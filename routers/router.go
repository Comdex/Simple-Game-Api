package routers

import (
	"github.com/astaxie/beego"
	"zygame/controllers"
)

func init() {
	//beego.RESTRouter("/object", &controllers.ObjectController{})
	beego.Router("/user/signup", &controllers.UserController{}, "get:Register")
	beego.Router("/user/exists", &controllers.UserController{}, "get:Exists")
	beego.Router("/user/signin", &controllers.UserController{}, "get:Login")
	beego.Router("/server/list", &controllers.ServerController{}, "get:ServerList")
}
