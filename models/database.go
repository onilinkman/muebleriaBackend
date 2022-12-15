package models

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var db *sql.DB

const nameDir = "DataBase"
const dirDB = "./"
const nameDB = "DataBussiness"

func init() {
	connectDB()
}

func connectDB() {
	var err error
	if !existDir(nameDir) {
		erro := os.Mkdir(nameDir, 0750)
		if erro != nil {
			panic("Error to create dir")
		}
	}
	db, err = sql.Open("sqlite", filepath.Join(dirDB, nameDir, nameDB))
	if err != nil {
		panic("Error to connect with database")
	}
}

func Exec(query string, args ...any) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Query(query string, args ...any) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func CloseDB() {
	db.Close()
}

func existDir(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
