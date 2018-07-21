package client

import (
	"muke_distributed/crawier/engine"
	"muke_distributed/crawier_distributed/rpcSupport"
	"muke_distributed/crawier_distributed/config"
	"muke_distributed/crawier_distributed/worker"
)

func CreateProcessor() (engine.Processor,error) {
	client,err := rpcSupport.NewClient(config.WorkPort0)
	if err != nil {
		return nil,err
	}
	return func(req engine.Request) (engine.ParseResult,error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		err := client.Call(config.CrawlServiceRpc,sReq,&sResult)
		if err != nil {
			return engine.ParseResult{},err
		}
		return worker.DeserializeResult(sResult),nil
	},nil
}

/*func createClient(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _,h := range hosts {
		client,err := rpcSupport.NewClient(h)
		if err == nil {
			clients = append(clients,client)
			log.Printf("connected to %s",h)
		}else {
			log.Printf("error connecting to %s",h)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for _,client := range clients {
			out <- client
		}
	}()
	return out
}*/
