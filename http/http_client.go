package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	client := http.Client{}

	// 这里 url 必须要加 https://
	//response, err := client.Get("www.baidu.com")
	response, err := client.Get("https://www.baidu.com")
	if err != nil {
		log.Fatal("client.Get err:", err)
	}

	ct := response.Header.Get("Content-Type")
	date := response.Header.Get("Date")
	server := response.Header.Get("Server")
	statusCode := response.StatusCode
	fmt.Println("ct:", ct)                 // ct: text/html
	fmt.Println("date:", date)             // date: Tue, 20 Jun 2023 01:55:58 GMT
	fmt.Println("server:", server)         // server: BWS/1.1
	fmt.Println("statusCode:", statusCode) // 200

	body0 := response.Body
	r, err := io.ReadAll(body0)
	if err != nil {
		log.Fatal("io.ReadAll", err)
	}
	fmt.Println("body:", string(r))
}
