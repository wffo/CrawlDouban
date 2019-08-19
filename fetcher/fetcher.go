package fetcher

import (
	"crypto/tls"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
	"jy/douban/util"
	"log"
	"net/http"
	url2 "net/url"
	"strings"
	"time"
)

func Fetch(url string)(*html.Node,error){
	resp,err := loadurl4(url)
	if err != nil{
		return nil,errors.New("get page error:timeout")
	}
	return resp,nil
}
func loadurl4(url string) (*html.Node, error) {
	proxyip := util.Readurl("./douban/file/proxyip.txt")
	proxyip = strings.ReplaceAll(proxyip,"\r","")
	proxyip = strings.ReplaceAll(proxyip,"\n","")
	proxy1 := func(_ *http.Request) (*url2.URL, error) {
		return url2.Parse(proxyip)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy:           proxy1,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		log.Println("下载网页出现问题")
		for i:=0;i<10;i++{
			log.Printf("15秒后，重试第%d次...",i+1)
			time.Sleep(time.Second*15)
			respp,err := Retry(url)
			if err == nil{
				return respp,err
			}
			if i==9{
				return nil,errors.New("this is stubborn error")
			}
		}
	}
	if resp.StatusCode != 200{
		for i:=0;i<10;i++{
			log.Printf("返回403代码，15秒后，重试第%d次...",i+1)
			time.Sleep(time.Second*15)
			retryRes,err:= Retry(url)
			if err == nil{
				return retryRes,err
			}
			if i==9{
				return nil,errors.New("this is stubborn error")
			}
		}
	}
	r, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err!=nil{
		log.Println("网页格式转换出错")
	}
	return html.Parse(r)
}
func Retry(url string) (respp *html.Node, err error){
	log.Println("开始重试")
	defer func() {
		r := recover()
		if r!=nil{
			fmt.Println("重试又错了.......")
			respp= nil
			err= errors.New("retry error")
		}
	}()
	proxy := util.Readurl("./douban/file/proxyip.txt")
	proxy = strings.ReplaceAll(proxy,"\r","")
	proxy = strings.ReplaceAll(proxy,"\n","")
	log.Printf("代理是:%s",proxy)
	proxy1 := func(_ *http.Request) (*url2.URL, error) {
		return url2.Parse(proxy)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy:           proxy1,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		log.Println("重试又出错")
	}else {
		defer resp.Body.Close()
	}
	r, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err!=nil{
		log.Println("文本转换出错")
	}
	return html.Parse(r)

}