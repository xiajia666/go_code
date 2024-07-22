package main

import "fmt"

func main() {
	// 数组的声明和初始化，默认值是0
	var a1 [3]int //数组，不指定长度就是切片，数组不能扩容
	a1[0] = 100
	a1[1] = 100
	a1[2] = 100
	fmt.Println(a1)

	a2 := [3]string{"xiajia", "xiajia666", "xiajia777"}
	fmt.Println(a2)

	var a3 = [...]int{1, 2, 3} // 数组
	fmt.Println(a3)
	fmt.Printf("%T\n", a3)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
