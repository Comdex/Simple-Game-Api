/*
* 功能函数工具 util/functions
 */

package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"time"
)

func CheckError(err error, comment string) bool {
	tag := true
	if err != nil {
		fmt.Println("ERROR ACCER : ", err.Error(), " OTHERS : ", comment)
		//os.Exit(1)
		tag = false
	}
	return tag
}

func Md5Encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func Timespan(str string) int64 {
	t, err := time.Parse("2006-01-02 15:04:05", str)
	CheckError(err, "Timespan(str string) parse")
	return t.Unix()
}

func GenGUID() string {
	f, _ := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}
