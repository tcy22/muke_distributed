package client

import (
	"muke_distributed/crawier/engine"
	"log"
	"muke_distributed/crawier_distributed/rpcSupport"
	"muke_distributed/crawier_distributed/config"
)

func ItemSave(host string) (chan engine.Item,error) {
	client,err := rpcSupport.NewClient(host)
	if err != nil {
		return nil,err
	}
	out := make(chan engine.Item)
	go func(){
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saverr: got item" + "#%d: %v",itemCount,item)
			itemCount++
            //Call rpc to save item
            result := ""
			err = client.Call(config.ItemSaverRpc,item,&result)
			if err != nil {
				log.Print("Item Saver:error",item,err)
			}
		}
	}()
	return out,nil
}