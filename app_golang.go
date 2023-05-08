package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(10))

	sls := []string{"pdf", "doc", "doc", "png", "png", "txt", "doc", "txt", "pdf", "txt"}

	path := "/home/wurtow977/Development/others/testfiles1/"

	err := os.MkdirAll(filepath.Dir(path), 0766)
	if err != nil {
		fmt.Println("Error ", err.Error())
	}

	for i := 0; i < 10; i++ {

		f, err := os.Create(fmt.Sprintf("%sfile%d.%s", path, i, sls[r1.Intn(10)]))
		if err != nil {
			panic(err)
		}
		fmt.Println(f.Name())
	}
}
