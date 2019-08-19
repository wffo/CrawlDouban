package parser

import (
	"golang.org/x/net/html"
	"jy/douban/engine"
	"strconv"
)

func All_url_parser(resp *html.Node) engine.ParserResult  {
	result := engine.ParserResult{}
	for i:=0;i<40;i++{
		result.Requests = append(result.Requests,engine.Request{
			Url: "https://www.douban.com/group/baiyunzufang/discussion?start="+strconv.Itoa(25*i),
			ParserFunc: URLparser,
		})
	}
	return result
}
