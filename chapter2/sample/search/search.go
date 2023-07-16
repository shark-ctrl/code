package search

import (
	"log"
	"sync"
)

// A map of registered matchers for searching.
// 以string为key mactcher为value创建一个map
var matchers = make(map[string]Matcher)

// Run performs the search logic.
func Run(searchTerm string) {
	log.Printf("run函数运行了")
	// Retrieve the list of feeds to search through.
	//检索要搜索关键字的列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// Create an unbuffered channel to receive match results to display.
	//创建一个没有缓冲区的通到，类型为我们自定义的result类型
	results := make(chan *Result)

	// Setup a wait group so we can process all the feeds.
	//设置一个等待数组，等待所有结果的获取
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait for while
	// they process the individual feeds.
	//等待数组的长度设置为每个数据源的goroutines数量的综合
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the results.
	// 第一个是索引值，第二个就是jsonList里面的json值了
	for _, feed := range feeds {
		// Retrieve a matcher for the search.
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// Launch the goroutine to perform the search.
		//启动goroutine进行搜索工作
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			//完成后递减waitGroup
			waitGroup.Done()
		}(matcher, feed)
	}

	// Launch a goroutine to monitor when all the work is done.
	go func() {
		// Wait for everything to be processed.
		//设置一个协程等到所有goroutine执行完成
		waitGroup.Wait()

		// Close the channel to signal to the Display
		// function that we can exit the program.
		close(results)
	}()

	// Start displaying results as they are available and
	// return after the final result is displayed.
	log.Println("展示结果开始")
	Display(results)
}

// Register is called to register a matcher for use by the program.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
