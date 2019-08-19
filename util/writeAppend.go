package util

import (
	"fmt"
	"log"
	"os"
)

func Write_append(filepath string,url string){
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println("打开文件出错")
	} else {
		_,err =f.Write([]byte(url))
		log.Printf("追加写入url:%s成功",url)
	}
}