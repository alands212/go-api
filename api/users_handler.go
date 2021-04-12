package api

import (
	"github.com/gofiber/fiber/v2"
)

func (w *WebServices) SaveDniHandler(c *fiber.Ctx) error {
	var dni CreateDniCMD

	err := c.BodyParser(&dni)

	res, err := w.Services.users.SaveDni(dni)

	if err != nil {
		return fiber.NewError(400, "cannot save dni")
	}

	return c.JSON(struct {
		Dni string `json:"NumeroDni"`
	}{
		Dni: res.Numero,
	})

}

func (w *WebServices) CreateUserHandler(c *fiber.Ctx) error {

	var cmd CreateUserCMD
	var accesos []string
	var permisos []string

	err := c.BodyParser(&cmd)

	res, err := w.Services.users.SaveUser(cmd)

	if err != nil {
		return fiber.NewError(400, "cannot create user")
	}

	accesos, permisos = w.users.GetAccess(res.ID, res.UserSistema)

	t := signToken(w.tokenKey, res.ID, cmd.User, accesos, permisos)

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
		Acceso string `json:"acceso"`
		ID     string `json:"user_id"`
	}{
		Acceso: msg,
		ID:     userID,
	})
}

func (w *WebServices) LoginHandler(c *fiber.Ctx) error {

	var cmd LoginCMD
	var accesos []string
	var permisos []string

	err := c.BodyParser(&cmd)

	if err != nil {
		return fiber.NewError(400, "cannot parse params")
	}

	id, sistemaid, user := w.users.Login(cmd)

	accesos, permisos = w.users.GetAccess(id, sistemaid)

	if id == "" {
		return fiber.NewError(404, "user not found")
	}

	t := signToken(w.tokenKey, id, user, accesos, permisos)

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
