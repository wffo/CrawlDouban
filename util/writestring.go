package util

import (
	"fmt"
	"log"
	"os"
)

func WriteString(filepath string,proxy string)  {
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		fmt.Println("写入出错")
	} else {
		_,err =f.Write([]byte(proxy))
		log.Printf("写入proxyip:%s成功",proxy)
	}
}
