package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println(os.Args[0])
	//fmt.Println(os.Args[1:])
	for index, item := range os.Args[1:] {
		fmt.Printf("下标:%d, 参数: %s\n", index, item)
	}
}
