package main

import (
	"muke_distributed/crawier_distributed/rpcSupport"
	"muke_distributed/crawier_distributed/config"
	"muke_distributed/crawier_distributed/worker"
	"flag"
)

var port = flag.Int("port",0,"the port for me to listen on")

func main() {
	rpcSupport.ServeRpc(config.WorkPort0,worker.CrawlService{})
}
