package main

import "fmt"

type Point struct {
	x int
	y int
}
type Person struct {
	name string
	age  int
}

func main() {
	var p1 Point = Point{3, 4}
	fmt.Println("Point的x欄位是", p1.x)
	fmt.Println("Point的y欄位是", p1.y)

	var people1 Person = Person{"老王", 30}
	var people2 Person = Person{age: 20, name: "alice"}
	fmt.Println(people1)
	fmt.Println(people2)
	people2.name = "Else"
	fmt.Println(people2)
}
