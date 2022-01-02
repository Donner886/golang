package main

import (
	"fmt"
	"sync"
	"time"
)

type Quene struct {
	quene []string
	cond  *sync.Cond //
}

func (q *Quene) Enqueue(item string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.quene = append(q.quene, item)
	fmt.Printf("insert %s to quene, notify all \n", item)
	q.cond.Broadcast() //q condition broadcast
}

// 此处receiver为指针类型，才能正常使用。
func (q *Quene) Dequeue() string {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.quene) == 0 {
		fmt.Println("no data available, wait")
		q.cond.Wait()
	}
	result := q.quene[0]
	q.quene = q.quene[1:]
	return result
}

func main() {
	//initinates queue including quene slice and condition locker
	q := Quene{
		quene: []string{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	go func() {
		for {
			q.Enqueue("a")
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		q.Dequeue()
		time.Sleep(time.Second)
	}

}
