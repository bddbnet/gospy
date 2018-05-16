package main

import (
	"LearnGo/awe/spy2/engine"

	"LearnGo/awe/spy2/fetcher"

	"LearnGo/awe/spy2/scheduler"

	"LearnGo/awe/spy2/persist"

	"LearnGo/awe/spy2/config"

	"LearnGo/awe/spy2/parser/h.bilibili.com"
)

func main() {

	//url := "http://127.0.0.1:9001/cos"
	//seed := engine.Request{
	//	Url:       url,
	//	ParseFunc: h_bilibili_com.CosIndexList,
	//	FetchFunc: fetcher.Fetch,
	//}

	url := "https://api.vc.bilibili.com/link_draw/v2/Photo/index?type=recommend&page_num=0&page_size=45"
	seed := engine.Request{
		Url:       url,
		ParseFunc: h_bilibili_com.CosPageList,
		FetchFunc: fetcher.JsonFetch,
	}
	//engine.SimpleEngine{}.Run(seed)

	//concurrentEngine := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.SimpleScheduler{},
	//	WorkerCount: 10,
	//}
	//concurrentEngine.Run(seed)

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
