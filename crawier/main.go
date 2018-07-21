package main

import (
	"muke_distributed/crawier/engine"
	"muke_distributed/crawier/save"
	"muke_distributed/crawier/schedular"
	"muke_distributed/crawier/muke/parser"
)


func main(){
	/*	engine.SimpleEngine{}.Run(engine.Request{
			Url:        "http://www.zhenai.com/zhenghun",
			ParserFunc:  parser.ParseCityList,
		})*/
	itemChan,err := save.ItemSave()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: 		&schedular.QueuedScheduler{},
		WorkerCount:	10,
		ItemChan:  		itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:        "https://www.imooc.com/course/list",
		Parser:    engine.CreateFuncParser(parser.ParseCourseList,"ParseCourseList"),
	})
}