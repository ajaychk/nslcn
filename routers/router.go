package routers

import (
	"github.com/astaxie/beego"
	"github.com/nslcn/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
