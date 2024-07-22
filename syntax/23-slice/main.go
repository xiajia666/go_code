package main

import "fmt"

func main() {

	s := make([]string, 3, 8) // make (类型, len, cap)
	s[0] = "a"                // 切片不能通过下标进行扩容，需要使用append
	s[1] = "b"
	s[2] = "c"
	fmt.Printf("数组：%v - 类型：%T - 长度：%v", s, s, len(s)) // c

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s) // [a b c d e f]

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c) // [a b c d e f]

	fmt.Println(s[2:5]) // [c d e]
	fmt.Println(s[:5])  // [a b c d e]
	fmt.Println(s[2:])  // [c d e f]

	good := []string{"g", "o", "o", "d"}
	fmt.Println(good) // [g o o d]

	var b = []string{1: "xiajia", 2: "xiajia666", 3: "xiajia777", 4: "xiajia888"}
	fmt.Println(b)

	for i, v := range b {
		fmt.Println(b[i])
		fmt.Println(v)
	}

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slicepart := slice[2:5]
	fmt.Println("长度：", len(slicepart), "容量：", cap(slicepart)) // 长度： 3 容量： 8
	// 长度是元素的个数，容量是首个元素到底层数组末尾元素的个数，容量cap >= 长度len

	slicepart1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slicepart = append(slicepart, slicepart1...)

	// 容量不够时会进行扩容，扩容规则是：小于1024，扩容为两倍，大于1024，扩容1/4
	sliceappend := []int{}
	for i := 1; i < 15; i++ {
		sliceappend = append(sliceappend, i)
		fmt.Printf("长度是%v,容量是%v\n", len(sliceappend), cap(sliceappend))
	}

	// 值类型：改变切片元素值，不会影响底层数组
	// 引用类型：改变切片元素值，会改变底层数组，切片是引用类型
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := sliceA
	sliceB[0] = 100
	fmt.Println(sliceA, sliceB) // [100 2 3 4 5] [100 2 3 4 5]

	sliceC := []int{1, 2, 3, 4, 5}
	sliceD := make([]int, len(sliceC), cap(sliceC))
	copy(sliceD, sliceC)
	sliceD[0] = 100
	fmt.Println(sliceC, sliceD) // [1 2 3 4 5] [100 2 3 4 5]

	// 删除切片中的元素
	sliceE := []int{1, 2, 3, 4, 5}
	sliceE = append(sliceE[:2], sliceE[3:]...) // 删除下标为2的元素，左包右不包
	fmt.Println(sliceE)

}
