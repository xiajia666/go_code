package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := "hello"
	fmt.Println(strings.Contains(a, "ll"))                // true
	fmt.Println(strings.Count(a, "l"))                    // 2
	fmt.Println(strings.HasPrefix(a, "he"))               // true
	fmt.Println(strings.HasSuffix(a, "llo"))              // true
	fmt.Println(strings.Index(a, "ll"))                   // 2
	fmt.Println(strings.Join([]string{"he", "llo"}, "-")) // he-llo
	fmt.Println(strings.Repeat(a, 2))                     // hellohello
	fmt.Println(strings.Replace(a, "e", "E", -1))         // hEllo
	fmt.Println(strings.Split("a-b-c", "-"))              // [a b c]
	fmt.Println(strings.ToLower(a))                       // hello
	fmt.Println(strings.ToUpper(a))                       // HELLO
	fmt.Println(len(a))                                   // 5
	b := "你好"
	fmt.Println(len(b)) // 6

	//类型转换
	var str string = "10"
	var intValue int = 10

	isInt, _ := strconv.ParseInt(str, 0, 0) // base代表进制，bitSize解析的整数范围，0代表自动解析，决定溢出检查，最终结果会转成 int64
	fmt.Printf("%T\n", isInt)

	trans2Str_1 := strconv.Itoa(intValue)      // 数字转字符串
	trans2Str_2 := fmt.Sprintf("%d", intValue) // 数字转字符串
	fmt.Printf("%T\n", trans2Str_1)
	fmt.Printf("%T\n", trans2Str_2)

}
