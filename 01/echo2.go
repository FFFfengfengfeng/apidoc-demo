package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var result []string
	for _, item := range os.Args[1:] {
		result = append(result, item)
	}
	fmt.Println("参数列表: ", strings.Join(result, ","))
}
