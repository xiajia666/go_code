package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	s := "hello"
	n := 123
	p := point{1, 2}
	fmt.Println(s, n) // hello 123
	fmt.Println(p)    // {1 2}

	fmt.Printf("s=%v\n", s)  // s=hello
	fmt.Printf("n=%v\n", n)  // n=123
	fmt.Printf("p=%v\n", p)  // p={1 2}
	fmt.Printf("p=%+v\n", p) // p={x:1 y:2}
	fmt.Printf("p=%#v\n", p) // p=main.point{x:1, y:2}

	f := 3.141592653
	fmt.Println(f)          // 3.141592653
	fmt.Printf("%.2f\n", f) // 3.14

	//修改字符串
	k := "我是帅哥"        //汉字占有3字节，字母占有1字节，
	b := "abc"         // rune 是unicode编码，所以rune类型是int32类型，4个字节
	runeK := []rune(k) // byte是字节，1个字节
	byteB := []byte(b)
	runeK[0] = '你'
	byteB[0] = 'l'

	fmt.Println(string(runeK))
	fmt.Println(string(byteB))

}
