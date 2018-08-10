package controllers

import (
	_ "encoding/json"
	"fakesoxobook/models"
	"github.com/astaxie/beego"
	"io"
	"os"
	_ "strconv"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

type ImageController struct {
	beego.Controller
}

func (this *ImageController) Post() {
	file, header, er := this.GetFile("image")
	defer file.Close()
	if er != nil {
		return
	}
	fileName := header.Filename
	fileName = strings.Replace(fileName, " ", "", -1)
	tmpFile := "./static/img/" + fileName
	newFile, er := os.Create(tmpFile)
	if er != nil {
		beego.Info("Can't create file ", er)
	}

	writenBytes, err := io.Copy(newFile, file)
	if err != nil {
		beego.Info("Can't copy file ", err)
	}
	newFile.Sync()
	newFile.Close()

	beego.Info("succes upload file, %d", writenBytes)
	image := &models.Image{fileName, tmpFile, nil}
	image, err = models.CreatePrevie(image)
	if image != nil {
		beego.Info(image)
	}
	if err != nil {
		beego.Info(err)
	}
	data, _ := models.Json(image)

	this.Ctx.Output.JSON(data, false, false)
	return
}
