package main

import (
	"github.com/alands212/go-api/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError
			var msg string
			// Retrieve the custom status code if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				msg = e.Message
			}

			if msg == "" {
				msg = "cannot process the http call"
			}

			// Send custom error page
			err = ctx.Status(code).JSON(internalError{
				Message: msg,
			})
			return nil
		},
	})

	key := "tokenKey"

	app.Use(recover.New())
	api.SetupUsersRoutes(app, key)

	_ = app.Listen(":3993")

}

type internalError struct {
	Message string `json:"message"`
}
