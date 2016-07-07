package routers

import (
	"dashboard_beego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/subtopo/:subtopo", &controllers.MainController{})
}
