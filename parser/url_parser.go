package parser

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"jy/douban/engine"
	"jy/douban/util"
	"log"
	"strings"
)

func URLparser(resp *html.Node) engine.ParserResult{
	result := engine.ParserResult{}
	for _, n :=range htmlquery.Find(resp,`//table[@class="olt"]//tr[@class!="th"]`){
		aaa := htmlquery.FindOne(n, "//a")
		url := strings.TrimSpace(htmlquery.SelectAttr(aaa,"href"))
		allurl := util.Readurl("./douban/file/url.txt")
		if !strings.Contains(allurl,url){
			result.Requests = append(result.Requests,engine.Request{
				Url: url,
				ParserFunc: Detail_parser,
			})
		} else {
			log.Printf("%s已经爬过了",url)
		}
	}
	return result
}