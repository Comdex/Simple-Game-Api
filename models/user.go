/*
 *用户信息操作model
 *2014-5-12
 */

package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alphazero/Go-Redis"
	"github.com/astaxie/beego"
	_ "github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
	"time"
	"zygame/util"
)

type User struct {
	Id       int
	Username string
	Password string
}

//根据用户名判断是否已注册过  true:已注册 false:未注册
func UserExists(uname string) (bool, bool) {
	dbx := getMySQL()
	dbx.Query("set names 'utf8'")

	defer putMySQL(dbx)
	flag := true
	row, _, err := dbx.QueryFirst("SELECT id FROM %vuser WHERE username='%v'", prefix, uname)
	tag := util.CheckError(err, "Exists(uname string)bool")
	if row == nil {
		flag = false
	}
	return flag, tag
}
func UserAdd(uname string, pwd string) (string, bool) {
	dbx := getMySQL()
	defer putMySQL(dbx)

	pinsert, err := dbx.Prepare(fmt.Sprintf("INSERT INTO %vuser (username,password) VALUES(?,?)", prefix))

	tag1 := util.CheckError(err, "UserAdd(uname string,pwd string) Prepare")
	if !tag1 {
		return "", false
	}
	_, err = pinsert.Run(uname, pwd)
	tag2 := util.CheckError(err, "UserAdd(uname string,pwd string) Run")
	if !tag2 {
		return "", false
	}
	if tag1 && tag2 {
		return uname, true
	}
	return "", false
}
func UserLogin(uname, pwd string) (map[string]string, bool) {
	dbx := getMySQL()
	defer putMySQL(dbx)
	pinsert, err := dbx.Prepare(fmt.Sprintf("SELECT id,username FROM %vuser WHERE username=? AND password=?", prefix))
	tag1 := util.CheckError(err, "UserLogin(uname string,pwd string) Prepare")
	if !tag1 {
		return map[string]string{}, false
	}
	res, err := pinsert.Run(uname, pwd)
	tag2 := util.CheckError(err, "UserLogin(uname string,pwd string) Run")
	if !tag2 {

		return map[string]string{}, false
	}
	row, err := res.GetFirstRow()
	util.CheckError(err, "UserLogin(uname string,pwd string) GetFirstRow")

	tag := tag1 && tag2
	if row != nil {
		guid := util.GenGUID()
		redis_host := beego.AppConfig.String("redis_host")
		redis_port, _ := beego.AppConfig.Int("redis_port")
		err1 := _handleGUID(row.Str(1), guid, redis_host, redis_port)
		tag3 := util.CheckError(err1, "_handleGUID(uname,guid,host string,port int) handle")

		return map[string]string{"Sessionid": guid, "UserName": uname}, tag && tag3
	}

	return map[string]string{}, tag
}
func UserInfo(uname string) map[string]interface{} {
	dbx := getMySQL()
	//db := mysql.New("tcp", "", host, user, pwd, dbname)
	//err := db.Connect()
	//util.CheckError(err, "数据库连接")
	//db.Query("set names 'utf8'")
	defer putMySQL(dbx)
	pinsert, err := dbx.Prepare(fmt.Sprintf("SELECT id,username FROM %vuser WHERE username=?", prefix))
	util.CheckError(err, "UserInfo(uname string) Prepare")
	//pwd = util.Md5Encode(pwd)
	res, err := pinsert.Run(uname)
	util.CheckError(err, "UserInfo(uname string) Run")
	row, err := res.GetFirstRow()
	util.CheckError(err, "UserInfo(uname string) GetFirstRow")
	if row != nil {
		result := map[string]interface{}{"Uid": row.Int64(0), "Signtime": (time.Now().Unix()) * 1000}
		//db.Close()
		return result
	}
	return map[string]interface{}{}
}

//TODO: 往redis服务器中插入一条记录,记录格式为key(value) 其中 key 上面生成的guid value为json {uid ,uname, logintime}
func _handleGUID(uname, guid, host string, port int) error {
	spec := redis.DefaultSpec().Host(host).Port(port)
	client, e := redis.NewAsynchClientWithSpec(spec)
	if e != nil {
		return errors.New("redis server connect failed")
	}
	//根据uname获取用户信息
	info := UserInfo(uname)
	if len(info) == 0 {
		return errors.New("获取用户信息" + uname + "失败")
	}
	info["Sessionid"] = guid
	infob, _ := json.Marshal(info)
	client.Set(uname, infob)
	return nil

	//util.CheckError(e,"_handleGUID(uname,guid,host string,port int)")
}
