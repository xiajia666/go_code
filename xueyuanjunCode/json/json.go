package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string
	Website string
	Age     uint
	Male    bool
	Skills  []string
}

func main() {
	user := User{
		"学院君",
		"https://xueyuanjun.com",
		18,
		true,
		[]string{"Golang", "PHP", "C", "Java", "Python"},
	}

	u, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("JSON 编码失败：%v\n", err)
		return
	}

	fmt.Printf("JSON 格式数据：%s\n", u)
}
