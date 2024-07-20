package main

import "fmt"

func main() {
	//语言中 break 语句用于以下几个方面:G0
	//用于循环语句中跳出循环，并开始执行循环之后的语句。
	//break 在 switch(开关语句)中在执行一条 case 后跳出语句的作用。在多重循环中，可以用标号 1abe1标出想 break 的循环。
	//表示当i=2的时候跳出当前循环
	for i := 1; i <= 10; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("继续执行")

	for i := 1; i <= 10; i++ { // 只跳出当前循环
		for j := 1; j <= 10; j++ {
			if i == 2 {
				break
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("继续执行")

lable1: // 跳出多层循环
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if i == 2 {
				break lable1
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("继续执行")
}
