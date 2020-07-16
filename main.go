package main

import "fmt"

func main() {
	m := make(map[string]int)

	v, ok := m["Abc"]

	fmt.Println("value: ", v, "是否含有该键?", ok)
}
