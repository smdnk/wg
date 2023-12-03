package gloabl

import (
	"context"
	lo "gorm.io/gorm/logger"
	"log"
	"os"
)

func openLog() (file *os.File) {
	dir, _ := os.Getwd()
	log.Printf("当前工作目录在%v", dir)
	f, err := os.Create(dir + "\\" + "log\\" + "wg.log")
	if err != nil {
		log.Fatal("系统日志目录创建失败")
	}

	return f
}

func initLogger() (wgLog lo.Interface) {
	file := openLog()

	writer := log.New(file, "\r\n", log.LstdFlags)
	config := lo.Config{
		LogLevel:                  lo.Info,
		SlowThreshold:             200,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	}
	wgLog = lo.New(writer, config)
	wgLog.Info(context.Background(), "全局日志初始化完成----")

	return wgLog
}
