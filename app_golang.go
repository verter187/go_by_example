package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	bs := sha1.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	hmd := md5.New()

	hmd.Write([]byte(s))

	bsmd := md5.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bsmd)
}
