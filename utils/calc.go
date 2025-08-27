package utils

import "fmt"

func init() {
	fmt.Println("我是calc函数的init")
}
func Calc(x, y int) int {
	return x + y
}
