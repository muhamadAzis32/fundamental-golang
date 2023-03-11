package main

import "fmt"

func main() {

	// Loopings (first way of looping)
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}

	// Loopings (second way of looping)
	var x = 0
	for x < 3 {
		fmt.Println("Angka ke-", x)
		x++
	}

	// Loopings (third way of looping)
	var z = 0
	for {
		fmt.Println("Angka nya", z)

		z++
		if z == 3 {
			break
		}
	}

	// Loopings (break and continue keyword)
	fmt.Printf("=========== break and continue keyword \n\n")
	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	// Loopings (nested looping)
	fmt.Printf("=========== Loopings (nested looping) \n\n")
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	// Loopings (label)
	fmt.Printf("=========== Loopings (label)\n\n")
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke - ", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}

}
