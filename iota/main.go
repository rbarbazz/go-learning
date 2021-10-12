package main

import "fmt"

func main() {
	const (
		a1 = iota * 2
		a2
		a3
		a4
	)

	fmt.Println(a1, a2, a3, a4)

	const (
		b1 = -(iota + 2)
		_
		b3
		b4
	)

	fmt.Println(b1, b3, b4)
}
