package worker

import (
	"muke_distributed/crawier/engine"
)

type CrawlService struct {}

func (CrawlService) Process(req RequestC,result *ParseResult) error {
	engineReq,err :=DeserializeRequest(req)
	if err != nil {
		return err
	}
	engineResult,err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	*result = SerializeResult(engineResult)
	return nil
}
