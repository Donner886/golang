package main

import (
	"fmt"
	"sync"
	"time"
)

//使用场景， 单一负责任务分解成多个子任务， 当所有子任务完成的时候，主任务在继续进行写去

func main() {
	// Here will create 100 sub-proccess by function waitbysleep
	waitByGroup()
}

func waitBySleep() {
	for i := 0; i < 100; i++ {
		go fmt.Println("Sub process created by waitbysleep: ", i)
	}
	time.Sleep(time.Second)
}

func waitByChannel() {
	c := make(chan bool, 100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println("Sub process created by channel", i)
			//当子线程完成执行的时候，会给channel中插入对应的状态
			c <- true
		}(i)
	}

	for i := 0; i < 100; i++ {
		//该执行步骤只有在对应位置的子线程执行结束后插入了状态的时候
		//才可以执行下一次循环，否则一直等待
		<-c
	}
}

func waitByGroup() {
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println("Sub process created by waitGroup:", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
