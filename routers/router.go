package routers

import (
	"fakesoxobook/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/create_previe", &controllers.ImageController{})
}
