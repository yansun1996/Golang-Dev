package main

import (
	"fmt"
)

const a int16 = 27

// iota changes within const block
const (
	b0 = iota
	b1 = iota
	b2 = iota
)

const (
	_  = iota // ignore first zero value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
	EB = 1 << (10 * iota)
	ZB = 1 << (10 * iota)
	YB = 1 << (10 * iota)
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials
	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	const myConst int = 33
	fmt.Printf("%v, %T\n", myConst, myConst)

	// Cannot create const that needs to be calculated at runtime
	//const myConst2 float32 = math.Sin(1.57)
	//fmt.Printf("%v, %T\n", myConst2, myConst2)

	fmt.Printf("%v, %T\n", a, a)
	const a int = 22
	fmt.Printf("%v, %T\n", a, a)

	var b int = 1
	fmt.Printf("%v, %T\n", a+b, a+b)

	// iota changes within const block
	fmt.Printf("%v, %T\n", b0, b0)
	fmt.Printf("%v, %T\n", b1, b1)
	fmt.Printf("%v, %T\n", b2, b2)

	fileSize := 4000000000.
	fmt.Printf("%.2fGB", fileSize/GB)

	// admin usage
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v", isAdmin&roles == isAdmin)
}
