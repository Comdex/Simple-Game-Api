package controllers

import (
	_ "encoding/json"
	"github.com/astaxie/beego"
	"zygame/models"
)

type ServerController struct {
	beego.Controller
}

//获取服务器列表
func (this *ServerController) ServerList() {
	message := Message{}
	lists, tag := models.ServerList()
	if tag {
		message.RetCode = 0
		message.Message = lists
	} else {
		message.RetCode = 999
		message.Message = "服务器内部错误"
	}
	this.Data["json"] = message
	this.ServeJson()
}
