package main

import "fmt"

func bubble(arr *[8]int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	arr := [...]int{6, 4, 5, 3, 7, 1, 9, 8}
	bubble(&arr)
	fmt.Println(arr)
}
