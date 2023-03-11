package main

import "fmt"

func main() {

	// If else
	var cureentYear = 2021

	if age := cureentYear - 1998; age < 17 {
		fmt.Println("Kamu belum boleh membuat kartu sim")
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim ")
	}

	// switch
	var score = 80
	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// Switch With Relational Operators
	var scoree = 6
	switch {
	case scoree == 8:
		fmt.Println("perfect")
	case (scoree < 8) && (scoree > 3):
		fmt.Println("not bade")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}

	// Switch (fallthrough keyword)
	var x = 6
	switch {
	case x == 8:
		fmt.Println("perfect")
	case (x < 8) && (x > 3):
		fmt.Println("not bade")
		fallthrough
	case x < 5:
		fmt.Println("It is Ok")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}

	// Switch (fallthrough keyword)
	var y = 10

	if y > 7 {
		switch y {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice")
		}
	} else {
		if y == 5 {
			fmt.Println("not bad")
		} else if y == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if y == 0 {
				fmt.Println("try harder!")
			}
		}
	}

}
