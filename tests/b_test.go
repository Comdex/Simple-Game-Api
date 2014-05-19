package zygametest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	url = "http://192.168.1.3:8080/"
)

func BenchmarkServerList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res, _ := http.Get(fmt.Sprintf("%vserver/list", url))
		defer res.Body.Close()
		ioutil.ReadAll(res.Body)
	}
}
func BenchmarkServerList1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res, _ := http.Get(fmt.Sprintf("%vserver/list", url))
		defer res.Body.Close()
		ioutil.ReadAll(res.Body)
		//fmt.Println(string(bb))
	}
}
func BenchmarkRegister(b *testing.B) {
	i := 0
	uname := "test"
	pwd := "test"
	var unamenow, pwdnow string
	for ; i < b.N; i++ {
		unamenow = fmt.Sprintf("%v%v", uname, i)
		pwdnow = fmt.Sprintf("%v%v", pwd, i)

		res, _ := http.Get(fmt.Sprintf("%vuser/signup?Uname=%v&Pwd=%v", url, unamenow, pwdnow))
		defer res.Body.Close()
		ioutil.ReadAll(res.Body)
		//fmt.Println(string(bb))

	}
}

func BenchmarkLogin(b *testing.B) {
	i := 0
	uname := "test"
	pwd := "test"
	var unamenow, pwdnow string
	for ; i < b.N; i++ {
		unamenow = fmt.Sprintf("%v%v", uname, i)
		pwdnow = fmt.Sprintf("%v%v", pwd, i)

		res, _ := http.Get(fmt.Sprintf("%vuser/signin?Uname=%v&Pwd=%v", url, unamenow, pwdnow))
		defer res.Body.Close()
		ioutil.ReadAll(res.Body)
	}
}
