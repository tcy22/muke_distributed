package save

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"muke_distributed/crawier/engine"
	"github.com/pkg/errors"
)

func ItemSave() (chan engine.Item,error) {
	//elastic在docker，它没办法sniff
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.163.134:9200"),elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	out := make(chan engine.Item)
	go func(){
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saverr: got item" + "#%d: %v",itemCount,item)
			itemCount++

			index := "muke_data22"
			err := Save(client,index,item)
			if err != nil {
				log.Print("Item Saver:error",item,err)
			}
		}
	}()
	return out,nil
}

func Save(client *elastic.Client,index string,item engine.Item) error{
	if item.Type == ""{
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_,err := indexService.Do(context.Background())

	if err != nil {
		return err
	}
	return  nil
}
