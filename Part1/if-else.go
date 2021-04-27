package main

import "fmt"

func main() {
	// 注意点：由于 Go是 强类型，所以要求你条件表达式必须严格返回布尔型的数据（nil 和 0 和 1 都不行）
	// 因为Go的true ！= 1  false != 0

	age := 88

	if age > 60 {
		fmt.Println("你是老年人了，悠着点~")
	} else if age > 30 {
		fmt.Println("中年大学，开始要秃头了吗？")
	} else {
		fmt.Println("还小，可以继续浪")
	}

	// 高级写法：在 if 里可以允许先运行一个表达式，取得变量后，再对其进行判断
	if sex := "man"; sex == "man" {
		fmt.Println("你是个boy")
	}
}
