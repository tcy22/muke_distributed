package save

import (
	"muke_distributed/crawier/engine"
	"gopkg.in/olivere/elastic.v5"
	"muke_distributed/crawier/save"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item,result *string) error{
	err := save.Save(s.Client,s.Index,item)
	log.Printf("Item %v saved.",item)
	if err == nil {
		*result = "ok"
	}else {
		log.Printf("Error saving item %v:%v",item,err)
	}
	return err
}
