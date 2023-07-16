// main 函数保存在名为 main 的包里。如果 main 函数不在 main 包里，构建工
// 具就不会生成可执行的文件
package main

import (
	"log"
	"os"

	//-符号可以保证导入包而不使用，并且可以让matchers的rss.go完成init完成初始化
	_ "./matchers"
	"./search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
