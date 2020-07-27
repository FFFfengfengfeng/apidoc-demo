package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func minus(x, y int) int {
	return  x - y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("随机数字是:", rand.Int())
	fmt.Printf("7的平方根是: %g\n", math.Sqrt(7))
	fmt.Println("圆周率:", math.Pi)
	fmt.Println("两数之和:", add(1, 2))
	fmt.Println("x减y等于:", minus(3, 1))
	a, b := swap("123", "456")
	fmt.Printf("swap之后:%s, %s", a, b)
}
