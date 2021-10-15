package main

import "fmt"

func main() {
	var a [4]int = [4]int{1, 2, 3, 4} //one way of declaring array
	b := [3]int{1, 2, 3}              // other way of declaration
	d := [...]int{1, 2, 3, 4, 5, 6}   //arry of infinite size
	c := b                            //copying one array to another
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	c[0] = 4 // changing the value of particular index
	fmt.Println(c)
	e := &d //assigning the address to e variable
	e[2] = 7
	fmt.Println(e)
	fmt.Println(d) //changes in e variable affects d variable also

}
