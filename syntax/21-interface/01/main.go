package main

import (
	"fmt"
	"unsafe"
)

type Animaler interface {
	GetName() string
	SetName(string)
}

//type Behavior interface {
//	EatFood(string)
//}

type Cat struct {
	Name string
}

func (c Cat) GetName() string {
	return c.Name
}

func (c *Cat) SetName(name string) {
	c.Name = name
}

type Dog struct {
	Name string
}

func (d *Dog) GetName() string {
	return d.Name
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func main() {
	var CatL = Cat{
		Name: "xiaoMao",
	}

	var Animal Animaler = &CatL

	fmt.Println(Animal.GetName())

	Animal.SetName("xiaoGou")
	fmt.Println(Animal.GetName())

	var Cat222 = Cat{}
	fmt.Println("----->", Cat222.GetName())

	var Dog_1 = &Dog{
		Name: "xiaoDD",
	}

	var Animal2 Animaler = Dog_1
	var Animal3 Animaler = Dog_1
	Animal3.SetName("小猪")
	fmt.Println(Animal2.GetName())

	// 空接口和类型断言使用细节
	var userInfo = make(map[string]interface{}, 1)
	//var userInfo2 = map[string]interface{}{}
	userInfo["name"] = "zhangsan"
	userInfo["age"] = 18
	userInfo["info"] = []string{"Chinese", "Boy", "Coder"}

	userInfo["dog"] = Dog_1
	st, _ := userInfo["info"].([]string) // 会返回强制转换后的对象

	fmt.Println(st)
	fmt.Println((userInfo["info"]).([]string)[0]) // 强制类型转换 x.(T)
	fmt.Println(userInfo["dog"].(*Dog))

	v := 0x1111                   // 16进制
	fmt.Println(unsafe.Sizeof(v)) // 字节数

	fmt.Println("ssssss")
}
