package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 这里只导入，不使用
	"gw/src/gloabl"
	"gw/src/tableModel"
)

func main() {

	db := gloabl.GetDB()

	var u tableModel.User
	db.First(&u, 2)
	fmt.Printf("%v \n", u)

}
