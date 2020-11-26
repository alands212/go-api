package api

import (
	"github.com/gofiber/fiber/v2"
)

func (w *WebServices) CreateUserHandler(c *fiber.Ctx) error {

	var cmd CreateUserCMD

	err := c.BodyParser(&cmd)

	res, err := w.Services.users.SaveUser(cmd)

	if err != nil {
		return fiber.NewError(400, "cannot create user")
	}
	t := signToken(w.tokenKey, res.ID)

	savetoken := w.Services.users.Savetoken(t, res.UserSistema)

	if savetoken != nil {
		return fiber.NewError(404, "user not found")
	}

	return c.JSON(struct {
		Token string `json:"token"`
	}{
		Token: t,
	})
}

func (w *WebServices) PermisoHandler(c *fiber.Ctx) error {

	var cmd GetPermisoCMD

	_ = c.BodyParser(&cmd)

	bearer := c.Get("Authorization")

	userID := extractUserIDFromJWT(bearer, w.tokenKey)

	msg := w.users.GetPermiso(userID, cmd.SistemaId, cmd.PermisoSlug)

	return c.JSON(struct {
		R string `json:"acceso"`
	}{
		R: msg,
	})
}

func (w *WebServices) LoginHandler(c *fiber.Ctx) error {

	var cmd LoginCMD

	err := c.BodyParser(&cmd)

	if err != nil {
		return fiber.NewError(400, "cannot parse params")
	}

	id, usersistema := w.users.Login(cmd)

	if id == "" {
		return fiber.NewError(404, "user not found")
	}

	t := signToken(w.tokenKey, id)

	savetoken := w.users.Savetoken(t, usersistema)

	if savetoken != nil {
		return fiber.NewError(404, "user not found")
	}

	return c.JSON(struct {
		Token string `json:"token"`
	}{
		Token: t,
	})
}

type LoginCMD struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	SistemaId string `json:"sistema_id"`
}

type WishMovieCMD struct {
	MovieID string `json:"movie_id"`
	Comment string `json:"comment"`
}

type GetPermisoCMD struct {
	SistemaId   string `json:"sistema_id"`
	PermisoSlug string `json:"permiso_slug"`
}
