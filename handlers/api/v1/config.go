package v1

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const GET = "GET"
const POST = "POST"
const PUT = "PUT"
const DELETE = "DELETE"

func ConvertJson(w http.ResponseWriter, r *http.Request, obj any) error {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &obj)
	return err
}

//setHeaders verific method and Set headers
func setHeaders(w http.ResponseWriter, r *http.Request, method string) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	if r.Method != method {
		err := errors.New("method no permitted")
		return err
	}
	return nil
}
