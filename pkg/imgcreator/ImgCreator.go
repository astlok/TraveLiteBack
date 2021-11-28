package imgcreator

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"math/rand"
	"os"
	"strings"
)

type ImgCreatorI interface {
	CreateImg(imgBase64 string) (string, error)
}

type imgCreator struct {}

var imgCreatorInstance *imgCreator

const (
	fullAcces = 0777
	quality   = 75
)

func ImgCreator() ImgCreatorI {
	if imgCreatorInstance == nil {
		imgCreatorInstance = new(imgCreator)
	}
	return imgCreatorInstance
}

func (i *imgCreator) CreateImg(imgBase64 string) (string, error) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imgBase64))
	m, _, err := image.Decode(reader)
	if err != nil {
		return "", err
	}

	salt := make([]byte, 8)
	_, err = rand.Read(salt)
	if err != nil {
		return "", err
	}
	jpegFilename, err := randomFilename16Char()
	if err != nil {
		return "", err
	}
	jpegFilename += ".jpeg"
	f, err := os.Create("image/" + jpegFilename)
	if err != nil {
		return "", err
	}
	err = f.Chmod(fullAcces)
	if err != nil {
		return "", err
	}

	err = jpeg.Encode(f, m, &jpeg.Options{Quality: quality})
	if err != nil {
		return "", err
	}
	return jpegFilename, nil
}

func randomFilename16Char() (s string, err error) {
	b := make([]byte, 8)
	_, err = rand.Read(b)
	if err != nil {
		return
	}
	s = fmt.Sprintf("%x", b)
	return
}
