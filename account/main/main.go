package main

import (
	"fmt"
	"go_code/account/utils"
)

func main() {
	var name = "xijaia"
	fmt.Println(name)
	fmt.Println("面向对象的方式来完成.....")
	utils.NewMyFamilyAccount().MainMenu()

}
