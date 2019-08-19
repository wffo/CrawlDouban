package main

import (
	"jy/douban/engine"
	"jy/douban/parser"
	"jy/douban/scheduler"
)

func main() {
	e := engine.Concurrent{
		Scheduler:&scheduler.SimpleScheduler{},
		Workercount:10,
	}
	e.Run(engine.Request{
		Url:"https://www.douban.com/group/baiyunzufang/discussion?start=0",
		ParserFunc: parser.All_url_parser,
	})
}
