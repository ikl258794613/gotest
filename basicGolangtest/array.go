package main

import "fmt"

func main() {
	var arr1 [5]int = [5]int{2, 2, 1, 0, 5}
	var arr2 [2]string = [2]string{"hello", "world"}
	var arr3 [3]bool = [3]bool{true, false, true}
	fmt.Println("arr1全部資料是:", arr1)
	fmt.Println("arr1的index1資料是:", arr1[1])
	fmt.Println("arr2的index0資料是:", arr2[0])
	fmt.Println("arr3的index2資料是:", arr3[2])
	//取得陣列長度len()
	fmt.Println("arr1的陣列長度:", len(arr1))

	var i int
	for i = 0; i < len(arr1); i++ {
		fmt.Println(`arr1的index`, i, ":", arr1[i])
	}
}
