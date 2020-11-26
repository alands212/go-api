package api

import (
	"fmt"
	"time"

	"github.com/alands212/go-api/internal/database"
	"github.com/alands212/go-api/internal/logs"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserCMD struct {
	User          string `json:"user"`
	Apellido      string `json:"apellido"`
	Nombre        string `json:"nombre"`
	Cuit          string `json:"cuit"`
	Dni           string `json:"dni"`
	Nrotramitedni string `json:"nrotramitedni"`
	Password      string `json:"password"`
	Created_at    string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
	Activo        string `json:"activo"`
	SistemaId     string `json:"sistema_id"`
}

type UserSummary struct {
	ID          string `json:"id"`
	UserSistema string `json:"user_sistema"`
}

type UserGateway interface {
	SaveUser(cmd CreateUserCMD) (*UserSummary, error)
	Login(cmd LoginCMD) (string, string)
	GetPermiso(userID, sistemaID, permisoSlug string) string
	Savetoken(t, usersistemaid string) error
}

type UserService struct {
	*database.MySqlClient
}

func (us *UserService) SaveUser(cmd CreateUserCMD) (*UserSummary, error) {

	/* encriptar contraseña */

	/* contraseña sin encriptar */
	contrasenaPlana := cmd.Password

	contrasenaPlanaComoByte := []byte(contrasenaPlana)
	hash, _ := bcrypt.GenerateFromPassword(contrasenaPlanaComoByte, bcrypt.DefaultCost) //DefaultCost es 10

	/* contraseña encriptada */
	password := string(hash)

	/* fecha */
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	/* Crear usuario */
	_, err := us.Exec(CreateUserQuery(), cmd.User, cmd.Apellido, cmd.Nombre, cmd.Cuit, cmd.Dni, cmd.Nrotramitedni, password, fecha, fecha, 1)

	if err != nil {

		logs.Error("cannot insert user" + err.Error())

		return nil, err
	}

	/* Login usuario nuevo */

	var id string

	/* Buscando al usuario */
	_ = us.QueryRow(GetLoginQuery(), cmd.User).Scan(&id, &password)

	/* Asignar a sistema */
	_, erro := us.Exec(SaveUsersSistema(), id, cmd.SistemaId, fecha, fecha)

	if erro != nil {

		logs.Error("Error al asignar sistema al usuario: " + erro.Error())

		return nil, erro
	}

	var usersistemaid string

	_ = us.QueryRow(GetUsersSistema(), id, cmd.SistemaId).Scan(&usersistemaid)

	return &UserSummary{
		ID:          id,
		UserSistema: usersistemaid,
	}, nil
}

func (us *UserService) Login(cmd LoginCMD) (string, string) {

	var id string
	var password string

	/* Buscando al usuario */
	err := us.QueryRow(GetLoginQuery(), cmd.Username).Scan(&id, &password)

	if err != nil {

		logs.Error(err.Error())

		return "", ""
	}

	/* Busacando al usuario en algun sistema */
	var usersistemaid string

	usersistema := us.QueryRow(GetUsersSistema(), id, cmd.SistemaId).Scan(&usersistemaid)

	if usersistema != nil {

		logs.Error(usersistema.Error())

		return "", ""
	}

	/* comparar contraseña */
	hash := password
	hashComoByte := []byte(hash)

	contraseña := cmd.Password
	contraseñaComoByte := []byte(contraseña)

	error := bcrypt.CompareHashAndPassword(hashComoByte, contraseñaComoByte)
	if error != nil {
		logs.Error("Contraseña incorrecta")
		return "", ""
	}

	return id, usersistemaid

}

func (us *UserService) GetPermiso(userID, sistemaID, permisoSlug string) string {

	var permisoid string

	res := us.QueryRow(GetPermisoQuery(), userID, sistemaID, permisoSlug).Scan(&permisoid)

	if res != nil {
		return "false"
	}

	return "true"

}

func (us *UserService) Savetoken(t, usersistemaid string) error {

	/* fecha */
	ahora := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		ahora.Year(), ahora.Month(), ahora.Day(),
		ahora.Hour(), ahora.Minute(), ahora.Second())

	_, err := us.Exec(SavetokenQuery(), t, fecha, usersistemaid)

	if err != nil {
		return err
	}

	return nil

}
