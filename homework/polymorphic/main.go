package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "dog"
}

type Cat struct {
}

func (c Cat) Speak() string { // 结构体的方法
	return "cat"
}

func Speak(a Animal) string { //普通方法
	return a.Speak()
}

func main() {
	dog := Dog{}
	cat := Cat{}

	fmt.Println(dog.Speak())
	fmt.Println(cat.Speak())

	fmt.Println(Speak(dog))
	fmt.Println(Speak(cat))
}
