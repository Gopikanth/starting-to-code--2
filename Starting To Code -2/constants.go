package main

import "fmt"

const New_number = 15 //global Declaration
const (
	New_number1 = 16
	New_number2 = 17 // Declaring Blocks of constant
	New_number3 = 18
)

func main() {
	const First_number = 12  //constant which can exported to other packages
	const second_number = 13 //constant within the function
	fmt.Println(First_number + second_number)
	var a = 14
	fmt.Println(a + First_number) // operation of constant and variables
	fmt.Println(New_number1 - New_number2 - New_number3)
}
