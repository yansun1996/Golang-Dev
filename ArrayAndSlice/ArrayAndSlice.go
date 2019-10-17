package main

import (
	"fmt"
)

func main() {
	grades := [3]int{97, 98, 99}
	fmt.Printf("Grades: %v\n", grades)
	var students [3]string
	students[0] = "Lisa"
	students[1] = "Armold"
	students[2] = "Ahmed"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of students: %v\n", len(students))

	// array in array
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}

	// array deep copy
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// array shallow copy
	c := [...]int{1, 2, 3}
	d := &c
	d[1] = 5
	fmt.Println(c)
	fmt.Println(d)

	// slice deep copy
	slice := []int{1, 2, 3}
	slice2 := slice
	slice2[1] = 5
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Printf("Length: %v\n", len(slice))
	fmt.Printf("Capacity: %v\n", cap(slice))

	makedSlice := make([]int, 3)
	fmt.Println(makedSlice)
	fmt.Printf("Length: %v\n", len(makedSlice))
	fmt.Printf("Capacity: %v\n", cap(makedSlice))

	// Append
	test := []int{}
	fmt.Println(test)
	fmt.Printf("Length: %v\n", len(test))
	fmt.Printf("Capacity: %v\n", cap(test))
	test = append(test, 1, 2, 3, 4, 5)
	fmt.Println(test)
	fmt.Printf("Length: %v\n", len(test))
	fmt.Printf("Capacity: %v\n", cap(test))
	test = append(test, []int{1, 2, 3, 4, 5}...)
	fmt.Println(test)
	fmt.Printf("Length: %v\n", len(test))
	fmt.Printf("Capacity: %v\n", cap(test))
}
