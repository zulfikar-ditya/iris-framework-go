package main

import (
	"fmt"
	controllers "iris-learn/Controllers"
	"os"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	// app := iris.New()
	app := iris.Default()

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

	// Book Url
	bookUrl := baseUrl.Party("/book")
	// bookUrl.Get("/", Book.GetListBook)
	// bookUrl.Post("/", Book.CreateBook)

	// Book Url Controller
	m := mvc.New(bookUrl)
	m.Handle(new(controllers.BookController))

	// route parameter
	// can be any types
	bookUrl.Get("/{id:int}", func(ctx iris.Context) {
		id := ctx.Params().Get("id")
		ctx.JSON(iris.Map{
			"message": "Book id " + id,
		})
	})

	// query string parameter
	bookUrl.Get("/query", func(ctx iris.Context) {
		firstName := ctx.URLParamDefault("firstName", "Guest")
		lastName := ctx.URLParamDefault("lastName", "Guest")

		ctx.JSON(iris.Map{
			"message": "Success",
			"data": iris.Map{
				"firstName": firstName,
				"lastName": lastName,
			},
		});
	})

	// Multipart/Urlencoded Form
	bookUrl.Post("/multipart", func(ctx iris.Context) {
		bookTitle := ctx.PostValue("title");
		
		if bookTitle == "" {
			ctx.StatusCode(iris.StatusUnprocessableEntity)
			ctx.JSON(iris.Map{
				"message": "Title is required",
				"errors": iris.Map{
					"title": "Title is required",
				},
			})

			return;
		}

		ctx.StatusCode(iris.StatusCreated)
		ctx.JSON(iris.Map{
			"message": "Book created",
		})
	})

	fmt.Println("Server running on port " + APP_PORT)
	app.Listen(APP_PORT)
}