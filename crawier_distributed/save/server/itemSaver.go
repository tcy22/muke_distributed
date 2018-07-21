package main

import (
	"muke_distributed/crawier_distributed/rpcSupport"
	"gopkg.in/olivere/elastic.v5"
	"muke_distributed/crawier_distributed/save"
	"muke_distributed/crawier_distributed/config"
	"flag"
)

var(
	itemSaverHost = flag.String(
		"itemSaver_host","","itemSaver host")
	workerHosts = flag.String(
		"worker_hosts","","worker hosts (comma separated)")
)

func main()  {
	err := serveRpc(config.ItemSavePort,config.ElasticIndex)
	if err != nil {
		panic(err)
	}
}

func serveRpc(host,index string) error{
	client, err := elastic.NewClient(
		elastic.SetURL(config.ElasticUrl),elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return rpcSupport.ServeRpc(host,
		&save.ItemSaverService{
			Client:client,
			Index:index,
		})
}