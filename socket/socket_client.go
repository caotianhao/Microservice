package main

import (
	"fmt"
	"log"
	"net"
)

// 先启动 server，再新开一个控制台启动 client
func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dial success")

	var sendData string
	_, _ = fmt.Scanf("%s", &sendData)
	//fmt.Println("ss", sendData)

	cnt, err := conn.Write([]byte(sendData))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("client ====> server 长度：", cnt, "内容：", sendData)

	// 接收服务器返回的数据的容器
	buf := make([]byte, 1024)
	cnt, err = conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server ====> client 长度：", cnt, "内容：", string(buf[:cnt]))

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)
}
