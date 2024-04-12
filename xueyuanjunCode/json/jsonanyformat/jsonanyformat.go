package main

import (
	"encoding/json"
	"fmt"
)

type Sayer interface {
	say()
}

type dog struct {
}
type cat struct {
}

// dog实现了Sayer接口

func (d dog) say() {
	fmt.Println("汪汪汪")
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func main() {
	var x Sayer
	a := cat{}
	x = a
	x.say()
	u3 := []byte(`{"name": "学院君", "website": "https://xueyuanjun.com", "age": 18, "male": true, "skills": ["Golang", "PHP"]}`)
	var user4 interface{}
	fmt.Printf("%T\n", &user4)
	err := json.Unmarshal(u3, &user4)
	if err != nil {
		fmt.Printf("JSON 解码失败：%v\n", err)
		return
	}
	fmt.Printf("JSON 解码结果: %#v\n", user4) //user4 被定义为一个空接口。json.Unmarshal() 函数将一个 JSON 对象 u3 解码到空接口 user4 中，最终 user4 将会是一个键值对的 map[string]interface{} 结构
	fmt.Printf("%T\n", &user4)

}
