package main

import (
	"fmt"
	"net/http"

	"main.go/handlers"

	v1 "main.go/handlers/api/v1"
	"main.go/models"
)

func main() {

	defer models.CloseDB()

	models.CreateTableRol()
	mux := http.NewServeMux()

	endPoints(mux)

	fmt.Println("Corriendo servidor...")
	http.ListenAndServe(":8000", mux)
}

func endPoints(mux *http.ServeMux) {
	addRol := handlers.HeadersMiddleware(http.HandlerFunc(v1.AddRol), "POST")
	mux.Handle("/addRol", addRol)

	getRol := handlers.HeadersMiddleware(http.HandlerFunc(v1.GetAllRols), "PUT")
	mux.Handle("/getRols", getRol)

	getRolsData := handlers.HeadersMiddleware(http.HandlerFunc(v1.GetAllRolsData), "PUT")
	mux.Handle("/getRolsData", getRolsData)
}
