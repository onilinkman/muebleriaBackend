package v1

import (
	"fmt"
	"net/http"

	"main.go/configs"
	"main.go/handlers"
	"main.go/models"
)

func AddPersonal(w http.ResponseWriter, r *http.Request) {
	personal := &models.Personal{}

	if err := ConvertJson(w, r, personal); err != nil {
		fmt.Println("Error in AddPersonal")
		fmt.Println(err)
		models.SendUnprocessableEntity(w)
		return
	}

	var fileName = fmt.Sprintf(`%s%s`, personal.Lastname, personal.Ci)
	var pathActual = fmt.Sprintf(`%s/`, configs.Concad_dirPhotos_personal())
	var url_image = ""

	format, err := handlers.SaveImagefromString(personal.Url_image, pathActual, fileName)
	if err == nil {
		url_image = fmt.Sprintf(`%s.%s`, fileName, format)
	}
	_, err = personal.AddPersonal(url_image)
	if err != nil {
		fmt.Println("Error AddPersonal:")
		fmt.Println(err)
		http.Error(w, "error to addPersonal", http.StatusNotAcceptable)
	}
}

func GetPersonalsMount(w http.ResponseWriter, r *http.Request) {
	mountFrom := &models.MountFrom{}
	if err := ConvertJson(w, r, mountFrom); err != nil {
		fmt.Println("error GetPersonalsMount:")
		fmt.Println(err)
		models.SendUnprocessableEntity(w)
		return
	}
	personals, err := models.GetPersonalsMount(mountFrom.Mount, mountFrom.From)
	if err != nil {
		fmt.Println("Error GetPersonalsMount")
		fmt.Println(err)
		http.Error(w, "Error in call database", http.StatusNotFound)
		return
	}
	models.SendData(w, personals)
}

func GetPersonals(w http.ResponseWriter, r *http.Request) {
	mountFrom := &models.MountFrom{}
	if err := ConvertJson(w, r, mountFrom); err != nil {
		fmt.Println("error GetPersonalsMount:")
		fmt.Println(err)
		models.SendUnprocessableEntity(w)
		return
	}
	personals, err := models.GetPersonals(mountFrom.Mount, mountFrom.From)
	if err != nil {
		fmt.Println("Error GetPersonalsMount")
		fmt.Println(err)
		http.Error(w, "Error in call database", http.StatusNotFound)
		return
	}
	models.SendData(w, personals)
}

func PutPersonalById(w http.ResponseWriter, r *http.Request) {
	personal := &models.Personal{}

	if err := ConvertJson(w, r, personal); err != nil {
		fmt.Println("Error in AddPersonal")
		fmt.Println(err)
		models.SendUnprocessableEntity(w)
		return
	}

	var fileName = fmt.Sprintf(`%s%s`, personal.Lastname, personal.Ci)
	var pathActual = fmt.Sprintf(`%s/`, configs.Concad_dirPhotos_personal())
	var url_image = ""

	format, err := handlers.SaveImagefromString(personal.Url_image, pathActual, fileName)
	if err == nil {
		url_image = fmt.Sprintf(`%s.%s`, fileName, format)
	}
	err = personal.PutPersonalById(url_image)
	if err != nil {
		fmt.Println("Error PutPersonalById:")
		fmt.Println(err)
		http.Error(w, "error to PutPersonalById", http.StatusNotAcceptable)
	}
}

func PutPersonalByIdWithoutImage(w http.ResponseWriter, r *http.Request) {
	personal := &models.Personal{}

	if err := ConvertJson(w, r, personal); err != nil {
		fmt.Println("Error in PutPersonalByIdWithoutImage")
		fmt.Println(err)
		models.SendUnprocessableEntity(w)
		return
	}

	if err := personal.PutPersonalByIdWithoutImage(); err != nil {
		fmt.Println("Error PutPersonalByIdWithoutImage:")
		fmt.Println(err)
		http.Error(w, "error to PutPersonalByIdWithoutImage", http.StatusNotAcceptable)
	}
}
