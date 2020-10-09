package main

import (
	"github.com/alands212/go-api/api"
	"github.com/alands212/go-api/internal"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main() {

	app := fiber.New()

	key := "tokenKey"

	internal.SetErrorHandler(app)
	app.Use(middleware.Recover())
	api.SetupMoviesRoutes(app, key)
	api.SetupUsersRoutes(app, key)

	_ = app.Listen("3001")

}
