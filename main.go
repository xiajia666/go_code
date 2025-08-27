package main

import (
	"fmt"
	"model/utils"
)

func init() {
	fmt.Println("我是main函数的init")
}

func main() {
	fmt.Println(utils.Calc(2, 6))
}
