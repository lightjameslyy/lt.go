package persist

import (
	"context"
	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"github.com/lightjameslyy/lt.go/imooclearngo/crawler/engine"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v\n", itemCount, item)
			itemCount++

			err := save(item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out
}

func save(item engine.Item) error {
	client, err := elastic.NewClient(
		// must turn off sniff in docker
		elastic.SetSniff(false),
		elastic.SetURL("http://localhost:9200"))
	if err != nil {
		return err
	}

	if item.Type == "" {
		return errors.New("Must supply Type")
	}

	indexService := client.Index().Index("dating_profile").Type(item.Type).Id(item.Id).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.Do(context.Background())
	if err != nil {
		return err
	}

	return nil

}
