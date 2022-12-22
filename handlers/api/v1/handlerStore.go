package v1

import (
	"fmt"
	"net/http"

	"main.go/models"
)

func AddStore(w http.ResponseWriter, r *http.Request) {
	store := &models.Store{}
	if err := ConvertJson(w, r, store); err != nil {
		fmt.Println("error AddStore:")
		fmt.Println(err)
		models.SendUnprocessableEntity(w)
		return
	}
	if err := store.AddStore(); err != nil {
		fmt.Println("Error AddStore:")
		fmt.Println(err)
		http.Error(w, "error to addStore", http.StatusNotAcceptable)
	}

}

func GetAllStore(w http.ResponseWriter, r *http.Request) {
	mountFrom := &models.MountFrom{}
	if err := ConvertJson(w, r, mountFrom); err != nil {
		fmt.Println("error GetAllStore:")
		fmt.Println(err)
		models.SendUnprocessableEntity(w)
		return
	}
	stores, err := models.GetAllStore(mountFrom.Mount, mountFrom.From)
	if err != nil {
		fmt.Println("Error GetAllStore")
		fmt.Println(err)
		http.Error(w, "Error in call database", http.StatusNotFound)
		return
	}
	models.SendData(w, stores)
}

func GetAllStoreAndMount(w http.ResponseWriter, r *http.Request) {
	mountFrom := &models.MountFrom{}
	if err := ConvertJson(w, r, mountFrom); err != nil {
		fmt.Println("error GetAllStoreAndMount:")
		fmt.Println(err)
		models.SendUnprocessableEntity(w)
		return
	}
	stores, err := models.GetAllStoreAndMount(mountFrom.Mount, mountFrom.From)
	if err != nil {
		fmt.Println("Error GetAllStoreAndMount")
		fmt.Println(err)
		http.Error(w, "Error in call database", http.StatusNotFound)
		return
	}
	models.SendData(w, stores)
}
