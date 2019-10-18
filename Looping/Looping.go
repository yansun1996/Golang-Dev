package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 5; i += 2 {
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// Labeled loop
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	// iterate slice
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	str := "Hello, GO!"
	for k, v := range str {
		fmt.Println(k, v)
		fmt.Println(k, string(v))
	}
}
