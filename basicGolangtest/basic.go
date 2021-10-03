package main

import "fmt"

func main() {
	// var x int
	// fmt.Scanln(&x)
	// fmt.Print(x)
	var x int
	var y int
	fmt.Print("輸入第一個數值")
	fmt.Scanln(&x)
	fmt.Print("輸入第二個數值")
	fmt.Scanln(&y)
	var result int = x + y
	fmt.Println(result)
}
