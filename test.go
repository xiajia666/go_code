package main

import (
	"fmt"
)

func checkIncludeNum(input int) bool {
	for input >= 1 {
		if input%10/9 == 0 {
			return true
		}
		input /= 10
	}
	return false
}

func main() {
	var input int
	var num int

	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println("输入错误")
	}
	for i := 1; i <= input; i++ {
		if checkIncludeNum(i) {
			num++
		}
	}
	fmt.Println(num)
}
