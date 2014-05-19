/*
 *游戏服务器信息操作model
 *2014-5-12
 */

package models

import (
	_ "github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
	"time"
	"zygame/util"
)

type Server struct {
	Id      int
	Name    string
	Status  int
	Address string
}

func ServerList() ([]Server, bool) {

	//err := db.Connect()
	dbx := getMySQL()
	//util.CheckError(err, "数据库连接")
	//db.Query("set names 'utf8'")
	defer putMySQL(dbx)
	rows, _, err := dbx.Query("SELECT id,name,starttime,address FROM %vserver", prefix)
	tag := util.CheckError(err, "ServerList() Query")
	lists := make([]Server, 0)
	if len(rows) > 0 {
		for _, row := range rows {
			lists = append(lists, Server{row.Int(0), row.Str(1), _getStatus(row.Str(2)), row.Str(3)})
		}
	}
	return lists, tag
}

//备注：1 : 新 2 : 热
func _getStatus(str string) int {
	status := 2
	now := time.Now()
	added, err := time.Parse("2006-01-02 15:04:05", str)
	util.CheckError(err, "_getStatus(str string) Parse")
	duration := now.Sub(added)
	if duration.Hours() < 7*24 {
		status = 1
	}
	return status
}
