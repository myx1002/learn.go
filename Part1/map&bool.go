package main

import "fmt"

func main() {
	fmt.Println("字典Map：")
	typeMap()
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("布尔Bool：")
	typeBool()
}

/**
字典Map，是由若干个 key:value 这样的键值对映射组合在一起的数据结构
它是哈希表的一个实现，这就要求它的每个映射里的key，都是唯一的，可以使用 == 和 != 来进行判等操作，换句话说就是key必须是可哈希的。
可哈希的意思就是，你的 key 不能是切片，不能是字典，不能是函数。
在声明字典时，必须指定好你的key和value是什么类型的，然后使用 map 关键字来告诉Go这是一个字典。
*/
func typeMap() {
	fmt.Println("三种声明并初始化字典的方法")
	fmt.Println("第一种：")
	var scores1 map[string]int = map[string]int{"english": 80, "chinese": 85}

	fmt.Println("第二种：")
	scores2 := map[string]int{"english": 80, "chinese": 85}

	fmt.Println("第三种：")
	scores3 := make(map[string]int)
	scores3["english"] = 80
	scores3["chinese"] = 85
	fmt.Println(scores1, scores2, scores3)

	// 添加元素
	scores1["math"] = 95

	// 更新元素
	scores1["english"] = 60

	// 读取元素
	fmt.Println(scores1["math"])

	// 删除元素
	delete(scores1, "chinese")

	// 判断key是否存在
	chinese, ok := scores1["chinese"]
	if ok {
		fmt.Printf("chinese 的值是: %d", chinese)
	} else {
		fmt.Println("chinese 不存在")
	}

	// map的循环
	for subject, score := range scores1 {
		fmt.Printf("key: %s, value: %d\n", subject, score)
	}
}

/**
布尔值Bool，无非就两个值，true和false
但是在Go中，true != 1, false != 0
*/
func typeBool() {
	var (
		male1 bool = true
		male2 bool = false
	)

	fmt.Println(male1, male2)
	// 如果这样判断，会直接报错：两个类型不匹配，不能作比较
	//fmt.Println(male1 == 0)
}
