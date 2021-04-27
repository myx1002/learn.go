package main

import "fmt"

func main() {
	typePtr()
}

/**
指针
*/
func typePtr() {
	var name string = "Go编程时光"
	//此时，name 是变量名，它只是编程语言中方便程序员编写和理解代码的一个标签。
	//当我们访问这个标签时，机算机会返回给我们它指向的内存地址里存储的值：Go编程时光。
	//出于某些需要，我们会将这个内存地址赋值给另一个变量名，通常叫做 ptr（pointer的简写），而这个变量，我们称之为指针变量。
	//换句话说，指针变量（一个标签）的值是指针，也就是内存地址
	fmt.Println(name)

	//根据变量指向的值，是否是内存地址，我把变量分为两种：
	//普通变量：存数据值本身
	//指针变量：存值的内存地址

	// 指针的创建

	// 1.先定义对应的变量，再通过变量取得内存地址，创建指针
	fmt.Println("方法一：")
	aint := 1
	ptr := &aint
	fmt.Println("定义普通变量aint：", aint)
	fmt.Println("定义指针变量ptr：", ptr)

	// 2.先创建指针，分配好内存后，再给指针指向的内存地址写入对应的值。
	fmt.Println("方法二：")
	astr := new(string)
	fmt.Println("创建一个指针astr：", astr)
	*astr = "学习Go真尼玛快乐"
	fmt.Println("给指针复制*astr：", *astr)

	// 3.先声明一个指针变量，再从其他变量取得内存地址赋值给它
	fmt.Println("方法三：")
	cint := 2
	var dint *int
	fmt.Println("声明一个指针变量dint：", dint)
	dint = &cint
	fmt.Println("把cint的内存地址赋值给dint：", dint)
	fmt.Println("读取指针变量dint的值：", *dint)

	// 上面三个方法种，指针操作都离不开这两个符号
	// &: 从一个普通变量中取得内存地址
	// *: 当*在赋值操作符（=）的右边，是从一个指针变量中取得变量值，当*在赋值操作符（=）的左边，是指该指针指向的变量
	fmt.Println("来个例子哈：")
	mint := 1     // 定义普通变量
	nptr := &mint // 定义指针变量
	fmt.Println("普通变量存储的是：", mint)
	fmt.Println("普通变量存储的是：", *nptr)
	fmt.Println("指针变量存储的是：", &mint)
	fmt.Println("指针变量存储的是：", nptr)

	// 注意点：指针和切片一样，都是引用类型
	// 如果我们想通过一个函数改变一个数组的值，有两种方法

	arr := [3]int{34, 56, 66}
	// 1.将这个数组的切片作为参数传递函数(一般使用这种，比较简洁易懂)
	modifySlice(arr[:])
	// 2.将这个数组的指针作为参数传递函数
	modifyPoint(&arr)
	fmt.Println(arr)
}

func modifySlice(slice []int) {
	slice[0] = 90
}

func modifyPoint(point *[3]int) {
	(*point)[1] = 88
}
