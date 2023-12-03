package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 这里只导入，不使用
	"gw/src/gloabl"
	"gw/src/http"
	"gw/src/tableModel"
	"sync"
)

func main() {

	db := gloabl.GetDB()

	var u tableModel.User
	db.First(&u, 2)
	fmt.Printf("%v \n", u)

	var wg sync.WaitGroup
	wg.Add(1)
	http.StartWeb(&wg)

	wg.Wait()
}
