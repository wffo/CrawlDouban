package engine

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"jy/douban/fetcher"
	"jy/douban/model"
	"jy/douban/proxy"
	"jy/douban/util"
	"log"
	"time"
)

type Concurrent struct {
	Scheduler Scheduler
	Workercount int
}
type Scheduler interface {
	Submit(Request)
	ConfigurWorkerChan(chan Request)
}
type ChanProxy struct {
	ProxyChan chan string
}
func (e *Concurrent)Run(seeds...Request)  {
	go proxy.GetProxy()
	time.Sleep(time.Second)

	in :=make(chan Request)
	out:=make(chan ParserResult)
	e.Scheduler.ConfigurWorkerChan(in)

	for i :=0; i<e.Workercount; i++{
		createWorker(in,out)
	}
	for _,r:=range seeds{
		e.Scheduler.Submit(r)
	}
	for {
		result := <-out
		for _,item :=range result.Items{
			go save(item)
			log.Printf("Got URL:%s",item.(model.Profile).Url)
		}
		for _,request := range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}
func save(a_item interface{})  {
	client,err := elastic.NewClient(
		elastic.SetURL("YOURELASTICURL"),
		elastic.SetBasicAuth("USERNAME", "USERPASSWORD"),
		elastic.SetSniff(false))
	if err!= nil{
		log.Printf("%s建立服务器连接出错",a_item.(model.Profile).Url)
	}
	_,err = client.Index().Index("DATABASE").Type("TABLE").Id(a_item.(model.Profile).Id).BodyJson(a_item).Do(context.Background())
	if err!=nil{
		log.Printf("%s发送到服务器出错",a_item.(model.Profile).Url)
	}
	util.Write_append("./douban/file/url.txt",a_item.(model.Profile).Url)
	log.Printf("保存到数据库成功：%s",a_item.(model.Profile).Id)

}
func createWorker(in chan Request,out chan  ParserResult){
	go func() {
		for{
			request :=<-in
			result,err := worker(request)
			if err != nil{
				continue
			}
			out <- result
		}
	}()
}
func worker(r Request)(ParserResult,error){
	body, err := fetcher.Fetch(r.Url)
	time.Sleep(time.Millisecond*200)
	if err!=nil{
		log.Println("获取页面失败，直接跳过")
		return ParserResult{},err
	}
	parseResult := r.ParserFunc(body)
	return parseResult,nil
}