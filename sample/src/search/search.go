package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

//Run执行搜索逻辑
func Run(searchTerm string) {
	//获取数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	//创建一个无缓冲的通道，接受匹配后的结果
	results := make(chan *Result)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}
		//启动一个goroution来执行搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}
	//启动一个goroution来监控所有的工作是否都完成
	go func() {
		
		waitGroup.Wait()
		
		//关闭通道channel
		close(results)
		
		Display(results)
	}
}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}
	
	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}