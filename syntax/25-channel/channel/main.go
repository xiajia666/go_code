package main

import (
	"fmt"
	"time"
)

func testwrite(ch chan<- int) {
	return
}

func testread(ch <-chan int) {

}

//var ch1 chan int   //双向
//var ch2 chan<- int  //单向写
//var ch3 <-chan int  //单向读
//ch1 := make(chan int)    //基于双向通道 ch1，我们通过类型转化初始化了两个单向通道：单向只读的 ch2 和单向只写的 ch3。注意这个转化是不可逆的，
//ch2 := make(chan<- int)  //双向通道可以转化为任意类型的单向通道，但单向通道不能转化为双向通道，读写通道之间也不能相互转化。
//ch3 := make(<-chan int)

func test(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	start := time.Now()
	ch := make(chan int, 20)
	go test(ch)         //	ch := make(<-chan int, 20)
	for i := range ch { //	ch := make(chan<- int, 20)会报错
		fmt.Println("接收到的数据:", i)
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
