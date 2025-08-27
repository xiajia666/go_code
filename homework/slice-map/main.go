package main

import "fmt"

//// map 错误示例
//func main() {
//	//var m map[string]int
//	//m["one"] = 1 // error: panic: assignment to entry in nil map
//	m := make(map[string]int) // map 的正确声明，分配了实际的内存
//	m["one"] = 1              // error: panic: assignment to entry in nil map
//	fmt.Println(m)
//}

// slice 正确示例
//func main() {
//var s []int
//s = append(s, 1)
//}

// 错误的 key 检测方式
//func main() {
//	x := map[string]string{"one": "2", "two": "", "three": "3"}
//	if v := x["two"]; v == "" {
//		fmt.Println("key two is no entry") // 键 two 存不存在都会返回的空字符串
//	}
//}
//

//// 正确示例
//func main() {
//	x := map[string]string{"one": "2", "two": "", "three": "3"}
//	if _, ok := x["two"]; !ok {
//		fmt.Println("key two is no entry")
//	}
//}

func main() {
	mapData := map[string]string{"one": "1", "": "2", "three": "3"}
	for k, v := range mapData {
		fmt.Println(k, v)
	}
}
