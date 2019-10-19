package main

import (
	"fmt"
)

type myStruct struct {
	foo int
}

func main() {
	// Pointer
	var a int = 22
	var b *int = &a
	fmt.Println(a, *b)
	fmt.Println(&a, b)
	a = 27
	fmt.Println(a, *b)
	*b = 77
	fmt.Println(a, *b)

	arr := [3]int{1, 2, 3}
	arr1 := &arr[1]
	arr2 := &arr[2]
	fmt.Printf("%v %p %p", arr, arr1, arr2)

	// New Function
	var ms *myStruct
	fmt.Println(ms)
	ms = new(myStruct)
	fmt.Println(ms)

	// Dereference
	// No need to add *
	ms.foo = 123
	fmt.Println(ms.foo)
}
