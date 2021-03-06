# go-api

> #### Contenido
>
> - Framework: [fiber v2.1.4](https://docs.gofiber.io/v/2.x/).
> - DB: MySQL.
> - Autentificacion: JWT.

#### Videos - tomas lingotti
- [Parte I - Arquitectura inicial, mods y las rutas web](https://www.youtube.com/watch?v=vQtkgavdxk4).
- [Parte II - MySQL y centralizar el manejo de errores](https://www.youtube.com/watch?v=hhpR825EXAY).
- [Parte III - Autenticación con JWT](https://www.youtube.com/watch?v=LXr1RJaaGhA).
- [Parte IV - Autenticacion del cliente y validacion JWT con Fiber](https://www.youtube.com/watch?v=3Uscn6CNEVU).
- [Extra - Fiber v2](https://www.youtube.com/watch?v=GW-vQpecS2M)

#### Conceptos - codigofacilito
- [Punteros](https://www.youtube.com/watch?v=V0cdxZCEzHE).
- [Structs](https://www.youtube.com/watch?v=aBkPQr2VTMc).
- [Métodos](https://www.youtube.com/watch?v=quA5nX8mceY).
- [Interfaces](https://www.youtube.com/watch?v=OeCtHLvf-Eo).
- [Goroutines](https://www.youtube.com/watch?v=rF3VP10S9SM).


### Hashear contraseña en Go
- [bcrypt](https://parzibyte.me/blog/2018/05/31/hasheando-comprobando-contrasenas-golang/).


*************************

### Documentacion

#### Registro

>
> - POST
>
> - Url: /users
>
> - Headers:
>
>   **Key**: Content-Type  **Value**: application/json
>
> - Body:
>
>```json
> {
>	"user": "facundo11",
>	"apellido": "Rios",
>	"nombre": "Facundo",
>	"cuit": "20213243129",
>	"dni": "21324312",
>	"tramite": "1231152224312456789",
>	"password": "contrasena123",
>	"sistema_id": "1"
> }
>```
>
> - Rerturn: token
>
>```json
> {
>    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDY2NTQ3NzUsImlkIjoiNiJ9.BJmXeGbZtikI2JtXst6s1ogP6L-y4n9Mi79SwdsHtHI"
> }
>```
>

#### Login

>
> - POST
>
> - Url: /users/login
>
> - Headers:
>
>   **Key**: Content-Type  **Value**: application/json
>
> - Body:
>
>```json
> {
>	"username": "alands212",
>	"password": "contrasena123",
>	"sistema_id": "1"
> }
>```
>
> - Rerturn: token
>```json
> {
>    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDY2NTUwNjUsImlkIjoiMSJ9.lIXRoxQGpywGeyp7r98aSrjPIGmQ22utaH2-aeK9X0o"
> }
>```
>



#### Permiso acceso

>
> - POST
>
> - Url: /users/permisos
>
> - Headers:
>
>   **Key**: Authorization  **Value**: Bearer token-del-usuario
>
>   **Key**: Content-Type  **Value**: application/json
>
> - Body:
>
>```json
> {
>	"sistema_id": "1",
>	"permiso_slug": "permiso.index"
> }
>```
>
> - Rerturn: (true o false)
>```json
> {
>    "acceso": "true"
> }
>```
>

#### Token contenido

>
>	claims["exp"]
>
>	claims["id"]
>

*************************

## TLS

Para ejecutar api en produccion con TLS descomentar las lineas correspondientes en [main.go](https://github.com/Corrientes-Telecomunicaciones/loginservice/blob/main/cmd/main/main.go)

Editar en [main.go](https://github.com/Corrientes-Telecomunicaciones/loginservice/blob/main/cmd/main/main.go) **example.com** por el dominio correspondiente

*************************

## Conexion DB

Para cambiar la conexion a la base de datos editar en el archivo [mysql_client.go](https://github.com/Corrientes-Telecomunicaciones/loginservice/blob/main/internal/database/mysql_client.go) 

```
db, err := sql.Open("mysql", "root:root@tcp(mariadb:3306)/mydb")
```
- **root** 1er usuario de la base de datos
- **root** 2do contraseña de la base de datos
- **mariadb:3306** direccion ip y puerto de la base de datos
- **mydb** nombre de la tabla

*************************

## Docker Compose

### Build
```shell script
docker-compose build
```

### Up
```shell script
docker-compose up -d
```
**-d** para ejecutar el servicios en segundo plano