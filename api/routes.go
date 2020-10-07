package api

import (
	"github.com/gofiber/fiber"
)

func SetupMoviesRoutes(app *fiber.App) {
	s := start()
	grp := app.Group("/movies")
	grp.Get("/", s.SearchMovieHandler)
}

func SetupUsersRoutes(app *fiber.App) {
	s := start()
	grp := app.Group("/users")
	grp.Post("/", s.CreateUserHandler)
}
