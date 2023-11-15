package Socks5

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const socks5Ver = 0x05
const cmdBind = 0x01
const atypeIPV4 = 0x01
const atypeHOST = 0x03
const atypeIPV6 = 0x04

func GoTest() {
	// 创建网络监听器 从指定端口接收消息
	server, err := net.Listen("tcp", "127.0.1.1:8080")
	fmt.Printf("go test")
	if err != nil {
		panic(err)
	}
	for {
		// 创建网络连接
		client, err := server.Accept()
		if err != nil {
			log.Printf("accept failed %v", err)
			continue
		}
		go process(client)
	}
}

// 处理网络连接
func process(conn net.Conn) {
	// 闭包函数
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	// 带缓冲流的作用是可以减少系统底层调用的次数
	reader := bufio.NewReader(conn)
	// auth
	err := auth(reader, conn)
	if err != nil {
		log.Printf("Client %v auth failed :%v", conn.RemoteAddr(), err)
		return
	}
	log.Printf("auth success")
	// read addr
	err = connect(reader, conn)
	if err != nil {
		log.Printf("2client %v auth failed :%v ", conn.RemoteAddr(), err)
		return
	}

	// 读什么还什么
	//for {
	//	// 从连接中读取数据并处理
	//	b, err := reader.ReadByte()
	//	if err != nil {
	//		break
	//	}
	//	// 写回到客户端连接中
	//	_, err = conn.Write([]byte{b})
	//	if err != nil {
	//		break
	//	}
	//}
}
