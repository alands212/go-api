package api

import (
	"github.com/gofiber/fiber/v2"
)

func SetupUsersRoutes(app *fiber.App, tokenKey string) {
	s := start(tokenKey)
	grp := app.Group("/users")
	dni := app.Group("/dni")

	grp.Post("/", s.CreateUserHandler)
	grp.Post("/login", s.LoginHandler)
	dni.Post("/save", s.SaveDniHandler)

	grp.Use(jwtMiddleware(tokenKey)).Post("/permisos", s.PermisoHandler)
}
