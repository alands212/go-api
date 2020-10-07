package api

import "github.com/gofiber/fiber"

func (w *WebServices) CreateUserHandler(c *fiber.Ctx) {
	c.JSON(struct {
		Test string `json:"test"`
	}{
		Test: "working",
	})
}
