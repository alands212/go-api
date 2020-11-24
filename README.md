# go-api

> #### Contenido
>
> - Framework: [fiber 1.14.6](https://docs.gofiber.io/v/1.x/).
> - DB: MySQL.
> - Autentificacion: JWT.

#### Videos - tomas lingotti
- [Parte I - Arquitectura inicial, mods y las rutas web](https://www.youtube.com/watch?v=vQtkgavdxk4).
- [Parte II - MySQL y centralizar el manejo de errores](https://www.youtube.com/watch?v=hhpR825EXAY).
- [Parte III - Autenticación con JWT](https://www.youtube.com/watch?v=LXr1RJaaGhA).
- [Parte IV - Autenticacion del cliente y validacion JWT con Fiber](https://www.youtube.com/watch?v=3Uscn6CNEVU).


#### Conceptos - codigofacilito
- [Punteros](https://www.youtube.com/watch?v=V0cdxZCEzHE).
- [Structs](https://www.youtube.com/watch?v=aBkPQr2VTMc).
- [Métodos](https://www.youtube.com/watch?v=quA5nX8mceY).
- [Interfaces](https://www.youtube.com/watch?v=OeCtHLvf-Eo).
- [Goroutines](https://www.youtube.com/watch?v=rF3VP10S9SM).


### Hashear contraseña en Go
- [bcrypt](https://parzibyte.me/blog/2018/05/31/hasheando-comprobando-contrasenas-golang/).



### Documentacion

#### Registro

>
> - POST
>
> - Url: /users
>
> - Headers:
>
>   Key: Content-Type 
>
>   Value: application/json
>
> - Body:
>
> {
>
>	"username": "UsuarioNuevo",
>
>	"password": "contrasena123"
>
> }
>
> - Rerturn:
>
> {
>
>    "id": "7149849b-9743-40bc-afa2-15338e70e892",
>
>    "username": "Nombreusuario",
>
>   "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZW1haWwiOiJtYXJ0aW4uZHMuMjEyQGdtYWlsLmNvbSIsImV4cCI6MTYwNjQ4MDQ2MCwic3ViIjoiNzE0OTg0OWItOTc0My00MGJjLWFmYTItMTUzMzhlNzBlODkyIiwidXNlcm5hbWUiOiJOb21icmV1c3VhcmlvIn0.vmHiN1HeLcCWuOygeLLqatL1RediaKQxZOUd7gA7FOI"
>
> }
>

#### Login

>
> - POST
>
> - Url: /users/login
>
> - Headers:
>
>   Key: Content-Type 
>
>   Value: application/json
>
> - Body:
>
> {
>
>	"username": "Nombreusuario",
>
>	"password": "contrasena123"
>
> }
>
> - Rerturn:
>```json
> {
>
>    "id": "7149849b-9743-40bc-afa2-15338e70e892",
>
>    "username": "Nombreusuario",
>
>   "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZW1haWwiOiJtYXJ0aW4uZHMuMjEyQGdtYWlsLmNvbSIsImV4cCI6MTYwNjQ4MDY0Niwic3ViIjoiNzE0OTg0OWItOTc0My00MGJjLWFmYTItMTUzMzhlNzBlODkyIiwidXNlcm5hbWUiOiJOb21icmV1c3VhcmlvIn0.cQ6OpjSEP56Jq8TxZJFD9FWARHim_EVxMVFyx7REZBo"
>
> }
>```
>

#### Movies

>
> - GET
>
> - Url: /movies
>
> - Rerturn:
>
>[
>
>    {
>
>        "id": "tt0468569",
>
>        "title": "The Dark Knight",
>
>        "cast": "Christian Bale",
>
>        "release_date": "2008-08-13",
>
>        "genre": "Acción,Drama",
>
>        "director": "Christopher Nolan"
>
>    },
>
>    {
>
>        "id": "tt4154796",
>
>        "title": "Avengers: Endgame",
>
>        "cast": "Robert Downey Jr, Chris Evans, Mark Ruffalo, Chris Hemsworth, Scarlett Johansson, etc.",
>
>        "release_date": "2019-04-25",
>
>        "genre": "Superhéroes, Acción, Ciencia ficción",
>
>        "director": "Anthony y Joe Russo"
>
>    }
>
>]
>



#### Token contenido

>
>	claims["admin"] = true
>
>	claims["username"] = user
>
>	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
>
>	claims["sub"] = id
>
>	claims["email"] = "martin.ds.212@gmail.com"
>
