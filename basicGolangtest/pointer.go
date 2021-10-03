package main

import "fmt"

func main() {
	var x int
	fmt.Println("請輸入資料")
	fmt.Scanln(&x)
	fmt.Println("你輸入的資料是：", x)
	fmt.Println("你輸入的資料記憶體位置是：", &x)
	var Xponiter *int = &x
	fmt.Println("你輸入的資料記憶體位置是：", Xponiter)
	fmt.Println("反解的值為", *Xponiter)
}
