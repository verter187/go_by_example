package main

import (
	"fmt"
	"strconv"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int, res string) int {
	res = strconv.Itoa(a) + "+" + strconv.Itoa(b) + "+" + strconv.Itoa(c) + "="
	fmt.Println(res)
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println(res)

	var str string
	res = plusPlus(1, 2, 3, str)
	fmt.Println(str, res)
}
