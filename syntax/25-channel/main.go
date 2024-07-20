package main

import (
	"fmt"
	"time"
)

// 通道写入数据和从通道接收数据都是原子操作，或者说是同步阻塞的，当我们向某个通道写入数据时，就相当于该通道被加锁，直到写入操作完成才能执行从该通道读取数据的操作，
// 反过来，当我们从某个通道读取数据时，其他协程也不能操作该通道，直到读取完成，
// 如果通道中没有数据，则会阻塞在这里，直到通道被写入数据。
// 因此，可以看到通道的发送和接收操作是互斥的，同一时间同一个进程内的所有协程对某个通道只能执行发送或接收操作，两者不可能同时进行，这样就保证了并发的安全性，数据不可能被污染。
func add(a, b int, ch chan int) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	ch <- 1
	fmt.Println("我在写入")
}

func main() {
	start := time.Now()
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go add(1, i, chs[i])
	}
	for _, ch := range chs {
		<-ch
		fmt.Println("我在读取")
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
