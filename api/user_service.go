package api

import (
	"fmt"
	"time"

	"github.com/alands212/go-api/internal/database"
	"github.com/alands212/go-api/internal/logs"
	"golang.org/x/crypto/bcrypt"
)

type CreateDniCMD struct {
	Numero string `json:"numero"`
}
type CreateUserCMD struct {
	User       string `json:"user"`
	Apellido   string `json:"apellido"`
	Nombre     string `json:"nombre"`
	Cuit       string `json:"cuit"`
	Dni        string `json:"dni"`
	Tramite    string `json:"tramite"`
	Password   string `json:"password"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Activo     string `json:"activo"`
	SistemaId  string `json:"sistema_id"`
}

type UserSummary struct {
	ID          string `json:"id"`
	UserSistema string `json:"user_sistema"`
}

type UserGateway interface {
	SaveUser(cmd CreateUserCMD) (*UserSummary, error)
	SaveDni(cmd CreateDniCMD) (*CreateDniCMD, error)
	Login(cmd LoginCMD) (string, string, string)
	GetAccess(userID, sistemaID string) ([]string, []string)
	GetPermiso(userID, sistemaID, permisoSlug string) string
}

type UserService struct {
	*database.MySqlClient
}

func (us *UserService) SaveDni(cmd CreateDniCMD) (*CreateDniCMD, error) {
	/* fecha */
	t := time.Now()

	fecha := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	/* Crear usuario */
	_, err := us.Exec(CreateDniQuery(), cmd.Numero, fecha)

	if err != nil {

		logs.Error("cannot insert user" + err.Error())

		return nil, err
	}

	return &CreateDniCMD{
		Numero: cmd.Numero,
	}, nil
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
	_, err := us.Exec(CreateUserQuery(), cmd.User, cmd.Apellido, cmd.Nombre, cmd.Cuit, cmd.Dni, cmd.Tramite, password, fecha, fecha, 1)

	if err != nil {

		logs.Error("cannot insert user" + err.Error())

		return nil, err
	}

	/* Login usuario nuevo */

	var id string
	var user string

	/* Buscando al usuario */
	_ = us.QueryRow(GetLoginQuery(), cmd.User).Scan(&id, &password, &user)

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

func (us *UserService) Login(cmd LoginCMD) (string, string, string) {

	var id string
	var password string
	var user string

	/* Buscando al usuario */
	err := us.QueryRow(GetLoginQuery(), cmd.Username).Scan(&id, &password, &user)

	if err != nil {

		logs.Error(err.Error())

		return "", "", ""
	}

	/* Busacando al usuario en algun sistema */
	var usersistemaid string

	usersistema := us.QueryRow(GetUsersSistema(), id, cmd.SistemaId).Scan(&usersistemaid)

	if usersistema != nil {

		logs.Error(usersistema.Error())

		return "", "", ""
	}

	/* comparar contraseña */
	hash := password
	hashComoByte := []byte(hash)

	contraseña := cmd.Password
	contraseñaComoByte := []byte(contraseña)

	error := bcrypt.CompareHashAndPassword(hashComoByte, contraseñaComoByte)
	if error != nil {
		logs.Error("Contraseña incorrecta")
		return "", "", ""
	}

	return id, cmd.SistemaId, user

}

func (us *UserService) GetPermiso(userID, sistemaID, permisoSlug string) string {

	var permisoid string

	res := us.QueryRow(GetPermisoQuery(), userID, sistemaID, permisoSlug).Scan(&permisoid)

	if res != nil {
		return "false"
	}

	return "true"

}

func (us *UserService) GetAccess(userID, sistemaID string) ([]string, []string) {

	var rol string
	var rolid string
	var permiso string

	res, err := us.Query(GetAccessQuery(), userID, sistemaID)

	if err != nil {
		logs.Error(err)
	}

	defer res.Close()

	var acceso []string
	var perm []string

	for res.Next() {
		err := res.Scan(&rolid, &rol)
		if err != nil {
			logs.Error(err)
		}

		resperm, err := us.Query(GetPermisosQuery(), userID, sistemaID, rolid)

		defer resperm.Close()

		for resperm.Next() {

			err := resperm.Scan(&permiso)
			if err != nil {
				logs.Error(err)
			}

			perm = append(perm, permiso)

		}

		acceso = append(acceso, rol)

	}

	permisos := unique(perm)

	if err := res.Err(); err != nil {
		logs.Error(err)
	}

	return acceso, permisos

}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
