package api

import (
	"github.com/gofiber/fiber/v2"
)

func SetupUsersRoutes(app *fiber.App, tokenKey string) {
	s := start(tokenKey)
	grp := app.Group("/users")

	grp.Post("/", s.CreateUserHandler)
	grp.Post("/login", s.LoginHandler)

	grp.Use(jwtMiddleware(tokenKey)).Post("/permisos", s.PermisoHandler)
}
