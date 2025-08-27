package main

import "fmt"

type Human struct {
	sex    string
	height float64
	width  float64
}

type Boy struct {
	name string
	Human
}

type Gril struct {
	name string
	Human
}

func (h Human) eat() {
	fmt.Println("人类都会吃饭")
}

func (b Boy) playGame() {
	fmt.Println("男孩子喜欢打游戏")
}

func (g Gril) eat() {
	fmt.Println("女孩子重新吃东西")
}

func main() {

	gril := Gril{}
	gril.eat()

}
