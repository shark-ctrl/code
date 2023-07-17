package search

import "log"

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
