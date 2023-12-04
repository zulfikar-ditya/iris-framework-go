package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	// load env
	envError := godotenv.Load()
	if envError != nil {
		panic("Error loading .env file")
	}

	APP_NAME := os.Getenv("APP_NAME")
	APP_PORT := os.Getenv("APP_PORT")

	baseUrl := app.Party("/api")
	baseUrl.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"application name": APP_NAME,
			"message": "Welcome to api " + APP_NAME,
		})
	})

	fmt.Println("Server running on port " + APP_PORT)
	app.Listen(APP_PORT)
}