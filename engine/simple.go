package engine

import (
	"fmt"
	"log"
)

type SimpleEngine struct {
}

func (s SimpleEngine) Run(seeds ...Request) {
	// 队列
	var requests []Request
	// 遍历seeds 添加到队列
	for _, r := range seeds {
		requests = append(requests, r)
	}

	// 循环处理队列
	for len(requests) > 0 {
		// 从队列中取出一个
		r := requests[0]
		requests = requests[1:]

		parseResult, err := Worker(r)
		if err != nil {
			fmt.Errorf("error %s", err)
			continue
		}

		// 添加到requests队列
		requests = append(requests, parseResult.Requests...)

		// 输出item
		for _, item := range parseResult.Items {
			log.Printf("Got item: %v", item)
		}
	}
}
