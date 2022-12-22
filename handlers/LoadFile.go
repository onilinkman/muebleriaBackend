package handlers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"strings"

	"os"
	"path/filepath"
)

const maxUploadSize = 15 * 1024 * 1024 //2MB

//RecibirArchivo el archivo que el usuario intente subir se guardara
func RecibirArchivo(r *http.Request, uploadPath, fileName string) error {
	//fmt.Println("recibiendo...")
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		return err
	}
	//fmt.Println("recibiendo...2", r.MultipartForm)
	file, fileHeader, err := r.FormFile("picture")
	if err != nil {
		return err
	}
	defer file.Close()

	fileSize := fileHeader.Size
	//fmt.Println("recibiendo...3")
	if fileSize > maxUploadSize {
		return errors.New("EXCEEDED SIZE LIMIT")
	}
	//fmt.Println("recibiendo...4")
	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	detectedFileType := http.DetectContentType(fileByte)
	switch detectedFileType {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
		//case "application", "pdf":
		break
	default:
		return errors.New("INVALID FILE TYPE")
	}

	fileEnding, err := mime.ExtensionsByType(detectedFileType)

	if err != nil {
		return errors.New("CANT READ FILE TYPE")
	}
	newPath := filepath.Join(uploadPath, fileName+fileEnding[0])
	newFile, err := os.Create(newPath)
	if err != nil {
		return errors.New("ERROR_TO_CREATE_NEW_FILE")
	}
	defer newFile.Close()

	if _, err := newFile.Write(fileByte); err != nil || newFile.Close() != nil {
		return errors.New("error to wirte file")
	}

	return nil
}

func SaveImagefromString(imageBase, uploadPath, fileName string) (string, error) {

	imageBytes, err := decodeDataURI(imageBase)

	if err != nil {
		imageBytes, err = base64.StdEncoding.DecodeString(imageBase)
		if err != nil {
			return "", err
		}
	}

	contentType := http.DetectContentType(imageBytes)

	fileSize := len(imageBytes)
	//fmt.Println("recibiendo...3")
	if fileSize > maxUploadSize {
		return "", errors.New("EXCEEDED SIZE LIMIT")
	}

	switch contentType {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
		break
	default:
		return "", errors.New("INVALID FILE TYPE")
	}

	format := strings.Replace(contentType, "image/", "", 1)
	var pathFile = fmt.Sprintf(`%s%s.%s`, uploadPath, fileName, format)

	err = ioutil.WriteFile(pathFile, imageBytes, 0644)

	return format, err
}

func decodeDataURI(dataURI string) ([]byte, error) {
	// Separamos el prefijo del resto de la cadena
	parts := strings.SplitN(dataURI, ",", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid Data URI")
	}

	// Decodificamos el resto de la cadena como base64
	bytes, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
