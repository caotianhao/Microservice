package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("listening......")
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("conn success!")

	// 创建容器，接收读取到的数据
	buf := make([]byte, 1024)
	// cnt 真正从 client 读取到数据的长度
	cnt, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	data := string(buf[:cnt])
	fmt.Println("client ====> server 长度：", cnt, "内容：", data)

	// server 对 client 数据处理，转大写
	upperData := strings.ToUpper(data)

	cnt, err = conn.Write([]byte(upperData))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server ====> client 长度：", cnt, "内容：", upperData)

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)
}
