package main

import (
	"fmt"
)

func MySum(a []int, sum chan int) {
	value := 0
	for _, v := range a {
		value += v
	}
	//将数据发送到channel中去
	sum <- value
}

func main() {
	//channel的声明形式为： var chanName chan ElementType
	var yzj_string chan string
	// 这是我们声明一个map, 元素是bool类型的chan
	var yzj_map map[string]chan bool

	yzj_channel := make(chan []map[string]int)

	fmt.Println(yzj_string)
	fmt.Println(yzj_map)
	fmt.Println(yzj_channel)

	yzj := []int{1, 2, 3, 4, 5, -10}
	//用chan定义一个channel对象，名称为sum
	sum := make(chan int)
	//将切片前面一半发送给channel sum
	go MySum(yzj[:len(yzj)/2], sum)
	//将切片后面一半发送给channel sum
	go MySum(yzj[len(yzj)/2:], sum)

	x, y := <-sum, <-sum
	fmt.Println("X    =", x)
	fmt.Println("Y    =", y)
	fmt.Println("X+Y  =", x+y)

}
