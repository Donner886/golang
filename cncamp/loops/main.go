package main

import "fmt"

func deleteItemInSlice(slice []string, index int) []string {
	fixedSlice := append(slice[:index], slice[index+1:]...)
	return fixedSlice
}

func main() {

	//if语句
	var a int = 10
	if a >= 10 {
		fmt.Println("a is larger than or equal to 10")
	} else if a < 10 {
		fmt.Println("a is less than 10")

	} else {
		fmt.Println("a has no scope")
	}

	// if的简短语句。同for一样， if语句可以在条件表达式前执行
	// 一个简单的语句
	if v := a - 1; v < 0 {
		fmt.Println("v is less than 0")
	}

	//for loop
	var sum int
	// for 初始化语句; 条件语句; 修饰语句
	for i := 0; i < 100; i++ {
		sum += 1
	}
	// for 初始化语句和修饰语句可选
	for sum < 100 {
		sum += 2
	}

	// for 实现的无限循环
	for {
		if sum == 150 {
			break
		} else if sum < 150 {
			sum += 1
			continue
		}
	}

	// for-range, 遍历数组，切片，字符串，map等
	var myarr [3]string
	myarr[0] = "I LOVE GYM"
	myarr[1] = "and i go to gym two days a week"
	myarr[2] = "北京"

	myarr2 := [3]string{"I LOVE GYM", "and i go to gym two days a week", "北京"}
	myslice := myarr2[:]

	removedMySlice := deleteItemInSlice(myslice, 1)

	for index, val := range removedMySlice {
		fmt.Printf("index is %d, and value is %s \n", index, val)
		//切片 切片是对数组一个连续片段的引用。
		//数组定义中，不指定长度即为切片。 var identifier []type
		//切片在未初始化之前默认为nil,长度为0
		//fmt.Printf("myslice is %v \n", myslice)
	}

}
