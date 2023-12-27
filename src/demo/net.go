package demo

import (
	"fmt"
	"net"
)

func Server() {

	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("服务器启动监听失败,程序结束", err)
		return
	}

	for { // 循环等待连接
		client, err := listen.Accept() // 阻塞 等待连接
		if err != nil {
			fmt.Println("服务器监听客户端连接失败,断开此客户连接", err)
			return
		}
		fmt.Printf("客户端连接成功 %v，开始数据交互。。。。 \n", client.RemoteAddr().String())

		// 有连接进来后分配一个后台进程处理
		go func(client net.Conn) {
			buffer := make([]byte, 1024)
			for {
				n, err := client.Read(buffer) // 阻塞 等待数据发送
				if err != nil {
					fmt.Println("读取客户端数据失败")
					return
				}
				fmt.Printf("读取客户端数据为 %v \n", string(buffer[0:n]))
			}
		}(client)

	}

}

func Client() {
	fmt.Println("客户端启动")
	connect, err := net.Dial("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Println("客户端连接服务器失败！退出")
		return
	}
	fmt.Println("客户端连接服务器成功,", connect.LocalAddr().String())

	n, err := connect.Write([]byte("1"))

	if err != nil {
		fmt.Println("客户端发送数据到服务器失败！退出")
		return
	}
	fmt.Printf("客户端发送数据成功,数据长度%d \n", n)

}
