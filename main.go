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
	//mux := http.NewServeMux()

	http.HandleFunc("/addRol", handlers.AddHeaders(v1.AddRol, "POST"))

	//getRol := handlers.HeadersMiddleware(http.HandlerFunc(v1.GetAllRols), "PUT")
	//http.Handle("/getRols", handlers.CORS(http.HandlerFunc(v1.GetAllRols)))
	http.HandleFunc("/getRols", v1.GetAllRols)

	//http.Handle("/addRol", handlers.HeadersMiddleware(htt, "POST"))
	fmt.Println("Corriendo servidor...")
	http.ListenAndServe(":8000", nil)
}
