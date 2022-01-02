package a

import (
	"fmt"

	"github.com/golang/cncamp/funcs/init/b"
)

func init() {
	fmt.Println(".........initate package a.......")
}

func Get() {
	b.Get()
	fmt.Println("this is getting function of a")
}
