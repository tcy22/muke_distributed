package main

import (
	"muke_distributed/crawier/engine"
	"muke_distributed/crawier/muke/parser"
	"muke_distributed/crawier/schedular"
	itemSaver"muke_distributed/crawier_distributed/client"
	"muke_distributed/crawier_distributed/config"
	worker "muke_distributed/crawier_distributed/worker/client"
)

func main(){
/*	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc:  parser.ParseCityList,
	})*/
	itemChan,err := itemSaver.ItemSave(config.ItemSavePort)
	if err != nil {
		panic(err)
	}

	processor,err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler: 		&schedular.QueuedScheduler{},
		WorkerCount:	10,
		ItemChan:  		itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:        "https://www.imooc.com/course/list",
		Parser:    engine.CreateFuncParser(parser.ParseCourseList,"ParseCourseList"),
	})
}