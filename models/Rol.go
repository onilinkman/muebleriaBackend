package models

//Rol modelo de rol
type Rol struct {
	Id_rol int    `json:"id_rol"`
	Nombre string `json:"nombre"`
	Mount  int
	From   int
}

type Roles []Rol

//queryRol this is a query for create table rol
const queryRol = `CREATE TABLE IF NOT EXISTS rol(
	id_rol INTEGER PRIMARY KEY AUTOINCREMENT,
	nombre TEXT
);`

func CreateTableRol() {
	Exec(queryRol)
}

func (r *Rol) AddRol() error {
	query := `INSERT INTO rol(nombre) VALUES(?);`
	_, err := Exec(query, &r.Nombre)
	return err
}

/**
GetAllRol return a structure type Roles

	-mount: that mount limite the rows
	-from: from which row do you want to continue?
*/
func GetAllRol(mount, from int) (Roles, error) {
	roles := Roles{}
	query := `SELECT id_rol,nombre from rol LIMIT ? OFFSET ?;`
	rows, err := Query(query, mount, from)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rol := Rol{}
		rows.Scan(&rol.Id_rol, &rol.Nombre)
		roles = append(roles, rol)
	}
	return roles, nil
}

func GetAllRolsData(mount, from int) (Roles, error) {
	roles := Roles{}
	query := `SELECT id_rol,nombre,(SELECT COUNT(*) FROM rol) as numRows from rol LIMIT ? OFFSET ?;`
	rows, err := Query(query, mount, from)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rol := Rol{}
		rows.Scan(&rol.Id_rol, &rol.Nombre, &rol.Mount)
		roles = append(roles, rol)
	}
	return roles, nil
}

func DeleteRolByID(id_rol int) error {
	query := `DELETE FROM rol WHERE id_rol=?`
	_, err := db.Query(query, id_rol)
	return err
}
