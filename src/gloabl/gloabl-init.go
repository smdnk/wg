package gloabl

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db    *gorm.DB
	wgLog logger.Interface
)

func init() {
	wgLog = initLogger()
	db = initDB()
}

func GetDB() *gorm.DB {
	if db == nil {
		panic("数据库初始化失败,请检查")
	}
	return db
}

func GetWgLog() (wgLog logger.Interface) {
	if wgLog == nil {
		panic("全局日志未初始化,请检查")
	}
	return wgLog
}
