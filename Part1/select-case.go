package main

import (
	"fmt"
	"time"
)

/**
跟 switch-case 相比，select-case 用法比较单一，它仅能用于 信道/通道 的相关操作。
1.select 只能用于 channel 的操作(写入/读出)，而 switch 则更通用一些
2.select 的 case 是随机的，而 switch 里的 case 是顺序执行
3.select 要注意避免出现死锁，同时也可以自行实现超时机制
4.select 里没有类似 switch 里的 fallthrough 的用法
5.select 不能像 switch 一样接函数或其他表达式
*/
func main() {

	// 1、简单的例子
	simpleChannel()

	// 2、避免造成死锁
	lockChannel()

	// 3、select设置超时
	//fmt.Println("select设置超时")
	//timeOutChannel()

	// 4、读取写入都可以
	fmt.Println("select中，case写入数据")
	writeInCase()
}

/**
在运行 select 时，会遍历所有（如果有机会的话）的 case 表达式，只要有一个信道有接收到数据，那么 select 就结束，所以输出如下
*/
func simpleChannel() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	// 注意：select的随机性，可能会被c1接收，也可能会被c2接收
	c1 <- "hello1"
	c2 <- "hello"

	select {
	case msg1 := <-c1:
		fmt.Println("c1 received: ", msg1)
	case msg2 := <-c2:
		fmt.Println("c2 received: ", msg2)
	default:
		fmt.Println("No data received.")
	}
}

/**
select 在执行过程中，必须命中其中的某一分支。
但如果你没有写 default 分支，select 就会阻塞，直到有某个 case 可以命中，而如果一直没有命中，select 就会抛出 deadlock 的错误。
fatal error: all goroutines are asleep - deadlock!
*/
func lockChannel() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	// c2 <- "hello"

	select {
	case msg1 := <-c1:
		fmt.Println("c1 received: ", msg1)
	case msg2 := <-c2:
		fmt.Println("c2 received: ", msg2)
	default:
		// 为了避免fatal error: all goroutines are asleep - deadlock!
		// 尽量写个default
	}
}

/**
当 case 里的信道始终没有接收到数据时，而且也没有 default 语句时，select 整体就会阻塞
但是有时我们并不希望 select 一直阻塞下去，这时候就可以手动设置一个超时时间。
*/
func timeOutChannel() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	timeout := make(chan bool, 1)

	go makeTimeout(timeout, 2)

	select {
	case msg1 := <-c1:
		fmt.Println("c1 received: ", msg1)
	case msg2 := <-c2:
		fmt.Println("c2 received: ", msg2)
	case <-timeout:
		fmt.Println("timeout,exit")
	}
}

func makeTimeout(ch chan bool, t int) {
	time.Sleep(time.Second * time.Duration(t))
	ch <- true
}

/**
上面的例子中的case，都是从信道读取数据的，但实际上，select里的case只要求你是对信道的操作即可
不管向信道写入数据，还是从信道读出数据
*/
func writeInCase() {
	c1 := make(chan int, 2)
	c1 <- 1

	select {
	case c1 <- 4:
		fmt.Println("c1 received", <-c1)
		fmt.Println("c1 received", <-c1)
	default:
		fmt.Println("channel blocking")
	}
}
