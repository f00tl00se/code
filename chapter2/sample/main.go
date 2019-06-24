package main

import (
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

// init is called prior to main.
// init在main之前调用
func init() {
	// Change the device for logging to stdout.
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
// main 是整个程序的入口
func main() {
	// Perform the search for the specified term.
	// 使用特定的项做搜索
	search.Run("president")
}
