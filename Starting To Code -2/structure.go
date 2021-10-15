package main

import "fmt"

type student struct {
	name     string
	roll     int
	subjects []string
}

func main() {
	student1 := student{
		name: "gopikanth",
		roll: 815117104017,
		subjects: []string{
			"M1",
			"M2",
			"M3",
		},
	}
	fmt.Println(student1)
	fmt.Println(student1.name) //fetching a particular value
	student1.name = "gopi"     // changing the value
	fmt.Println(student1)
}
