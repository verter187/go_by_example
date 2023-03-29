package main

import "fmt"

func main() {
	fmt.Println("hello world")
	var b int = 0
	for i := 0; i < 10000000000; i++ {
		b = b + 2
	}
	fmt.Println("всего:", b)
}
