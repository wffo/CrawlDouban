package parser

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"jy/douban/fetcher"
	"log"
	"strings"
	"testing"
)

func Test_DetailParser(t *testing.T) {
	var url  = "https://www.douban.com/group/topic/147838182"
	doc,err := fetcher.Fetch(url)
	if err!=nil{
		log.Println("出错了")
	}
	urlparser(doc)
}
func urlparser(resp *html.Node) {
	tt := strings.TrimSpace(htmlquery.InnerText(htmlquery.FindOne(resp,`//div[@id="content"]/h1`)))
	tt = tt[len(tt)-3:len(tt)]
	ttt := strings.TrimSpace(htmlquery.InnerText(htmlquery.FindOne(resp,`//td[@class="tablecc"]`)))[9:]
	url := htmlquery.SelectAttr(htmlquery.FindOne(resp,`//div[@class="tabs"]/a`),"href")
	url = url[:len(url)-5]
	id := url[len(url)-9:len(url)]
	images := htmlquery.Find(resp,`//div[@class="topic-richtext"]/div[@class="image-container image-float-center"]`)
	fmt.Println(len(images))
	for _,image := range images{
		image_url := htmlquery.SelectAttr(htmlquery.FindOne(image,`//div[@class="image-wrapper"]/img`),"src")
		fmt.Println(image_url)
	}
	fmt.Println(tt,ttt,url,id)
}