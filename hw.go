package main

import (
	"fmt"
	"os"
)

func main() {
	// 1.1
	fmt.Println(os.Args)
	// 1.2
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
	// 1.3
	// ...
}
