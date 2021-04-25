package main

import "fmt"

func main() {
	fmt.Println("数组Array：")
	typeArray()
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("切片Slice：")
	typeSlice()
}

/**
数组
是一个固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成
因为数组的长度是固定的，所以实际中很少用到，这与php的数组有很大区别
*/
func typeArray() {
	// 声明并直接初始化数组
	// 第一种方法：
	var arr1 [3]int = [3]int{1, 2, 3}

	// 第二种方法：
	arr2 := [4]int{4, 5, 6, 7}

	// 上面的3表示数组的元素个数，万一你哪天想往该数组中增加元素，
	// 你得对应修改这个数字，为了避免这种硬编码，你可以这样写，使用 ... 让Go语言自己根据实际情况来分配空间
	arr := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)
	fmt.Println("arr:", arr)

	// [3]int 和 [4]int 虽然都是数组，但他们却是不同的类型，使用 fmt 的 %T 可以查得。
	fmt.Printf("%d 的类型是：%T\n", arr1, arr1)
	fmt.Printf("%d 的类型是：%T\n", arr2, arr2)

	// 如果你觉得每次写 [3]int 有点麻烦，你可以为 [3]int 定义一个类型字面量，也就是别名类型
	type arr3 [3]int
	myarr := arr3{'a', 'b', 'c'}
	fmt.Printf("%d 的类型是：%T\n", myarr, myarr)

	// 另外，定义数组还有一种偷懒的方法，如下：
	// 可以看出[4]int{2:3}，4表示数组有4个元素，2 和 3 分别表示该数组索引为2（初始索引为0）的值为3，而其他没有指定值的，就是 int 类型的零值，即0
	arr4 := [4]int{2: 3}
	fmt.Println("arr4:", arr4)
}

/**
切片与数组一样，也是可以容纳若干类型相同的元素容器
与数组不同的是，无法通过切片类型来确定其值的长度。每个切片值都会将数组作为其底层数据结构。
我们也把这样的数组称为切片的底层数组。
*/
func typeSlice() {
	//切片是对数组的一个连续片段的引用，所以切片是一个引用类型，这个片段可以是整个数组，也可以是由起始和终止索引标识的一些项的子集，
	//需要注意的是，终止索引标识的项不包括在切片内（意思是这是个左闭右开的区间）
	myarr := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%d 的类型是 %T\n", myarr[0:2], myarr[0:2])

	// 切片的构造有四种方法

	fmt.Println("Slice声明方式1：")
	// 1.对数组进行片段截取，主要有如下两种方式
	// 从下标为1开始，下标为3结束，即[1,3)
	// 第三个参数：在切片时，若不指定第三个数，那么切片终止索引会一直到原数组的最后一个数。而如果指定了第三个数，那么切片终止索引只会到原数组的该索引值。
	mysli1 := myarr[1:3]
	mysli2 := myarr[1:3:4]

	// 切片的第三个数，影响的只是切片的容量，而不会影响长度
	fmt.Printf("myarr 的长度为：%d，容量为：%d\n", len(myarr), cap(myarr))
	fmt.Printf("mysli1 的长度为：%d，容量为：%d\n", len(mysli1), cap(mysli1))
	fmt.Printf("mysli2 的长度为：%d，容量为：%d\n", len(mysli2), cap(mysli2))

	// 2.从头声明赋值
	fmt.Println("Slice声明方式2：")
	var (
		strList      []string
		numList      []int
		numListEmpty = []int{}
	)
	fmt.Println(strList, numList, numListEmpty)

	// 3.使用make函数构造，make函数的格式：make([]Type, size, cap)
	// 这个函数刚好指出了，一个切片具备的三个要素：类型（Type），长度（size），容量（cap）
	fmt.Println("Slice声明方式3：")
	slice1 := make([]int, 2)
	slice2 := make([]int, 2, 10)
	fmt.Printf("slice1:%d 的长度为：%d，容量为：%d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2:%d 的长度为：%d，容量为：%d\n", slice2, len(slice2), cap(slice2))

	// 4.使用和数组一样，偷懒的方法
	fmt.Println("Slice声明方式4：")
	slice4 := []int{4: 2}
	fmt.Printf("slice4:%d 的长度为：%d，容量为：%d\n", slice4, len(slice4), cap(slice4))

	// 由于 切片是引用类型，所以你不对它进行赋值的话，它的零值（默认值）是 nil
	var myslices []int
	fmt.Println(myslices == nil) // true

	// 数组 与 切片 有相同点，它们都是可以容纳若干类型相同的元素的容器
	// 也有不同点，数组的容器大小固定，而切片本身是引用类型,我们可以对它 append 进行元素的添加。
	arr := []int{1}

	// 追加一个元素
	arr = append(arr, 2)

	// 追加多个元素
	arr = append(arr, 3, 4)

	// 追加一个切片, ... 表示解包，不能省略
	arr = append(arr, []int{7, 8}...)

	// 在第一个位置插入元素
	arr = append([]int{0}, arr...)

	// 在中间插入一个切片
	arr = append(arr[:5], append([]int{5, 6}, arr[5:]...)...)
	fmt.Println(arr)

	// 额外知识点：切片是数组的引用
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6:8]
	fmt.Printf("myslice为 %d, 其长度为: %d, 其容量：%d\n", myslice, len(myslice), cap(myslice))
	fmt.Println("由于myslice的容量为4，所有获取切片的切片时，获取到的值等同于[4:8]")
	myslice1 := myslice[:cap(myslice)]
	fmt.Printf("myslice的第四个元素为: %d\n", myslice1[3])

	// 切片是一个引用类型,那么是否意味着，切片其实是数组的一个引用变量，更改切片的值会影响到原数组和其它引用该数组切片的值
	myslice[1] = 999
	fmt.Println(numbers4, myslice, myslice1)

}
