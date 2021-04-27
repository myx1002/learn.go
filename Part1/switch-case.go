package main

import "fmt"

func main() {
	// switch 后接一个你要判断变量 education （学历），然后 case 会拿这个变量去和它后面的表达式（可能是常量、变量、表达式等）进行判等。

	// 1、简单示例
	education := "本科"
	switch education {
	case "博士":
		fmt.Println("博士？牛皮牛皮!")
	case "研究生":
		fmt.Println("研究生？有点东西!")
	case "本科":
		fmt.Println("本科？可以呀!")
	case "大专":
		fmt.Println("大专？不错呀!")
	default:
		fmt.Println("对不起，你不符号我们的学历要求")
	}

	// 2、一个case多个条件
	month := 2

	switch month {
	case 3, 4, 5:
		fmt.Println("春天")
	case 6, 7, 8:
		fmt.Println("夏天")
	case 9, 10, 11:
		fmt.Println("秋天")
	case 12, 1, 2:
		fmt.Println("冬天")
	default:
		fmt.Println("输入有误...")
	}

	// 3、switch后可接函数，但是前提是要保证 case 后的值类型与函数的返回值一致
	chinese := 77
	english := 96
	math := 83
	switch getAchievement(chinese, english, math) {
	// case的数据类型要和switch的一致
	case true:
		fmt.Println("该同学所有成绩都合格")
	case false:
		fmt.Println("该同学有挂科记录")
	}

	// 4、switch不可接表达式（相当于if-elseif-else）
	score := 30
	switch {
	case score >= 95 && score <= 100:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 60:
		fmt.Println("合格")
	case score >= 0:
		fmt.Println("不合格")
	default:
		fmt.Println("输入有误...")
	}

	// 5、switch的穿透能力
	// 是当 case 使用关键字 fallthrough 的时候开启穿透能力，
	// 但是值得注意的是，fallthrough 只能穿透一层，意思是它让你直接执行下一个case的语句，而且不需要判断条件。
	s := "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s == "xxxx":
		fmt.Println("xxxx")
	case s != "world":
		fmt.Println("world")
	}
}

func getAchievement(args ...int) bool {
	for _, arg := range args {
		if arg < 60 {
			return false
		}
	}
	return true
}
