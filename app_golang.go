package main

import (
	"fmt"
	"time"
)

func f(from string) {
	sum := 0
	start := time.Now()
	for i := 0; i < 10000000000; i++ {
		sum += 1
	}
	duration := time.Since(start)
	fmt.Println(from, ":", duration)
}

func main() {
	go f("поток1")
	go f("поток2")
	go f("поток3")
	go f("поток4")
	go f("поток5")
	go f("поток6")
	go f("поток7")
	go f("поток8")
	go f("поток9")
	go f("поток10")
	go f("поток11")
	go f("поток12")
	go f("поток13")
	go f("поток14")
	go f("поток15")
	go f("поток16")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(10 * time.Second)
	fmt.Println("done")
}
