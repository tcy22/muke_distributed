package main

import (
	"testing"
	"muke_distributed/crawier_distributed/rpcSupport"
	"muke_distributed/crawier_distributed/worker"
	"time"
	"muke_distributed/crawier_distributed/config"
	"fmt"
)

func TestCrawlService(t *testing.T){
	const host  = ":9000"
	go rpcSupport.ServeRpc(host,worker.CrawlService{})
	time.Sleep(time.Second)

	client,err:= rpcSupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.RequestC{
		Url:"https://www.imooc.com/coursescore/947",
		Parser:worker.SerializedParser{
			Name:config.ParseProfiler,
		},
	}
	var result worker.ParseResult

	err = client.Call(config.CrawlServiceRpc,req,&result)
	if err != nil {
		t.Error(err)
	}else {
		fmt.Println(result)
	}
}
