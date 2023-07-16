package main

import (
	"log"
	"os"
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
}
