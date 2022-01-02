package main

import (
	"fmt"
	"time"
)

func fibonacci(channelName, quit chan int) {
	x, y := 0, 1
	TimeOut := time.After(5 * time.Second)
	for {
		select {
		case channelName <- x: //用channelName来接收数据
			x, y = y, x+y
		//表示当接收到quit的channel时，就执行如下代码，其实就是实现关闭channel的功能
		case <-quit:
			fmt.Println("EXIT")
			return
		case <-TimeOut:
			fmt.Println("Timeout and exit now!!!! ")
			break
		}
	}
}

//单想channel
func prod(ch chan<- int) {
	for {
		ch <- 1
	}
}

func consume(ch <-chan int) {
	for {
		<-ch
	}
}

func main() {
	channelName := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 11; i++ {
			fmt.Println(<-channelName)
		}
		quit <- 100
	}()

	fibonacci(channelName, quit)
}
