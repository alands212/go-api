package api

func CreateDniQuery() string {
	return "insert into dni (numero, created_at) values (?, ?)"
}
func CreateUserQuery() string {
	return "insert into users (user, apellido, nombre, cuit, dni, tramite, password, created_at, updated_at, activo) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
}

func GetLoginQuery() string {
	return "select id, password, user from users where user = ?"
}

func GetUsersSistema() string {
	return "select id from userssistemas where users_id = ? and sistemas_id = ? and activo = 1"
}

func GetPermisoQuery() string {
	return "select permisos.id as permisoid from permisos inner join rolspermisos ON rolspermisos.permisos_id=permisos.id inner join userssistemasrols ON userssistemasrols.rols_id=rolspermisos.rols_id inner join userssistemas ON userssistemas.id = userssistemasrols.userssistemas_id where userssistemas.users_id = ? and userssistemas.sistemas_id = ? and permisos.scope = ? and permisos.activo = 1"
}

func GetAccessQuery() string {
	return "select rols.id, rols.scope from rols join userssistemasrols ON userssistemasrols.rols_id = rols.id inner join rolspermisos ON rolspermisos.rols_id=rols.id inner join userssistemas ON userssistemas.id = userssistemasrols.userssistemas_id where userssistemas.users_id = ? and userssistemas.sistemas_id = ? and rols.activo = 1 GROUP BY rols.scope"
}

func GetPermisosQuery() string {
	return "select permisos.scope from permisos join rolspermisos ON rolspermisos.permisos_id=permisos.id inner join userssistemasrols ON userssistemasrols.rols_id=rolspermisos.rols_id inner join userssistemas ON userssistemas.id = userssistemasrols.userssistemas_id inner join rols ON rols.id = userssistemasrols.rols_id where userssistemas.users_id = ? and userssistemas.sistemas_id = ? and rols.id = ? and permisos.activo = 1"
}

func SaveUsersSistema() string {
	return "insert into userssistemas (users_id, sistemas_id, activo, created_at, updated_at) values (?, ?, 1, ?, ?)"
}
