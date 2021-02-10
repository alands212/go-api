package api

import (
	"time"

	"github.com/alands212/go-api/internal/logs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func jwtMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}

func signToken(tokenKey, id, user string, accesos []string, permisos []string) string {

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["id"] = id
	claims["user"] = user
	claims["roles"] = accesos
	claims["permisos"] = permisos

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(tokenKey))

	if err != nil {
		return ""
	}

	return t

}

func extractUserIDFromJWT(bearer, tokenKey string) string {

	token := bearer[7:]
	logs.Info(token)
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})

	if err != nil {
		return ""
	}

	if t.Valid {
		claims := t.Claims.(jwt.MapClaims)
		return claims["id"].(string)
	}

	return ""
}
