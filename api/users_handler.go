package api

import (
	"github.com/gofiber/fiber"
)

func (w *WebServices) CreateUserHandler(c *fiber.Ctx) {

	var cmd CreateUserCMD

	err := c.BodyParser(&cmd)

	res, err := w.Services.users.SaveUser(cmd)

	if err != nil {
		err = fiber.NewError(400, "cannot create user")
		c.Next(err)
		return
	}
	t := signToken(w.tokenKey, res.ID)

	savetoken := w.Services.users.Savetoken(t, res.UserSistema)

	if savetoken != nil {
		err = fiber.NewError(404, "user not found")
		c.Next(err)
		return
	}
	_ = c.JSON(struct {
		Token string `json:"token"`
	}{
		Token: t,
	})
}

func (w *WebServices) PermisoHandler(c *fiber.Ctx) {

	var cmd GetPermisoCMD

	_ = c.BodyParser(&cmd)

	bearer := c.Get("Authorization")

	userID := extractUserIDFromJWT(bearer, w.tokenKey)

	msg := w.users.GetPermiso(userID, cmd.SistemaId, cmd.PermisoSlug)

	_ = c.JSON(struct {
		R string `json:"acceso"`
	}{
		R: msg,
	})
}

func (w *WebServices) LoginHandler(c *fiber.Ctx) {

	var cmd LoginCMD

	err := c.BodyParser(&cmd)

	if err != nil {
		err = fiber.NewError(400, "cannot parse params")
		c.Next(err)
		return
	}

	id, usersistema := w.users.Login(cmd)

	if id == "" {
		err = fiber.NewError(404, "user not found")
		c.Next(err)
		return
	}

	t := signToken(w.tokenKey, id)

	savetoken := w.users.Savetoken(t, usersistema)

	if savetoken != nil {
		err = fiber.NewError(404, "user not found")
		c.Next(err)
		return
	}
	_ = c.JSON(struct {
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
