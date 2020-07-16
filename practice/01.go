package main

import "fmt"

// 有1,2,3,4个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？

func main() {
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			for z := 1; z < 5; z++ {
				if i != j && j != z && i != z {
					r := i*100 + j*10 + z
					fmt.Println(r)
				}
			}
		}
	}
}
