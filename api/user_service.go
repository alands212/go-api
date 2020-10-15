package api

import (
	"github.com/alands212/go-api/internal/database"
	"github.com/alands212/go-api/internal/logs"
	"github.com/gofiber/utils"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserCMD struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type UserSummary struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	JWT      string `json:"token"`
}

type UserGateway interface {
	SaveUser(cmd CreateUserCMD) (*UserSummary, error)
	Login(cmd LoginCMD) string
	AddWishMovie(userID, movieID, comment string) error
}

type UserService struct {
	*database.MySqlClient
}

func (us *UserService) SaveUser(cmd CreateUserCMD) (*UserSummary, error) {

	id := utils.UUID()

	/* encriptar contraseña */

	/* contraseña sin encriptar */
	contrasenaPlana := cmd.Password

	contrasenaPlanaComoByte := []byte(contrasenaPlana)
	hash, _ := bcrypt.GenerateFromPassword(contrasenaPlanaComoByte, bcrypt.DefaultCost) //DefaultCost es 10

	/* contraseña encriptada */
	password := string(hash)

	_, err := us.Exec(CreateUserQuery(), id, cmd.Username, password)

	if err != nil {

		logs.Error("cannot insert user" + err.Error())

		return nil, err
	}

	return &UserSummary{
		ID:       id,
		Username: cmd.Username,
		JWT:      "",
	}, nil
}

func (us *UserService) Login(cmd LoginCMD) string {
	var id string
	var password string

	err := us.QueryRow(GetLoginQuery(), cmd.Username).Scan(&id, &password)

	if err != nil {

		logs.Error(err.Error())

		return ""
	}

	/* comparar contraseña */

	hash := password
	hashComoByte := []byte(hash)

	contraseña := cmd.Password
	contraseñaComoByte := []byte(contraseña)

	error := bcrypt.CompareHashAndPassword(hashComoByte, contraseñaComoByte)
	if error != nil {
		logs.Error("Contraseña incorrecta")
		return ""
	}

	return id
}

func (us *UserService) AddWishMovie(userID, movieID, comment string) error {
	_, err := us.Exec(GetAddWishMovieQuery(), userID, movieID, comment)

	if err != nil {
		return err
	}

	return nil

}
