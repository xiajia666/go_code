package main

import (
	"fmt"
	"sync"
)

var x int64
var y int64
var wg sync.WaitGroup
var lock sync.Mutex
var RWLock sync.RWMutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()   // 加锁 ，避免多个goroutine同时进入临界区
		x = x + 1     // 临界区指的是lock.Lock()和lock.Unlock()之间的代码
		y = y - 1     //互斥锁针对的是协程，并不是针对某些变量
		lock.Unlock() //互斥锁能够保证同一时间有且只有一个goroutine进入临界区
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
	fmt.Println(y)
}
