package search

import (
	"log"
	"sync"
)

// 创建一个存储匹配器的map，key为字符串，value为匹配器对象
var matcherMap = make(map[string]Matcher)

// Register 将匹配器存到map中，如果存在则直接报错
func Register(feedType string, matcher Matcher) {
	if _, exists := matcherMap[feedType]; exists {
		log.Fatalln("当前matcherMap已存在匹配器，key", feedType)
	}

	log.Println("注册匹配器", feedType)
	matcherMap[feedType] = matcher
}

func Run(searchItem string) {
	log.Println("开始搜索匹配项，匹配关键字:", searchItem)

	//获取所有的文章列表
	feeds, err := RetrieveFeeds()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("文章列表查询结束，文章长度:", len(feeds))

	//使用waitGroup进行协程流程控制，后续我们会为每个feeds开一个协程，所以waitGroup计数器个数就为feeds个数
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	var result = make(chan *Result)

	for _, feed := range feeds {
		matcher, exist := matcherMap[feed.Type]
		//如果不存在则取默认
		if !exist {
			matcher = matcherMap["default"]
		}

		//启动一个闭包函数的协程
		go func(matcher Matcher, feed *Feed) {
			log.Println("开启一个协程进行工作,请求URL:", feed.URI)
			Match(matcher, feed, searchItem, result)
			log.Println("请求url:", feed.URI, "工作完毕")
			//每一个协程完成工作后，waitGroup减1
			waitGroup.Done()
		}(matcher, feed)

	}

	go func() {
		log.Println("waitGroup等待所有协程工作完毕")
		//等待所有协程执行完成
		waitGroup.Wait()
		log.Println("所有协程工作完毕,关闭waitGroup")

		//将协程同意使用的通道关闭
		close(result)

	}()
	log.Println("输出结果:")
	Display(result)
}
