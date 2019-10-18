package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	// use defer to define resource closure next to resource open
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	// defer take the value when it is deferred
	a := "start"
	defer fmt.Println(a)
	a = "end"

	// panic
	//fmt.Println("start")
	//defer fmt.Println("middle")
	//panic("something bad happened") // panic happened after deferred lines
	//fmt.Println("end")

	// recover
	fmt.Println("goint to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("Done panicking")
}
