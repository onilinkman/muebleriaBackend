package models

import "database/sql"

type Store struct {
	Id_store int    `json:"id_store"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Active   int    `json:"active"`
	Mount    int    `json:"mount"`
}

type Stores []Store

const queryStore = `CREATE TABLE IF NOT EXISTS store(
	id_store INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	address TEXT NOT NULL,
	active INTEGER DEFAULT 1
)`

func CreateTableStore() {
	if _, err := Exec(queryStore); err != nil {
		panic(err)
	}
}

func (store *Store) AddStore() error {
	query := `INSERT INTO store(name,address) VALUES(?,?);`
	_, err := Exec(query, &store.Name, &store.Address)
	return err
}

func queryGetStore(query string, mount, from int) (Stores, error) {
	stores := Stores{}
	rows, err := Query(query, mount, from)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		store := Store{}
		var name sql.NullString //con esto se evita que cuando haya un valor nulo no deje de leer los demas elementos de la fila
		var address sql.NullString
		//var active sql.NullInt64
		rows.Scan(&store.Id_store, &name, &address, &store.Active, &store.Mount)
		if name.Valid {
			store.Name = name.String
		}
		if address.Valid {
			store.Address = address.String
		}
		//if active.Valid
		stores = append(stores, store)
	}
	return stores, nil
}

func GetAllStore(mount, from int) (Stores, error) {
	query := `SELECT id_store,name,address,active,0 as mount FROM store LIMIT ? OFFSET ?`
	stores, err := queryGetStore(query, mount, from)
	return stores, err
}

func GetAllStoreAndMount(mount, from int) (Stores, error) {
	query := `SELECT id_store,name,address,active,(select count(0) from store) as mount FROM store LIMIT ? OFFSET ?`
	stores, err := queryGetStore(query, mount, from)
	return stores, err
}
