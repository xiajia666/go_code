package utils

import "fmt"

func init() {
	fmt.Println("我是add函数的init")
}

func Add(x, y int) int {
	return Calc(x, y)
}
