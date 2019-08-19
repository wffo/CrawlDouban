package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Readurl(filepath string)string{
	proxy := []byte{}
	f, err := os.OpenFile(filepath, os.O_RDONLY,0600)
	defer f.Close()
	if err !=nil {
		fmt.Println("打开文件出错")
	} else {
		proxy,err =ioutil.ReadAll(f)
		if err!=nil{
			fmt.Println("读取文件出错")
		}
	}
	return string(proxy)
}
