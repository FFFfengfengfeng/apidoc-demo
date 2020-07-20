package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var result []string
	var s string
	for i := 1; i < len(os.Args); i++ {
		result = append(result, os.Args[i])
	}
	s = strings.Join(result, ",")
	fmt.Println("参数列表: ", s)
}
