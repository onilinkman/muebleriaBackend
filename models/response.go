package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Response forma de la estructura en la que se enviaran los datos
type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	write       http.ResponseWriter
}

//CreateDefaultResponse crea el response por defecto que se enviara
func CreateDefaultResponse(w http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		write:       w,
		contentType: "application/json",
	}
}

//SendNotFound envia el valor response
func SendNotFound(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotFound()
	response.Send()
}

//NotFound envia mensaje de not found cuando no se encuentra nada
func (r *Response) NotFound() {
	r.Status = http.StatusNotFound
	r.Message = "Resource Not Found"
}

//SendUnprocessableEntity envia cuando un dato no es posible procesar
func SendUnprocessableEntity(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.UnprocessableEntity()
	response.Send()
}

//UnprocessableEntity prepara el mensaje cuando el dato no es procesable
func (r *Response) UnprocessableEntity() {
	r.Status = http.StatusUnprocessableEntity
	r.Message = "Unprecessable Entity"
}

//SendNoContent envia un mensaje cuando no hay contenido
func SendNoContent(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NoContent()
	response.Send()
}

//NoContent se escribe el mensaje cuando no hay contenido
func (r *Response) NoContent() {
	r.Status = http.StatusNoContent
	r.Message = "No Content"
}

//SendData prepara para envia los datos
func SendData(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.Data = data
	response.Send()
}

//Send envia los datos
func (r *Response) Send() {
	r.write.Header().Set("Content-Type", r.contentType)
	r.write.WriteHeader(r.Status)
	output, _ := json.Marshal(&r)
	fmt.Fprintf(r.write, "%s", string(output))
}
