/*
* 处理数据库连接的创建，并设定查询编码为utf8格式
 */

package models

import (
	"github.com/astaxie/beego"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
	"zygame/util"
)

//var db mysql.Conn
var prefix string = beego.AppConfig.String("db_prefix")
var host = beego.AppConfig.String("db_host")
var user = beego.AppConfig.String("db_user")
var pwd = beego.AppConfig.String("db_pwd")
var dbname = beego.AppConfig.String("db_name")

func init() {
	//db = mysql.New("tcp", "", host, user, pwd, dbname)
}

var MAX_POOL_SIZE = 200
var MySQLPool chan mysql.Conn

func getMySQL() mysql.Conn {

	if MySQLPool == nil {
		MySQLPool = make(chan mysql.Conn, MAX_POOL_SIZE)
	}
	if len(MySQLPool) == 0 {
		go func() {
			for i := 0; i < MAX_POOL_SIZE/2; i++ {
				mysqlc := mysql.New("tcp", "", host, user, pwd, dbname)

				err := mysqlc.Connect()
				util.CheckError(err, "getMySQL() Connect")
				mysqlc.Query("set names 'utf8'")
				putMySQL(mysqlc)
			}
		}()
	}
	temp := <-MySQLPool
	if !temp.IsConnected() {
		temp.Reconnect()
	}
	return temp
}

func putMySQL(conn mysql.Conn) {
	if MySQLPool == nil {
		MySQLPool = make(chan mysql.Conn, MAX_POOL_SIZE)
	}

	if len(MySQLPool) == MAX_POOL_SIZE {
		conn.Close()
		return
	}
	MySQLPool <- conn
}
