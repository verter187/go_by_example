package main

import (
	"fmt"
	"strconv"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) (string, int, string, bool) {
	str := strconv.Itoa(a) + "+" + strconv.Itoa(b) + "+" + strconv.Itoa(c) + "="

	return str, a + b + c, "test", true
}

func main() {
	res := plus(1, 2)
	fmt.Println(res)

	str, res, test, _ := plusPlus(1, 2, 3)
	fmt.Println(str, res, test)

}
