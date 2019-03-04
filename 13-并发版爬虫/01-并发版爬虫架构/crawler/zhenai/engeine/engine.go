package engine

import (
	"crawler/fetcher"
	"crawler/zhenai/types"
	"log"
)

// 解析任务种子，获取更多任务种子
func Run(seeds ...types.Request) {
	var queue []types.Request //任务种子队列
	for _, seed := range seeds {
		queue = append(queue, seed)
	}
	// 不断用种子解析函数解析任务种子，同时获取更多任务种子加入队列
	for len(queue) > 0 {
		request := queue[0]
		queue = queue[1:]

		parseResult, err := worker(request)
		if err != nil {
			continue
		}

		queue = append(queue, parseResult.Requests...)	//使用...将解析结果的更多任务加入队列
		for _, item := range parseResult.Items {
			log.Printf("将目标[%s]加入了队列", item)	// 打印附加对象
		}
	}
}

// 抽取出抓取解析功能，以便并发执行
func worker(request types.Request) (types.ParseResult, error) {
	bytes, err := fetcher.ClientFetch(request.Url)
	log.Printf("抓取[%s]", request.Url)
	if err != nil {
		log.Printf("抓取[%s]页面出错:[%s]", request.Url, err)
		return types.ParseResult{}, err
	}
	return request.ParseFunc(bytes), nil
}