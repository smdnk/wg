package gloabl

import (
	"context"
	lo "gorm.io/gorm/logger"
	"log"
)

func initLogger() (wgLog lo.Interface) {
	writer := log.Logger{}
	config := lo.Config{
		LogLevel: lo.Info,
	}
	wgLog = lo.New(&writer, config)
	wgLog.Info(context.Background(), "全局日志初始化完成----")

	return wgLog
}
