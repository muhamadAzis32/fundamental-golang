package main

import (
	"fmt"
)

func main() {
	// 1
	var a uint8 = 10
	var b byte

	b = a
	_ = b

	// 2
	type second = uint

	var hour second = 3600
	fmt.Printf("hour type: %T\n", hour)
}
