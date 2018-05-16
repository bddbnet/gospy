package engine

import (
	"crypto/md5"
	"fmt"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	// 建立chan
	out := make(chan ParseResult)
	// 启动scheduler 调度器
	e.Scheduler.Run()

	// 启用count个worker
	for i := 0; i < e.WorkerCount; i++ {
		// Worker工作
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	// 使用调度器处理seeds
	for _, r := range seeds {
		//if !IsDuplicate(r.Url) {
		//}
		e.Scheduler.Submit(r)

	}

	// 接收返回的ParseResult

	for {
		result := <-out
		// 打印
		for _, item := range result.Items {
			go func(i Item) { e.ItemChan <- i }(item)
		}

		for _, request := range result.Requests {
			// 去重 url dedup
			//if !IsDuplicate(request.Url) {
			//	// 添加到调度器
			//}
			e.Scheduler.Submit(request)

		}
	}

}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {

	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				//log.Printf("error %s", err)
				continue
			}
			out <- result
		}
	}()
}

// url 去重
func IsDuplicate(url string) bool {
	//TODO:: 正式使用时候去掉下一行
	return false

	// 链接redis pool
	p := NewPool()
	r := p.Get()
	defer r.Close()

	// 生成MD5
	hash := md5.New()
	hash.Write([]byte(url))
	str := fmt.Sprintf("%x", hash.Sum(nil))

	// 检查是否存在此Url
	reply, err := r.Do("GET", str)
	if err != nil {
		log.Printf("redis error %s\n", err)
		return false
	}
	if reply != nil {
		return true
	} else {

		// 不存在则添加到redis
		reply, err = r.Do("SETEX", str, 60, 1)
		if err != nil {
			log.Printf("redis error %s\n", err)
			return false
		}

		return false
	}

}
