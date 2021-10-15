package main

import "fmt"

func main() {
	var a []int = []int{10, 20, 30, 40, 50} // one way of declaration
	var b []int = make([]int, 3, 10)        // other way of Declaration
	fmt.Println(a)
	fmt.Println(len(b)) //length of slices
	fmt.Println(cap(b)) //capacity of slices
	a = append(a, 60)   //appending of slices
	fmt.Println(a)

	b = append(a, 70)
	c := append(a, b...) //spread operator
	fmt.Println(c)
}
