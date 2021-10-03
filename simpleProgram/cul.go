package main

import "fmt"

func sum() {
	for true {
		var x int
		var result int
		var max int
		fmt.Println("請輸入最大值")
		fmt.Scanln(&max)
		for x = 0; x <= max; x++ {
			result += x
		}
		fmt.Println(result)
		fmt.Println("是否繼續，請輸入yes,no")
		var contuin string
		fmt.Scanln(&contuin)
		if contuin == "no" {
			break
		}
	}

}

func main() {
	sum()
}
