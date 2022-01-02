package main

import (
	"fmt"
)

// interface is a collection of method 接口是一组方法的集合
type IF interface {
	getName() string
}

type Human struct {
	firstName, lastName string
}

type Plane struct {
	vendor string
	model  string
}

//struct 无需显式申明interface，只需要直接实现方法
// 体现多态实现
func (h *Human) getName() string {
	return h.firstName + "," + h.lastName

}

// plane实现的getname方法, 注意：此处没有使用指针参数引用
func (p Plane) getName() string {
	return fmt.Sprintf("vendor: %s, model: %s", p.vendor, p.model)
}

type Car struct {
	factory, model string
}

//这里注意， Car中getName对象的接受者是 *Car。
//函数内修改参数的值不会影响函数外原始变量的值
//可以传递指针参数将变量地址传递给给调用函数，Go语言会复制该指针作为函数的内的地址， 但是指向同一地址
func (c *Car) getName() string {
	return c.factory + "-" + c.model
}

func main() {
	//定义一个IF interface类型的切片，并且初始化为空
	interfaces := []IF{}
	//创建一个结构体
	h := new(Human)
	h.firstName = "first"
	h.lastName = "last"
	//
	interfaces = append(interfaces, h)
	c := new(Car)
	c.factory = "benz"
	c.model = "s"
	interfaces = append(interfaces, c)
	// for index, value := range 循环
	// 因为在上述过程中通过append将不同的函数体实例添加到了interfaces中，但是仅仅可以访问interface定义的方法
	for _, f := range interfaces {
		fmt.Println(f.getName())
	}

	//
	p := Plane{}
	p.vendor = "testVendor"
	p.model = "testModel"
	fmt.Println("p pointer is: ", p)
	fmt.Println(p.getName())

	p1 := &Plane{}
	fmt.Println("p pointer is: ", *p1)
	fmt.Println("h pointer is: ", &h)

}
