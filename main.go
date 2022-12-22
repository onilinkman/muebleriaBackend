package main

import (
	"fmt"
	"net/http"
	"os"

	"main.go/handlers"

	"main.go/configs"
	v1 "main.go/handlers/api/v1"
	"main.go/models"
)

func main() {

	defer models.CloseDB()

	createDir(configs.DIR_PHOTOS)
	createDir(configs.Concad_dirPhotos_personal())

	models.CreateTableRol()
	models.CreateTableStore()
	models.CreateTablePersonal()

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("photos"))))

	endPoints(mux)

	fmt.Println("Corriendo servidor...")
	fmt.Println(http.ListenAndServe("192.168.0.123:8080", mux))
}

func endPoints(mux *http.ServeMux) {
	//EndPoints for Rol
	addRol := handlers.HeadersMiddleware(http.HandlerFunc(v1.AddRol), "POST")
	mux.Handle("/addRol", addRol)

	getRol := handlers.HeadersMiddleware(http.HandlerFunc(v1.GetAllRols), "PUT")
	mux.Handle("/getRols", getRol)

	getRolsData := handlers.HeadersMiddleware(http.HandlerFunc(v1.GetAllRolsData), "PUT")
	mux.Handle("/getRolsData", getRolsData)

	deleteRolById := handlers.HeadersMiddleware(http.HandlerFunc(v1.DeleteRolByID), "DELETE")
	mux.Handle("/deleteRol", deleteRolById)

	//EndPoints for Store
	addStore := handlers.HeadersMiddleware(http.HandlerFunc(v1.AddStore), "POST")
	mux.Handle("/addStore", addStore)

	getAllStore := handlers.HeadersMiddleware(http.HandlerFunc(v1.GetAllStore), "PUT")
	mux.Handle("/getAllStore", getAllStore)

	getAllStoreAndMount := handlers.HeadersMiddleware(http.HandlerFunc(v1.GetAllStoreAndMount), "PUT")
	mux.Handle("/getAllStoreAndMount", getAllStoreAndMount)

	//EndPoints for Personal
	addPersonal := handlers.HeadersMiddleware(http.HandlerFunc(v1.AddPersonal), "POST")
	mux.Handle("/addPersonal", addPersonal)

	getPersonalsMount := handlers.HeadersMiddleware(http.HandlerFunc(v1.GetPersonalsMount), "PUT")
	mux.Handle("/getPersonalsMount", getPersonalsMount)

	getPersonals := handlers.HeadersMiddleware(http.HandlerFunc(v1.GetPersonals), "PUT")
	mux.Handle("/getPersonals", getPersonals)

	putPersonalById := handlers.HeadersMiddleware(http.HandlerFunc(v1.PutPersonalById), "PUT")
	mux.Handle("/putPersonalById", putPersonalById)

	putPersonalByIdWithoutImage := handlers.HeadersMiddleware(http.HandlerFunc(v1.PutPersonalByIdWithoutImage), "PUT")
	mux.Handle("/putPersonalByIdWithoutImage", putPersonalByIdWithoutImage)
}

func createDir(nameDir string) {
	if !existDir(nameDir) {
		erro := os.Mkdir(nameDir, 0750)
		if erro != nil {
			panic("Error to create dir")
		}
	}
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
