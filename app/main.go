package main

import (
	"log"
	"nutech-test/config"
	"nutech-test/internal/controller"
	"nutech-test/internal/repository"
	"nutech-test/internal/service"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
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

	//service
	userService := service.NewUserService(userRepository)

	//controller
	userController := controller.NewUserController(userService, validate)

	//http routing using echo
	e := echo.New()

	port := os.Getenv("PORT")
	// if port empty just make it 8080
	if port == "" {
		port = "8080"
	}

	//route
	e.POST("/registration", userController.CreateUser)
	e.POST("/login", userController.LoginUser)

	//protected route
	e.GET("/profile", userController.GetUserProfileByEmail)
	e.GET("/profile/update", userController.UpdateUserByEmail)
	e.GET("/profile/image", userController.UpdateUserImageByEmail)

	//start server
	if err := e.Start(":" + port); err != nil {
		log.Printf("error connect to server")
		return
	}
}