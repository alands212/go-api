package api

import "github.com/alands212/go-api/internal/database"

type Services struct {
	users UserGateway
}

func NewServices() Services {
	client := database.NewMySQLClient()
	return Services{
		users: &UserService{client},
	}
}

type WebServices struct {
	Services
	tokenKey string
}

func start(tokenKey string) *WebServices {
	return &WebServices{NewServices(), tokenKey}
}
