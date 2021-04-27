package main

import "fmt"

var typeName string = "go"

func main() {
	// 1、延迟调用
	// defer 的用法很简单，只要在后面跟一个函数的调用，就能实现将这个 xxx 函数的调用延迟到当前函数执行完后再执行。
	deferOne()

	// 2、即时求值的变量快照
	// 使用 defer 只是延时调用函数，此时传递给函数里的变量，不应该受到后续程序的影响。
	fmt.Println("defer 即时求值的变量快照")
	deferTwo1()
	// 如果 defer 后面跟的是匿名函数，情况会有所不同， defer 会取到最后的变量值
	fmt.Println("defer 后面跟的是匿名函数")
	deferTwo2()

	// 3、多个defer 反序调用
	fmt.Println("多个defer 反序调用")
	deferThree()

	// 4、defer 与 return 孰先孰后
	myName := deferReturn()
	fmt.Println("main 函数里的typeName: ", typeName)
	// defer 是return 后才调用的。所以在执行 defer 前，myName 已经被赋值成 go 了
	fmt.Println("main 函数里的myName: ", myName)

	// 要点：defer一般用于资源释放、关闭文件等等

}

func deferOne() {
	defer deferFunc()
	fmt.Println("正常执行")
}

func deferTwo2() {
	name := "go"
	defer func() {
		fmt.Println(name)
	}()

	name = "php"
	fmt.Println(name)
}

func deferTwo1() {
	name := "go"
	defer fmt.Println(name) // 输出 go

	name = "php"
	fmt.Println(name)
}

func deferThree() {
	name := "go"
	defer fmt.Println("第一个defer")

	name = "php"
	defer fmt.Println("第二个defer")

	fmt.Println(name)
}

func deferReturn() string {
	defer func() { typeName = "php" }()

	fmt.Println("deferReturn函数里面的name", typeName)
	return typeName
}

func deferFunc() {
	fmt.Println("defer延迟执行")
}
