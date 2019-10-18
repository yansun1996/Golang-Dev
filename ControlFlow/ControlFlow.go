package main

import (
	"fmt"
)

func main() {
	if false {
		fmt.Println("The test is false!")
	}
	if true {
		fmt.Println("The test is false!")
	}

	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
	}

	if pop, ok := statePopulations["California"]; ok {
		// pop only exist in this block
		fmt.Println(pop)
	}

	number := 30
	guess := 50
	if guess < number {
		fmt.Println("Too low")
	} else if guess > number {
		fmt.Println("Too high")
	} else {
		fmt.Println("You got it !")
	}
	fmt.Println(number <= guess, number >= guess, number != guess)

	if guess < 1 || guess > 100 {
		fmt.Println("Guess number must be between 1 and 100 !")
	}
	fmt.Println(!false)
	fmt.Println(!true)

	// Switch statement
	n := 2
	switch n {
	case 1, 3, 5, 7, 9:
		fmt.Println("1")
	case 2, 4, 6, 8, 10:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	i := 10
	switch {
	case i <= 10:
		fmt.Println("Less than 10")
		// break already applied here
		// only Less than 10 will be printed
	case i <= 20:
		fmt.Println("Less than 20")
	default:
		fmt.Println("Default")
	}

	i2 := 10
	switch {
	case i2 <= 10:
		fmt.Println("Less than 10")
		fallthrough
		// use fall through here, less than 10 and less than 20 will be printed
	case i2 <= 20:
		fmt.Println("Less than 20")
	default:
		fmt.Println("Default")
	}

	var i3 interface{} = []int{}
	switch i3.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a float32")
	case string:
		fmt.Println("i is a string")
	case []int:
		fmt.Println("i is a slice")
		break
		fmt.Println("This will not be printed due to break")
	default:
		fmt.Println("i is another type")
	}
}
