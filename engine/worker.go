package engine

import "log"

func Worker(r Request) (ParseResult, error) {

	// 处理队列中url对应html文本
	body, err := r.FetchFunc(r.Url)
	//fetcher.Fetch(r.Url)

	// 文本处理错误时忽略错误
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	// 处理数据
	return r.ParseFunc(body), nil

}
