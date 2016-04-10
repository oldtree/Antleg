package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Environ() {
		fmt.Println(index, " : ", value)
	}

	fmt.Println("init commit")
}
