package controllers

import (
	_ "encoding/json"
	"github.com/astaxie/beego"
	"zygame/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Exists() {
	message := Message{}
	name := this.GetString("uname")
	if name == "" {
		//message.Status = "fail"
		message.Message = "uname is empty"
		message.RetCode = 0
	} else {
		bol, _ := models.UserExists(name)
		if bol == true {
			//message.Status = "fail"
			message.Message = "用户名已注册"
			message.RetCode = 2
		} else {
			//message.Status = "ok"
			message.Message = "用户名正确"
			message.RetCode = 1
		}
	}
	this.Data["json"] = message
	this.ServeJson()
}

//用户注册  1：用户名和密码为必填项； 0 :注册成功， 返回用户名; 2 : 用户名已存在
func (this *UserController) Register() {
	message := Message{}
	uname := this.GetString("Uname")
	pwd := this.GetString("Pwd")
	if uname == "" || pwd == "" {
		message.Message = "用户名和密码为必填项"
		message.RetCode = 1
	} else {
		exists, tag := models.UserExists(uname)
		if tag {
			if exists {
				message.Message = "用户名已存在"
				message.RetCode = 2
			} else {
				result, _ := models.UserAdd(uname, pwd)
				if len(result) > 0 {
					message.Message = result
					message.RetCode = 0
					//fmt.Println(message)
				} else {
					message.RetCode = 999
					message.Message = "服务器内部处理错误"
				}
			}
		} else {
			message.RetCode = 999
			message.Message = "服务器内部处理错误"
		}

	}
	this.Data["json"] = message
	this.ServeJson()
}

//用户登陆 1 : 用户名和密码为必填项； 0 : 登陆成功，返回随即生成的guid； 2 : 用户名或密码不正确
func (this *UserController) Login() {
	message := Message{}
	uname := this.GetString("Uname")
	pwd := this.GetString("Pwd")
	if uname == "" || pwd == "" {
		message.Message = "用户名和密码为必填项"
		message.RetCode = 1
	} else {
		result, tag := models.UserLogin(uname, pwd)
		if tag {
			if len(result) > 0 {
				message.Message = result
				message.RetCode = 0

			} else {
				message.Message = "用户名或密码不正确"
				message.RetCode = 2
			}
		} else {
			message.Message = "服务器内部错误"
			message.RetCode = 999
		}
	}
	this.Data["json"] = message

	this.ServeJson()
}
