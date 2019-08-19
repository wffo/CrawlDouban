package proxy

import (
	"io/ioutil"
	"jy/douban/util"
	"net/http"
	"time"
)
func GetProxy()  {
	for{
		b := ""
		new_url := "XICIDAILI"
		rep,_:= http.Get(new_url)
		body,_ := ioutil.ReadAll(rep.Body)
		b = "http://" + string(body)
		util.WriteString("./douban/file/proxyip.txt",b)
		time.Sleep(time.Second*30)
	}
}
