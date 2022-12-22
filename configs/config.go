package configs

import "fmt"

const DIR_PHOTOS = "./photos"
const DIR_PHOTOS_PERSONAL = "/personal"

func Concad_dirPhotos_personal() string {
	return fmt.Sprintf("%s%s/", DIR_PHOTOS, DIR_PHOTOS_PERSONAL)
}
