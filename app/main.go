package main

import (
	"log"
	"net/http"
	"nutech-test/config"
	"nutech-test/internal/controller"
	"nutech-test/internal/repository"
	"nutech-test/internal/service"
	"nutech-test/middleware"
	"os"
	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := config.ConnectionDb()
	if err != nil {
		log.Printf("error connect to db %s", err)
	}
	//close db connection if exit from main
	defer db.Close()

	
	//dependency injection
	//validator
	validate := validator.New()

	//repository
	userRepository := repository.NewUserRepository(db)
	bannerRepository := repository.NewBannerRepository(db)
	serviceRepository := repository.NewServiceRepository(db)

	//service
	userService := service.NewUserService(userRepository)
	bannerService := service.NewBannerService(bannerRepository)
	servicesService := service.NewServiceService(serviceRepository)

	//controller
	userController := controller.NewUserController(userService, validate)
	bannerController := controller.NewBannerController(bannerService)
	servicesController := controller.NewServiceController(servicesService)

	//http routing using echo
	e := echo.New()
	e.Use(middleware.LoggingMiddleware)
	e.HTTPErrorHandler = middleware.ErrorHandler
	//route
	e.POST("/registration", userController.CreateUser)
	e.POST("/login", userController.LoginUser)
	e.GET("/banner", bannerController.GetAllBanner)
	
	//protected route
	secretKey := os.Getenv("SECRET_KEY") 
	protectedRoute := e.Group("")
	protectedRoute.Use()
	protectedRoute.Use(echojwt.WithConfig(echojwt.Config{
        SigningKey: []byte(secretKey),
        ErrorHandler: func(c echo.Context, err error) error {
			auth := c.Request().Header.Get("Authorization")
			log.Printf("RAW Authorization header: [%s]", auth)
			log.Println("JWT key length:", len(secretKey))
            log.Printf("JWT Error: %v", err) 
            return c.JSON(http.StatusUnauthorized, map[string]interface{}{
                "message": "Token tidak tidak valid atau kadaluwarsa",
            })
        },
		SigningMethod: "HS256",
    }))

	protectedRoute.GET("/profile", userController.GetUserProfileByEmail)
	protectedRoute.PUT("/profile/update", userController.UpdateUserByEmail)
	protectedRoute.PUT("/profile/image", userController.UpdateUserImageByEmail)
	protectedRoute.GET("/services", servicesController.GetAllService)
	
	port := os.Getenv("PORT")
	// if port empty just make it 8080
	if port == "" {
		port = "8080"
	}
	//start server
	if err := e.Start(":" + port); err != nil {
		log.Printf("error connect to server")
		return
	}
}