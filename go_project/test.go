package main

import "fmt"

func main() {
	s := "夏嘉xiajia666"
	fmt.Println(s[:2])         //输出乱码�
	fmt.Println([]rune(s)[:1]) //输出阿福
}
