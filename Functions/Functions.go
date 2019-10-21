package main

import (
	"fmt"
)

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func main() {
	sayMessage("Hello, Go!")
	greeting := "Hello"
	name := "Stacey"
	sayGreeeting(greeting, name)
	// name will bot be changed by the function sayGreeting
	fmt.Println(name)

	//  if pointer is used
	// the underlying data is used
	greeting2 := "Hello"
	name2 := "Stacey"
	sayGreeting2(&greeting2, &name2)
	fmt.Println(name2)

	// Retrieve multiple return values
	res, err := divide(4.0, 3.0)
	fmt.Println(res, err)
	res2, err2 := divide(4.0, 0.0)
	fmt.Println(res2, err2)

	// Anonymous function
	func() {
		fmt.Println("Hello, Go!")
	}()

	for i := 0; i < 5; i++ {  
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	// function as method
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

func sayMessage(msg string) {
	fmt.Println(msg)
	for i := 0; i < 5; i++ {
		sayMessage2("Hello, GO!", i)
	}
}

func sayMessage2(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

func sayGreeeting(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Ted"
	fmt.Println(name)
}

func sayGreeting2(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

// return multiple variables
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
