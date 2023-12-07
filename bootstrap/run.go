package bootstrap

import (
	config_database "iris-learn/config"
	"iris-learn/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func StartApp() {
	loadEnv()
	startDatabase()
	
	app := app()
	startRoute(app)

	app.Listen(os.Getenv("APP_PORT"))
}

func loadEnv() {
	envError := godotenv.Load()
	if envError != nil {
		panic("Error loading .env file")
	}
}

func startDatabase() {
	config_database.Connect()
}

func startRoute(app *iris.Application) {
	routes.LoadRoutes(app)
}

func app() (*iris.Application) {
	app := iris.Default()
	return app
}