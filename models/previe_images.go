package models

import (
	"encoding/json"
	"github.com/disintegration/imaging"
	"log"
)

var newImage *Image

type Previe struct {
	Name string
	Path string
	Size string
}

type Image struct {
	Name   string
	Path   string
	Previe []*Previe
}

func CreatePrevie(image *Image) (*Image, error) {
	src, err := imaging.Open(image.Path)
	if err != nil {
		log.Fatal("cannot open image %v", err)
		return nil, err
	}
	prev := imaging.Resize(src, 150, 150, imaging.Lanczos)

	tmpPath := "./static/img/previe/150x150" + image.Name
	tmpName := "150x150" + image.Name

	err = imaging.Save(prev, tmpPath)
	if err != nil {
		log.Fatal("cannot create previe ", err)
		return nil, err
	}
	image.Previe = append(image.Previe, &Previe{tmpName, tmpPath, "150x150"})

	return image, nil
}

func Json(prev *Image) (string, error) {
	data, _ := json.Marshal(prev)
	log.Print(string(data))
	dataJson := string(data)
	return dataJson, nil
}

func NewImage() *Image {
	return &Image{}
}

func init() {
	newImage = NewImage()
}
