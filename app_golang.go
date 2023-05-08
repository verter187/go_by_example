package main

import (
	"fmt"

	"os"
	"path/filepath"
)

func main() {
	path := "/home/wurtow977/Development/others/testfiles1/"

	err := os.MkdirAll(filepath.Dir(path), 0766)
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
	for i := 0; i < 7; i++ {

		err := os.Remove(fmt.Sprintf("%sfile%d.txt", path, i))
		if err != nil {
			panic(err)
		}

	}
}
