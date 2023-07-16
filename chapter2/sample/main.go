package main

import (
	"log"
	"os"
	//没有用到的包可以直接用下划线注释一下
	_ "./matchers"
	"./search"
)

/*
*
相当于Java中的构造方法
*/
func init() {
	//将日志的打印结果在控制台输出
	log.SetOutput(os.Stdout)
}

func main() {
	log.Println("Hello Go in Action")
	//爬取带有pulled1的内容
	search.Run("pulled")
}
