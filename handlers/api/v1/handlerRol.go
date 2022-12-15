package v1

import (
	"fmt"
	"net/http"

	"main.go/models"
)

func AddRol(w http.ResponseWriter, r *http.Request) {
	if r.Method == POST {
		rol := &models.Rol{}
		if err := ConvertJson(w, r, rol); err != nil {
			fmt.Println(err)
			models.SendUnprocessableEntity(w)
		}
		if err := rol.AddRol(); err != nil {
			fmt.Println(err)
			models.SendUnprocessableEntity(w)
		}
	}
}

func GetAllRols(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	if r.Method == PUT {
		/* if err := setHeaders(w, r, PUT); err != nil {
			http.Error(w, "Method not permitted", http.StatusMethodNotAllowed)
			return
		} */
		fmt.Println("hola")
		rol := &models.Rol{}
		if err := ConvertJson(w, r, rol); err != nil {
			fmt.Println(1, err)
			models.SendUnprocessableEntity(w)
			return
		}
		roles, err := models.GetAllRol(rol.Mount, rol.From)
		if err != nil {
			fmt.Println(3, err)
			models.SendUnprocessableEntity(w)
			return
		}
		models.SendData(w, roles)

	}
}
