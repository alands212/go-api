
package api

import "github.com/alands212/go-api/internal/database"

type Services struct {
	search MovieSearch
	users  UserGateway
}

func NewServices() Services {
	client := database.NewMySQLClient()
	return Services{
		search: &MovieService{client},
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
