package main

import (
	"fmt"
	"time"
)

func MyEcho(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}

}

func main() {
	go MyEcho("dongdong ma - sub thread")
	MyEcho("dongdong ma - main thread")

	yzj := []int{2, 7, 1, 6, 4, 3, 11, 15, 17, 5, 8, 9, 12}
	fmt.Println("the lenght of slice yzj is: ", len(yzj))

	for i, n := range yzj {
		//定义一个匿名函数，并对该函数开启协程，每次循环都会开启一个
		//也就是说一共开启13个
		fmt.Println("index of slice is ", n)
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Println(time.Duration(n) * time.Second)
		}(i)
	}
	time.Sleep(60 * time.Second)

}
