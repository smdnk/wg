package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// JsonHandler 实现了 http.Handler 接口
type JsonHandler struct{}

// ServeHTTP 处理 HTTP 请求
func (h *JsonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	// 处理请求
	marshal, _ := json.Marshal("Hello, this is the index page!")
	fmt.Fprintln(w, string(marshal))
}

func StartWeb(wg *sync.WaitGroup) {
	// 创建一个 IndexHandler 实例
	jsonHandler := &JsonHandler{}

	// 注册处理器
	http.Handle("/", jsonHandler)
	// 启动 HTTP 服务器，监听端口 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start the server:", err)
		wg.Done()
	}

	wg.Done()

}
