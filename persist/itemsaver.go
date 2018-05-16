package persist

import (
	"errors"
	"log"

	"context"

	"LearnGo/awe/spy2/engine"

	"fmt"

	"LearnGo/awe/spy2/config"

	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver(index string) (chan engine.Item, error) {

	// client
	client, err := elastic.NewClient(
		elastic.SetURL(config.ElasticSearchNodeUrl),
		//elastic.SetSniff(false),
	)

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 1
		for {
			item := <-out
			fmt.Printf("Item Saver: got item "+
				"#%d: %v\n", itemCount, item)
			itemCount++

			switch item.Action {
			case "index":
				err := Index(client, index, item)
				if err != nil {
					log.Printf("Item Saver: error saving item %v: %v\n",
						item, err)
				}
			case "create":
				err := Save(client, index, item)
				if err != nil {
					log.Printf("Item Saver: error saving item %v: %v\n",
						item, err)
				}
			case "update":
				err := Update(client, index, item)
				if err != nil {
					log.Printf("Item Updater: error updating item %v: %v\n",
						item, err)
				}
			}

		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	if item.Id == "" {
		return errors.New("id is null")
	}

	// 添加数据
	indexService := client.Index().
		Index(index).Type(item.Type).Routing(item.ParentId).
		Id(item.Id).BodyJson(item.Payload)

	if item.ParentId != "" {
		indexService.Parent(item.ParentId)
	}

	_, err := indexService.
		Do(context.Background())
	if err != nil {
		return err
	}

	return nil

}

func Update(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	if item.Id == "" {
		return errors.New("id is null")
	}

	// 在原有数据上更新
	updateService := client.Update().
		Index(index).Type(item.Type).Id(item.Id).Routing(item.ParentId).
		Doc(item.Payload).
		DetectNoop(true)

	if item.ParentId != "" {
		updateService.Parent(item.ParentId)
	}

	// 执行
	_, err := updateService.Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func Index(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	if isIndex("index") {
		return nil
	}

	body := `
{
  "mappings": {
    "doc": {
      "_parent": {
        "type": "user" 
      }
    },
	"doccount": {
      "_parent": {
        "type": "user" 
      }
    }
  }
}
`
	// 添加索引
	indexService := client.CreateIndex(index).
		Index(index).Body(body)

	// 执行
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}

var visitedUrls = make(map[string]bool)

func isIndex(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}
