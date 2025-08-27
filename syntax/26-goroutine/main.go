package main

import (
	"fmt"
	"sync"
)

type Producer struct {
	Name string
	Age  int
	Sex  string
}

type Consumer struct {
	Name string
	Age  int
	Sex  string
}

func ProducerGet(channelProducer chan Producer, data Producer, wg *sync.WaitGroup) {
	defer wg.Done() // defer是go中一种延迟调用机制，defer后面的函数只有在当前函数执行完毕后才能执行，通常用于释放资源。
	channelProducer <- data
	fmt.Println("生产数据", channelProducer)
}
func ConsumerGet(channelProducer chan Producer, wg *sync.WaitGroup) Producer {
	defer wg.Done()
	GeyData := <-channelProducer
	fmt.Println("获取数据", GeyData)
	return GeyData
}
func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	channelProducer := make(chan Producer, 10)
	//channelConsumer := make(chan Consumer, 10)
	data := Producer{"xiajia", 18, "man"}
	for i := 0; i < 10; i++ {
		go ProducerGet(channelProducer, data, &wg)
		go ConsumerGet(channelProducer, &wg)
	}
	wg.Wait()

}
