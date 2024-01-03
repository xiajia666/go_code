package main

import (
	"fmt"
	"model/account/utils"
)

func main() {
	var name string = "xijaia"
	fmt.Println(name)
	fmt.Println("面向对象的方式来完成.....")
	utils.NewMyFamilyAccount().MainMenu()

}
