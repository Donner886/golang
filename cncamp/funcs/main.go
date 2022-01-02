package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/cncamp/funcs/init/a"
)

//main函数的入参
func main() {
	// parse input parameters with flag, if not containing, default value will be specified in code.
	name := flag.String("name", "world", "specify the name you want to say hi")
	flag.Parse()
	// os.Args contains all parameters of main function, including
	// function name;
	fmt.Println("os args is ", os.Args)
	fmt.Println("input parameters is", *name)

	a.Get()

}
