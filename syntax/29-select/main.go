package main

import (
	"fmt"
	"time"
)

// golang 中的 select 就是用来监听和 channel 有关的 IO 操作，
// 当 IO 操作发生时，触发相应的动作。
// select 只能应用于 channel 的操作，既可以用于 channel 的数据接收，也可以用于 channel 的数据发送。
// 如果 select 的多个分支都满足条件，则会随机的选取其中一个满足条件的分支执行。

func stopGoroutine(chan1 <-chan int) {
	for {
		select {
		case <-chan1:
			return
			// 如果 chan1 成功读到数据，则进行该 case 处理语句
		default:
			// 如果上面都没有成功，则进入default处理流程
			fmt.Println("未接到数据")
		}
	}

}
func main() {
	chan1 := make(chan int)
	go stopGoroutine(chan1)
	time.Sleep(1 * time.Second)
	chan1 <- 1
}
