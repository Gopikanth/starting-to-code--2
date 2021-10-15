package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	b := a[:]         // copies the entier array
	c := a[1:]        //starts with 1 and ends with index 3
	d := a[:3]        //ends with index 3
	e := a[:len(a)-1] //removing last element
	f := a[1:]        //removing first element
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
