package main

import (
	"fmt"
)

func main() {

	// if, else if, & else
	point := 5

	if point == 10 {
		fmt.Println("Perfect!")
	} else if point > 5 {
		fmt.Println("Good!")
	} else if point == 4 {
		fmt.Println("Not Bad!")
	} else {
		fmt.Printf("Bad!, your score is %d\n", point)
	}

	// temporary variable on if - else statement
	score := 88.50

	if percent := score / 100; percent >= 0.8 {
		fmt.Printf("Your point %.1f, Your grade is A\n", percent)
	} else if percent < 0.8 && percent >= 0.6 {
		fmt.Printf("Your point %.1f, Your grade is B\n", percent)
	} else if percent < 0.6 && percent >= 0.4 {
		fmt.Printf("Your point %.1f, Your grade is C\n", percent)
	} else {
		fmt.Printf("Your point %.1f, Your grade is D\n", percent)
	}

	// switch - case
	myScore := 6

	switch myScore {
	case 10:
		fmt.Println("Awesome")
	case 9:
		fmt.Println("Perfect")
	default:
		fmt.Println("Not bad")
	}

	// case with many condition
	fmt.Println("\ncase with many condition\n")
	switch myScore {
	case 10:
		fmt.Println("Awesome")
	case 9, 8, 7, 6, 5:
		fmt.Println("Perfect")
	default:
		fmt.Println("Not bad")
	}

	// case with curly braces
	fmt.Print("\ncase with curly braces\n")
	switch myScore {
	case 10:
		fmt.Println("Awesome")
	case 9, 8, 7:
		fmt.Println("Perfect")
	default:
		{
			fmt.Println("Not bad")
			fmt.Println("you can do it")
		}
	}

	// case with if else statement
	fmt.Print("\ncase with if else statement\n")
	switch myScore {
	case 10:
		fmt.Println("Awesome")
	case 9, 8, 7:
		fmt.Println("Perfect")
	default:
		if myScore == 6 {
			fmt.Println("Not bad")
		} else {
			fmt.Println("you can do it")
		}
	}

	// case with fallthrough
	fmt.Print("\ncase with fallthrough\n")
	switch {
	case myScore == 8:
		fmt.Println("Perfect")
	case (myScore < 8) && (myScore > 3):
		fmt.Println("Awesome")
		fallthrough
	case myScore < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("Not bad")
			fmt.Println("you can do it")
		}
	}

	// case with nested if
	fmt.Print("\ncase with nested if\n")
	if myScore > 7 {
		switch myScore {
		case 10:
			fmt.Println("awesome")
		default:
			fmt.Println("nice!")
		}
	} else {
		if myScore == 5 {
			fmt.Println("not bad")
		} else if myScore == 3 {
			fmt.Println("need to learn more")
		} else {
			fmt.Println("you can do it!")
			if myScore == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
