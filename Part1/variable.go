package main

import "fmt"

func main() {
	// 第一种：单个一行声明一个变量
	var name string = "奥特曼"
	var age = 18

	// 第二种：多个一起声明
	var (
		sex            = "man"
		height float32 = 1.75
	)

	// 第三种：声明和初始化一个变量
	// 注意：该方法有个限制，只能用于函数内部
	character := "funny"

	// 第四种：声明和初始化多个变量
	weight, education := 65, "undergraduate"

	// 该方法常用于变量的交换
	a, b := 100, 200
	a, b = b, a

	fmt.Println("name: ", name)
	fmt.Println("age: ", age)
	fmt.Println("sex: ", sex)
	fmt.Println("height: ", height)
	fmt.Println("character: ", character)
	fmt.Println("weight: ", weight)
	fmt.Println("education: ", education)
	point := &education // &后面接变量名，表示取出该变量的内存地址
	fmt.Println(point)

	/**
	第五种：new 函数声明一个指针变量
	指针变量存放的时数据的地址
	使用表达式 new(Type) 将创建一个Type类型的匿名变量，初始化为Type类型的零值，然后返回变量地址，返回的指针类型为*Type。
	*/
	ptr := new(int)
	fmt.Println("ptr address: ", ptr)
	fmt.Println("ptr value: ", *ptr) // * 后面接指针变量，表示从内存地址中取出值

	/**
	特殊：匿名变量
	也称作占位符，或者空白标识符，用下划线_表示
	优点如下：
		1.不分配内存，不占用内存空间
		2.不需要你为命名无用的变量名纠结
		3.多次声明不会有任何问题
	*/
	c, _ := getData()
	_, d := getData()
	fmt.Println(c, d)
}

func getData() (int, int) {
	return 100, 200
}
