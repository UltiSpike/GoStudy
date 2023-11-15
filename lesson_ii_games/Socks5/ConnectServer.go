package Socks5

import (
	"bufio"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

func connect(reader *bufio.Reader, conn net.Conn) (err error) {
	// +----+-----+-------+------+----------+----------+
	// |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER 代理协议版本号，socks5的值为0x05
	// CMD 0x01表示CONNECT请求 0x02监听 0x03 关联
	// RSV 保留字段，值为0x00
	// ATYP 目标地址类型，DST.ADDR的数据对应这个字段的类型。
	//   0x01表示IPv4地址，DST.ADDR为4个字节
	//   0x03表示域名，DST.ADDR是一个可变长度的域名
	// DST.ADDR 根据目标地址类型可能不同/ 如果是host的话 第一个字节是域名字符的长度
	// DST.PORT 目标端口，固定2个字节
	buf := make([]byte, 4)
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		return fmt.Errorf("read header failed:%w", err)
	}
	ver, cmd, atyp := buf[0], buf[1], buf[3]
	if ver != socks5Ver {
		return fmt.Errorf("not support ver:%v", ver)
	}
	if cmd != cmdBind {
		return fmt.Errorf("not support cmd:%v", cmd)
	}
	addr := ""
	switch atyp {
	case atypeIPV4:
		// 读取reader中的数据 直到把buf填满
		_, err = io.ReadFull(reader, buf)
		if err != nil {
			return fmt.Errorf("read atyp failed:%w", err)
		}
		addr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])
	case atypeHOST:
		hotSize, err := reader.ReadByte()
		if err != nil {
			return fmt.Errorf("read hostSize failed:%w", err)
		}
		host := make([]byte, hotSize)
		_, err = io.ReadFull(reader, host)
		if err != nil {
			return fmt.Errorf("read host failed:%w", err)
		}
		addr = string(host)
	case atypeIPV6:
		return errors.New("ivp6 : No supported yet")
	default:
		return errors.New("invalid type")
	}
	// 读取 端口号
	_, err = io.ReadFull(reader, buf[:2])
	if err != nil {
		return fmt.Errorf("read port failed:%w", err)
	}
	// 网络字节序 默认的数据表示方式
	// 保证不同主机之间数据可以正确传输 ： 字节序一般采用大端字节序
	// 高位字节在前 ,低位字节在后(从左到右传输）
	// 数字与字节序列的转换  将端口号按大端字节序转为16进制
	port := binary.BigEndian.Uint16(buf[:2])

	dest, err := net.Dial("tcp", fmt.Sprintf("%v:%v", addr, port))
	if err != nil {
		return fmt.Errorf("dial dst failed:%w", err)
	}
	defer dest.Close()

	// 打印服务器接收到的源地址和端口地址
	log.Println("dial", addr, port)
	// +----+-----+-------+------+----------+----------+
	// |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER socks版本，这里为0x05
	// REP Relay field,内容取值如下 X’00’ succeeded
	// RSV 保留字段
	// ATYPE 地址类型 这里选择ipv4 所以填充4位
	// BND.ADDR 服务绑定的地址
	// BND.PORT 服务绑定的端口DST.PORT 两位
	//
	_, err = conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0, 0, 0, 0, 0, 0})
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	// 端口转发
	// 创建一个具有取消功能的上下文
	// 返回一个context.Context对象和一个cancel函数
	ctx, cancel := context.WithCancel(context.Background())
	// 可以在其中一个goroutine提前退出时保证程序正常返回
	defer cancel()
	go func() {
		_, _ = io.Copy(dest, reader)
		cancel()
	}()
	go func() {
		_, _ = io.Copy(conn, dest)
		cancel()
	}()
	<-ctx.Done()
	return nil
}
