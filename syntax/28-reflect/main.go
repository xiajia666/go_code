package main

import (
	"fmt"
	"reflect"
)

type J struct {
	a string //小写无tag
	b string `json:"B"` //小写+tag
	C string //大写无tag
	D string `json:"DD" otherTag:"good"` //大写+tag
}

func printTag(stru interface{}) {
	t := reflect.TypeOf(stru).Elem() // 类型的反射对象
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("结构体内第%v个字段 %v 对应的json tag是 %v , 还有otherTag？ = %v \n", i+1, t.Field(i).Name, t.Field(i).Tag.Get("json"), t.Field(i).Tag.Get("otherTag"))
	}
}

func main() {

	i := 55 //int

	j := 67.8 //float64

	sum := i + int(j) //j is converted to int
	fmt.Println(sum)
}
