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

	rol := &models.Rol{}
	if err := ConvertJson(w, r, rol); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	roles, err := models.GetAllRol(rol.Mount, rol.From)
	if err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	models.SendData(w, roles)

}

func GetAllRolsData(w http.ResponseWriter, r *http.Request) {
	rol := &models.Rol{}
	if err := ConvertJson(w, r, rol); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	roles, err := models.GetAllRolsData(rol.Mount, rol.From)
	if err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	models.SendData(w, roles)
}

func DeleteRolByID(w http.ResponseWriter, r *http.Request) {
	rol := &models.Rol{}
	if err := ConvertJson(w, r, rol); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	err := models.DeleteRolByID(rol.Id_rol)
	if err != nil {
		http.Error(w, "error not delete element", http.StatusNotImplemented)
		return
	}
	w.WriteHeader(http.StatusOK)
}
