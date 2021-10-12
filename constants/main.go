package main

import (
	"fmt"
)

func main() {
	const days int = 7

	var i int
	fmt.Println(i)

	const pi float64 = 3.14
	const secondsInHour = 3600
	fmt.Println(pi)

	duration := 234
	fmt.Printf("Duration in seconds %v\n", duration*secondsInHour)

	x, y := 5, 1
	fmt.Println(x / y)

	const a, b int = 5, 0
	// fmt.Println(a / b)

	const (
		min1 = -500
		min2 = -300
		min3 = 100
	)

	fmt.Println(min1, min2, min3)

	// Constants rules
	// 1. Can't change its value
	const temp int = 100
	// temp = 50

	// 2. Can't initialize at runtime
	// const power = math.Pow(2, 3)

	// 3. Can't use a variable to init a constant
	// t := 5
	// const tc = t

	// 4.
	const l1 = len("hello")
}
