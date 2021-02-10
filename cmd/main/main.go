package main

import (

	/*descomentar en produccion

	"crypto/tls"
	"net"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"

	*/

	"github.com/alands212/go-api/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError
			var msg string
			// Retrieve the custom status code if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				msg = e.Message
			}

			if msg == "" {
				msg = "cannot process the http call"
			}

			// Send custom error page
			err = ctx.Status(code).JSON(internalError{
				Message: msg,
			})
			return nil
		},
	})

	key := "1G1f233G&HD)=F?¡{ÑFffJS}KFJHS754dhkks88eikrms{.xxjjf}55555"

	app.Use(recover.New())
	api.SetupUsersRoutes(app, key)

	/*DESCOMENTAR PARA GENERAR CERTIFICADO EN PRODUCCION

	m := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example.com"), // replaced with your domain.
		Cache:      autocert.DirCache("./certs"),
	}

	ssl := &tls.Config{
		GetCertificate: m.GetCertificate,
		// Must add acme.ALPNProto to NextProtos for TLS-ALPN-01.
		// Or just use m.TLSConfig() and remove it's NextProtos "h2" string.
		NextProtos: []string{
			"http/1.1", acme.ALPNProto,
		},
	}

	ln, _ := net.Listen("tcp", ":443")

	ln = tls.NewListener(ln, ssl)

	app.Listener(ln)

	*/

	//TRABAJANDO EN LOCAL
	_ = app.Listen(":3993")

}

type internalError struct {
	Message string `json:"message"`
}
