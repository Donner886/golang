package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go lock()
	go rLock()
	go wLock()
	time.Sleep(5 * time.Second)
}

func lock() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("lock: ", i)
	}
}

//defer语句是在方法执行结束后执行的。就是说让rLock执行结束后才会执行defer后的语句
func rLock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println("read lock: ", i)
	}
}

func wLock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("write lock: ", i)
	}
}
