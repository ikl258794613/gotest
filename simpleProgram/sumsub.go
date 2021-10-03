package main

import "fmt"

func sumandsub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	fmt.Println(sum, sub)
	return sum, sub
}

func main() {
	var a int
	var b int
	fmt.Println("請輸入數字a")
	fmt.Scanln(&a)
	fmt.Println("請輸入數字b")
	fmt.Scanln(&b)
	sumandsub(a, b)
}
