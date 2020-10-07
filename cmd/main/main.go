package main

import (
	"github.com/alands212/go-api/api"
	"github.com/alands212/go-api/internal"
	"github.com/gofiber/fiber"
)

func main() {

	app := fiber.New()
	internal.SetErrorHandler(app)
	api.SetupMoviesRoutes(app)
	api.SetupUsersRoutes(app)
	_ = app.Listen("3001")
	
}
