package main

import (
	"fmt"
	"math"
)

func test() (int, int) {
	return 100, 200
}

var cc int = 1 //全局变量

func main() {

	ai, _ := test() //  _表示匿名变量  可以不使用
	fmt.Println(ai)
	var (
		name  string //默认值空
		phone int    //默认值0
	)
	qqq := 1 //局部变量
	fmt.Println(qqq)
	fmt.Println(name, phone)
	name = "ok"
	fmt.Println(name)
	fmt.Println("%T", name)  //打印变量类型
	fmt.Println("%p", &name) //打印变量地址

	var a = "initial"

	var b, c int = 1, 2

	var d = true

	var e float64

	f := float32(e)

	g := a + "foo"
	fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0
	fmt.Println(g)                // initialapple

	const s string = "constant"
	const h = 500000000
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))

	var aa int = 100
	var bb int = 200
	aa, bb = bb, aa //直接互换
	println(aa, bb)

	const (
		a_ = iota //在const出现时，每行都会使得iota值自动加1（初始值为0），如果该行没有赋值，则默认为iota的值，否则就是自己赋的值
		b_ = iota
		c_ = iota
		d_
		e_
		f_ = "www"
	)
	const (
		g_ = iota //const出现时 iota会被重置
		h_ = iota
	)

	//定义一个整数
	var xia int = 18
	println("%T,%d", xia, xia) //int    3.140000

	//定义一个浮点数
	var jia float64 = 3.14
	fmt.Printf("%T,%f", jia, jia)   //float64    3.140000
	fmt.Printf("%T,%.2f", jia, jia) //float64    3.14  保留两位小数，四舍五入

	//定义一个字符串
	var shi string = "abc" + ""
	fmt.Printf("%T,%s", shi, shi) //string abc   改用%d会打印abc对应的ascii码

	//转义字符  \
	fmt.Println("abc\"") //abc"
}
