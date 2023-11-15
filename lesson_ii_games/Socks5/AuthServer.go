package Socks5

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func auth(reader *bufio.Reader, conn net.Conn) (error error) {
	// +----+----------+----------+
	// |VER | NMETHODS | METHODS  |
	// +----+----------+----------+
	// | 1  |    1     | 1 to 255 |
	// +----+----------+----------+
	// VER: 协议版本，socks5为0x05
	// NMETHODS: 支持认证的方法数量
	// METHODS: 对应NMETHODS，NMETHODS的值为多少，METHODS就有多少个字节。RFC预定义了一些值的含义，内容如下:
	// X’00’ NO AUTHENTICATION REQUIRED
	// X’02’ USERNAME/PASSWORD

	ver, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read ver failed :%w", err)
	}
	if ver != socks5Ver {

	}
	methodsSize, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read methodSize failed:%w", err)
	}
	method := make([]byte, methodsSize)
	// 尽可能读取数据到method中
	_, err = io.ReadFull(reader, method)
	if err != nil {
		return fmt.Errorf("read method failed : %w", err)
	}
	log.Println("ver", ver, "method", method)
	// 服务器response 包含两个字段 一个是version协议版本号 一个method ， 不需要鉴传的方式返回00
	_, err = conn.Write([]byte{socks5Ver, 0x00})
	if err != nil {
		return fmt.Errorf("wrtie failed:%w", err)
	}

	return nil
}
