package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

type Animal struct {
	Name   string `required max: "100"` // Tag
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	// Map
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 39250017,
		"Texas":      27862596,
	}
	fmt.Println(statePopulations)

	m := map[[3]int]string{}
	fmt.Println(m)

	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations)
	fmt.Println(len(statePopulations))

	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Georgia"])

	pop, ok := statePopulations["Georgia"]
	fmt.Println(pop, ok)

	// Struct
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	bDoctor := Doctor{
		number:    3,
		actorName: "Test",
		companions: []string{
			"L",
			"A",
			"J",
		},
	}

	cDoctor := struct{ name string }{name: "test"}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions)
	fmt.Println(bDoctor)
	fmt.Println(cDoctor)

	// Struct deep copy
	anotherDoctor := aDoctor
	anotherDoctor.actorName = "Changed Name"
	fmt.Println(aDoctor.actorName)
	fmt.Println(anotherDoctor.actorName)

	// Something similar to inherence
	bird := Bird{}
	bird.Name = "Emu"
	bird.Origin = "Australia"
	bird.SpeedKPH = 48
	bird.CanFly = false
	fmt.Println(bird)

	// Tag
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
