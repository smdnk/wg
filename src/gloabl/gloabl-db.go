package gloabl

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

const (
	dsn = "root:123456@tcp(192.168.0.6:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
)

func initDB() (db *gorm.DB) {
	// 创建 database/sql 连接
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	// 设置数据库连接池
	sqlDB.SetMaxIdleConns(10)                 // 设置空闲连接数
	sqlDB.SetMaxOpenConns(100)                // 设置最大连接数
	sqlDB.SetConnMaxLifetime(time.Minute * 3) // 设置最大空闲时间

	// 使用 gorm.Open 连接到数据库
	db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		wgLog.Error(context.TODO(), "Failed to open database:", err)
		return
	}
	wgLog.Info(context.TODO(), "全局数据库初始化完成")

	return db
}
