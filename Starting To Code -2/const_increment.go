package main

import "fmt"

const (
	a = iota
	b = iota // one way of incerement
	c = iota
)
const (
	i = iota
	j // other way of increment
	k
)
const (
	x = iota
	_ // skipping the values
	y
	z
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
