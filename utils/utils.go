package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/ice909/go-common/message"
)

func ReadPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("read pkg header fail, err=", err)
		return
	}

	// 获取数据长度
	var pkgLen uint32 = binary.BigEndian.Uint32(buf[0:4])

	// 根据pkgLen读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("read pkg dat fail, err = ", err)
		return
	}
	// 把buf[:pkgLen]反序列化成message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	fmt.Println("mes=", mes)
	if err != nil {
		fmt.Println("json.Unmarshal fail, err=", err)
		return
	}

	return
}

func WritePkg(conn net.Conn, data []byte) (err error) {
	// 先发送一个长度给对方
	var pkgLen uint32 = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	// 发送data本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return nil
}
