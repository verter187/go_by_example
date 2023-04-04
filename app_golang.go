package main

import (
	"fmt"
	"time"
)

func cat(messages chan string, str string) {
	messages <- str
}

func main() {
	msg := make(chan string, 3)

	go cat(msg, "c")
	go cat(msg, "a")
	go cat(msg, "t")

	time.Sleep(time.Second)
	test := <-msg
	test += <-msg
	test += <-msg
	fmt.Println(test)
}
