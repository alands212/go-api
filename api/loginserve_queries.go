package api

func CreateUserQuery() string {
	return "insert into users (user, apellido, nombre, cuit, dni, nrotramitedni, password, created_at, updated_at, activo) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
}

func GetLoginQuery() string {
	return "select id, password from users where user = ?"
}

func GetUsersSistema() string {
	return "select id from userssistemas where users_id = ? and sistemas_id = ? and activo = 1"
}

func GetPermisoQuery() string {
	return "select permisos.id from permisos inner join rolspermisos ON rolspermisos.permisos_id=permisos.id inner join usersistemasrols ON usersistemasrols.rols_id=rolspermisos.rols_id inner join userssistemas ON userssistemas.id = usersistemasrols.userssistemas_id where userssistemas.users_id = ? and userssistemas.sistemas_id = ? and permisos.slug = ? and permisos.activo = 1"
}

func SavetokenQuery() string {
	return "update userssistemas set token = ?, updated_at = ? where id = ?"
}

func SaveUsersSistema() string {
	return "insert into userssistemas (users_id, sistemas_id, activo, created_at, updated_at) values (?, ?, 1, ?, ?)"
}
