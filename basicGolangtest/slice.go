package main

import "fmt"

func main() {
	// slice := make([]int, 3, 5)
	// slice[0] = 1
	// slice[1] = 5
	// slice[2] = 3
	// fmt.Println(slice)
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))

	//宣告陣列時是用 [...]string{} 而使用切片時是用 []string{}
	// slice := []string{"小三", "小武", "阿金"}
	// fmt.Println(slice)
	// fmt.Println(len(slice)) // 長度
	// fmt.Println(cap(slice)) // 容量

	// append()
	// slice := []string{}
	// slice=append(slice,"小三")
	// slice=append(slice,"小王")
	// slice=append(slice,"Tefa")
	// fmt.Println(slice)

	// 用 for 進行走訪
	// slice := []string{"小櫻", "知世", "小可"}
	// for i := 0; i < len(slice); i = i + 1 {
	// 	fmt.Printf("%d : %s\n", i, slice[i])
	// }
	// 用 for + range 進行走訪
	slice := []string{"小櫻", "知世", "小可"}
	for k, v := range slice {
		fmt.Println("%d : %s\n", k, v)
	}

}
