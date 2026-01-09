package main

import (
	"log"
	"nutech-test/config"
	"os"

	"github.com/labstack/echo"
)

func main() {
	_, err := config.ConnectionDb()
	if err != nil {
		log.Printf("error connect to db %s", err)
	}

	//dependency injection

	//http routing using echo
	e := echo.New()

	port := os.Getenv("PORT")
	// if port empty just make it 8080
	if port == "" {
		port = "8080"
	}

	//route

	//protected route

	//start server
	if err := e.Start(":" + port); err != nil {
		log.Printf("error connect to server")
		return
	}
}