package main

import (
	"github.com/alands212/go-api/api"
	"github.com/gofiber/fiber"
)

func main() {

	app := fiber.New()
	api.SetupMoviesRoutes(app)
	_ = app.Listen("3001")
}
