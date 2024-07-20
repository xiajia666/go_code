package main

import "fmt"

func main() {

	fmt.Println("aaa")
	if 1 > 0 {
		goto label
	}
	fmt.Println("bbb")
	fmt.Println("ccc")

label:
	fmt.Println("ddd")
	fmt.Println("eee")

}
