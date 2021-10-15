package main

import "fmt"

func main() {
	shopping := map[string]int{
		"computer": 100,
		"laptop":   200,
		"mobile":   300,
	}
	fmt.Println(shopping)
	shopping["computer"] = 200  // updating the values
	shopping["webcam"] = 400    // adding the values
	_, ok := shopping["iphone"] // searhcing the element if it exist
	fmt.Println(shopping)
	fmt.Println(ok)
	delete(shopping, "laptop") //deleting the values
	fmt.Println(shopping)
}
