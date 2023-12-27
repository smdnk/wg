package demo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ContentType    = "content-type"
	XConTypeOption = "X-Content-Type-Options"
)

const (
	Utf8 = "charset=utf-8"
)

const (
	FormData = "multipart/form-data"
	Text     = "text/plain"
	Html     = "text/html"
	Css      = "text/css"
	Js       = "text/javascript"
	Json     = "application/json"
	Xml      = "application/xml"
	Octet    = "application/octet-stream"
)

func CurrencyHead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, Json+";"+Utf8)
	// 控制浏览器是否启用嗅探，这里关闭。避免漏洞攻击
	w.Header().Set(XConTypeOption, "nosniff")
	w.WriteHeader(http.StatusAccepted)
}

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

type MyMux struct {
}

// 创建一个 handler
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
	}

	closer := r.Body // 请求体中的数据

	buffer := make([]byte, 1024)
	n, err := closer.Read(buffer)
	fmt.Println(string(buffer[0:n]))

	for k, v := range r.Form { // form表单数据
		fmt.Println("key", k)
		fmt.Println("val", v)
	}

	if r.URL.Path == "/" {
		n, err := fmt.Fprintf(w, "hello word")
		fmt.Println(n, err)
	}

	w.Header().Set(ContentType, Json)
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusAccepted)

	http.NotFound(w, r)
	return
}

func HttpStart1() {
	mux := &MyMux{} // 实例化 handler
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		return
	}
}

func HttpStart2() {
	// 创建一个 IndexHandler 实例
	mux := &MyMux{}

	// 注册处理器
	http.Handle("/", mux)
	// 启动 HTTP 服务器，监听端口 8080
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		return
	}
}
