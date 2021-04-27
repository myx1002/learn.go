package main

import "fmt"

func main() {
	// 1、接一个条件表达式
	a := 1
	for a <= 5 {
		fmt.Println(a)
		a++
	}

	// 2、接三个表达式
	// 第一个表达式：初始化控制变量，在整个循环生命周期内，只运行一次；
	// 第二个表达式：设置循环控制条件，当返回true，继续循环，返回false，结束循环
	// 第三个表达式：每次循完开始（除第一次）时，给控制变量增量或减量
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 3、不接表达式：无线循环
	// 在 Go 语言中，没有 while 循环，如果要实现无限循环，也完全可以 for 来实现。
	sort := 1
	for {
		if sort > 5 {
			break
		}
		fmt.Println(sort)
		sort++
	}

	// 4、接 for-range 语句
	// 遍历一个可迭代对象， range 后可接数组、切片，字符串等
	myarr := [...]string{"english", "chinese", "math"}
	for _, item := range myarr {
		fmt.Println(item)
	}
}
