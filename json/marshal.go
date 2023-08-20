package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	// `json:"-"` 表示 json 编码时忽略该字段，注意不是下划线
	// 若是 student's name 则显示的仍然是 Name，是因为单引号
	Name string `json:"student_name"`
	// `json:"age,string"` 加了之后，输出会从 Age:18 变为 age:"18"
	Age int `json:"age,string"`
	// 如果该字段为空了，那么忽略掉，前面必须有个逗号
	Score float64 `json:",omitempty"`
}

func main() {
	s := Student{Age: 20, Name: "Alice"}
	data, err := json.Marshal(&s)
	if err != nil {
		log.Fatal("marshal err: ", err)
	}
	fmt.Println(string(data))

	// 反序列化时需要一个 map 的指针
	var s2 map[string]interface{}
	// 并且 data 用反引号修饰
	s2Info := `{"Age":20,"Name":"gg","city":"Chongqing"}`
	err = json.Unmarshal([]byte(s2Info), &s2)
	if err != nil {
		log.Fatal("unmarshal s2 err: ", err)
	}
	fmt.Println("s2 =", s2)

	// map 并不一定非要是 map[string]interface{}
	var s3 map[string]int
	s3Info := `{"First":1111,"Second":2222}`
	err = json.Unmarshal([]byte(s3Info), &s3)
	if err != nil {
		log.Fatal("unmarshal s3 err: ", err)
	}
	fmt.Println("s3 =", s3)
}
