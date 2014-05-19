package main

import (
	"github.com/astaxie/beego"
	_ "zygame/models"
	_ "zygame/routers"
)

//		Objects

//	URL					HTTP Verb				Functionality
//	/object				POST					Creating Objects
//	/object/<objectId>	GET						Retrieving Objects
//	/object/<objectId>	PUT						Updating Objects
//	/object				GET						Queries
//	/object/<objectId>	DELETE					Deleting Objects

func main() {
	//models.ServersList()
	beego.Run()
}
