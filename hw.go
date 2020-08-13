package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

	dup2()
}

// 1.4
func dup2() {
	counts := make(map[string]int)
	var filenames string
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if counts[line] > 1 && !strings.Contains(filenames, filename) {
				filenames += filename
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Println(line, n)
		}
	}
	fmt.Println(filenames)
}
