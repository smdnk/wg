package demo

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// 获取当前应用工作目录（mod文件所在目录）
var workDir, _ = os.Getwd()

// 获取当前系统文件路径分隔符
var separator = string(filepath.Separator)

// FileReadDemo 读取文件
func FileReadDemo() {

	// 使用当前系统文件路径分隔符拼接路径
	filePath := filepath.Join(workDir, "config", "db-config", "db-config.yml", separator)

	fmt.Println(filePath)
	file, _ := os.Open(filePath)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text() // 一次读取一行
		fmt.Println(text)
	}
}

// FileWriteDemo 写入文件
func FileWriteDemo() {
	// 使用当前系统文件路径分隔符拼接路径
	filePath := filepath.Join(workDir, "config", "db-config", "db-config-back.yml", separator)
	file, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString("hello word")
	if err != nil {
		return
	}
}
