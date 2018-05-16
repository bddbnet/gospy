package main

import (
	"github.com/bddbnet/gospy/config"
	"github.com/bddbnet/gospy/engine"
	"github.com/bddbnet/gospy/fetcher"
	"github.com/bddbnet/gospy/parser/h.bilibili.com"
	"github.com/bddbnet/gospy/persist"
	"github.com/bddbnet/gospy/scheduler"
)

func main() {

	//url := "https://api.vc.bilibili.com/link_draw/v2/Photo/index?type=recommend&page_num=0&page_size=45"
	//seed := engine.Request{
	//	Url:       url,
	//	ParseFunc: h_bilibili_com.CosPageList,
	//	FetchFunc: fetcher.JsonFetch,
	//}
	//engine.SimpleEngine{}.Run(seed)

	//concurrentEngine := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.SimpleScheduler{},
	//	WorkerCount: 10,
	//}
	//concurrentEngine.Run(seed)

	url := "http://127.0.0.1:9001/cos"
	seed := engine.Request{
		Url:       url,
		ParseFunc: h_bilibili_com.CosIndexList,
		FetchFunc: fetcher.Fetch,
	}

	itemChan, err := persist.ItemSaver(config.ElasticSearchIndex)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 3,
		ItemChan:    itemChan,
	}
	e.Run(seed)
}
