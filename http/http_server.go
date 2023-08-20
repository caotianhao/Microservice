package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// 注册路由
	// xxx/name ===> func1
	// xxx/age  ===> func2
	// xxx/score ===> func3
	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("request:", request)
		_, err := io.WriteString(writer, "this is name write")
		if err != nil {
			log.Fatal("io.WriteString name err: ", err)
		}
	})

	http.HandleFunc("/age", func(writer http.ResponseWriter, request *http.Request) {
		_, err := io.WriteString(writer, "this is age write")
		if err != nil {
			log.Fatal("io.WriteString age err: ", err)
		}
	})

	http.HandleFunc("/score", func(writer http.ResponseWriter, request *http.Request) {
		_, err := io.WriteString(writer, "this is score write")
		if err != nil {
			log.Fatal("io.WriteString score err: ", err)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("http.ListenAndServe err:", err)
	}
	fmt.Println("http server start......")
}
