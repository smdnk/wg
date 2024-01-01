package main

import (
	_ "github.com/go-sql-driver/mysql" // 这里只导入，不使用
	"gw/src/demo"
)

func main() {

	demo.StartSocket()
}
