package main

func main() {
	//for i := 0; i < 5; i++ { // continue结束当前循环，继续下一次循环，不跳出
	//	if i == 1 {
	//		continue
	//	}
	//	println(i)
	//}

label:
	for i := 0; i < 5; i++ { // 在 continue 语句后添加标签时，表示开始标签对应的循环
		for j := 0; j < 5; j++ {
			if j == 3 {
				continue label
			}
			println(i, j)
		}
	}
}
