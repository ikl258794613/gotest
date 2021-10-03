package main

import "fmt"

func main() {
	var numarray [3]int
	var i int
	fmt.Println("請輸入3個數值得到平均")
	for i = 0; i < len(numarray); i++ {
		fmt.Scanln(&numarray[i])
	}
	var result int
	for i = 0; i < len(numarray); i++ {
		result += numarray[i]
	}
	var average int
	average = result / len(numarray)
	fmt.Println("平均值為：", average)
}
