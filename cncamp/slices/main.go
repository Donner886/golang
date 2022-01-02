package main

import "fmt"

func main() {
	demoStr := [5]string{"i", "am", "stupid", "and", "weak"}
	fmt.Printf("original demoStr is %v\n", demoStr)
	for index, value := range demoStr {
		if value == "stupid" {
			demoStr[index] = "smart"
		} else if value == "weak" {
			demoStr[index] = "strong"
		}
	}
	fmt.Printf("newly ajusted demoStr is %v\n", demoStr)
}
