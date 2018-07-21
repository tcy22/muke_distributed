package engine

import (
	"log"
	"muke_distributed/crawier/fetcher"
)

func Worker(r Request)(ParseResult,error){
	log.Printf("Fetching %s",r.Url)
	body,err := fetcher.Fetcher(r.Url) //从网络上获取数据，然后由不同的解析器解析数据
	if err != nil {
		log.Printf("Fetcher:error fetching url %s,%v",r.Url,err)
		return ParseResult{},err
	}
	return r.Parser.Parser(body,r.Url),nil
}
