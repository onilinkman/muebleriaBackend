package models

import (
	"database/sql"
	"fmt"
)

type Personal struct {
	Id_personal int    `json:"id_personal"`
	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	Url_image   string `json:"url_image"`
	Ci          string `json:"ci"`
	Date        string `json:"date"`
	Active      int    `json:"active"`
	Mount       int    `json:"mount"`
}

type Personals []Personal

const queryPersonal = `CREATE TABLE IF NOT EXISTS personal(
	id_personal INTEGER PRIMARY KEY AUTOINCREMENT,
	name 		TEXT NOT NULL,
	lastname 	TEXT NOT NULL,
	url_image 	TEXT,
	ci 			TEXT NOT NULL,
	date 		TEXT NOT NULL,
	active 		INTEGER DEFAULT 1
);`

func CreateTablePersonal() {
	_, err := Exec(queryPersonal)
	if err != nil {
		fmt.Println("Error in CreateTablePersonal")
		panic(err)
	}
}

func (personal *Personal) AddPersonal(url_img string) (int64, error) {
	query := `INSERT INTO personal(name,lastname,url_image,ci,date) VALUES(?,?,?,?,?)`
	result, err := Exec(query, &personal.Name, &personal.Lastname, url_img, &personal.Ci, &personal.Date)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return id, err
}

func queryGetPersonal(query string, mount, from int) (Personals, error) {
	personals := Personals{}
	rows, err := Query(query, mount, from)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		personal := Personal{}
		var url_image sql.NullString
		rows.Scan(&personal.Id_personal, &personal.Name, &personal.Lastname, &url_image,
			&personal.Ci, &personal.Date, &personal.Active, &personal.Mount)
		if url_image.Valid {
			personal.Url_image = url_image.String
		}
		personals = append(personals, personal)
	}
	return personals, nil
}

func GetPersonalsMount(mount, from int) (Personals, error) {
	query := `SELECT id_personal,name,lastname,url_image,ci,date,active,(select count(*) from personal) as mount from personal LIMIT ? OFFSET ?`
	return queryGetPersonal(query, mount, from)
}

func GetPersonals(mount, from int) (Personals, error) {
	query := `SELECT id_personal,name,lastname,url_image,ci,date,active,0 as mount from personal LIMIT ? OFFSET ?`
	return queryGetPersonal(query, mount, from)
}

func (personal *Personal) PutPersonalById(url_image string) error {
	query := `UPDATE personal SET name=?,lastname=?,url_image=?,ci=?,date=? 
	WHERE id_personal=?`

	_, err := Exec(query, &personal.Name, &personal.Lastname, url_image, &personal.Ci, &personal.Date, &personal.Id_personal)
	return err
}

func (personal *Personal) PutPersonalByIdWithoutImage() error {
	query := `UPDATE personal SET name=?,lastname=?,ci=?,date=? 
	WHERE id_personal=?`

	_, err := Exec(query, &personal.Name, &personal.Lastname, &personal.Ci, &personal.Date, &personal.Id_personal)
	return err
}
