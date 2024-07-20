package main

import (
	"fmt"
)

func main() {

	//i := 1
	//for {
	//	fmt.Println("loop")
	//	break
	//}
	//for j := 7; j < 9; j++ {
	//	fmt.Println(j)
	//}
	//
	//for n := 0; n < 5; n++ {
	//	if n%2 == 0 {
	//		07-continue
	//	}
	//	fmt.Println(n)
	//}
	//for i <= 3 {
	//	fmt.Println(i)
	//	i = i + 1
	//}

	s := "你好  golang"     // 3 + 3 + 1 + 1 + 1 .....
	for _, v := range s { // rune类型 (=int 32, 4个字节)
		fmt.Printf("%v(%c)\n", v, v) // %c该值对应的unicode码值
	}

	for i := 0; i < len(s); i++ { // byte类型字节
		fmt.Printf("%v\n", s[i])
	}

}
