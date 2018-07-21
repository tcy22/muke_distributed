package main

import (
	"testing"
	"muke_distributed/crawier_distributed/rpcSupport"
	"muke_distributed/crawier/engine"
	"muke_distributed/crawier/model"
	"time"
)

func TestItemSave(t *testing.T)  {
	//start ItemSaveServer
	go serveRpc(":2222","test1")
	time.Sleep(time.Second)
	//start ItemSaveClient
	client,err := rpcSupport.NewClient(":2222")
	if err != nil {
		panic(err)
	}
	//Call save
	var comm []string
	comm = append(comm,"eeee")
	comm = append(comm,"kkkk")
	item := engine.Item{
		Type: "zhenai",
		Id:   "111111",
		Payload: model.Profile{
			Score:"2222",
			Num: "22",
			Comments:comm,
		},
	}
	result := ""

	err = client.Call("ItemSaverService.Save",item,&result)
	if err != nil || result != "ok" {
		t.Errorf("result:%s;err:%s",result,err)
	}
}
