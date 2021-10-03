package main

import "fmt"

func main() {
	var x int
	fmt.Println("請輸入數字:10")
	fmt.Scanln(&x)
	if x == 10 {
		fmt.Println("good")
	} else {
		fmt.Println("not good")
	}
}
