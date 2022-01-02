package main

import (
	"fmt"
)

func FibonacciSequence(num int, Producer chan int) {
	x, y := 1, 1
	for i := 0; i < num; i++ {
		Producer <- x
		x, y = y, x+y
	}
	//可以显示的关闭channel, 生产者通过关键字close函数关闭channel.
	//关闭channel之后就无法在接收或者发送任何数据了，
	//应该在生产者的地方关闭channel
	//而不是由消费的地方去关闭他，这样容易引起panic.
	// 另外记住一点的就是channel 不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束 range 循环之类的。
	close(Producer)
}

func main() {
	yzj := make(chan int, 5)
	go FibonacciSequence(cap(yzj), yzj)

	//注意，这里的“value”相当对“yzj”这个channel进行读取一次数据哟。“status”的值如何为“true”则表明channel还没有被关闭哟。
	value, status := <-yzj
	fmt.Println(value, status)
	for i := range yzj {
		fmt.Println(i)
	}

	value1, status1 := <-yzj
	fmt.Println(value1, status1)

}
