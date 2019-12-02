package main

import (
	"fmt"
	"reflect"
)

type teststruct struct {
	id   int    `json:"key,omitempty"`
	name string `json:"name,omitempty"`
	tag  string `json:"tag"`
	text string `json:"text"`
}

func (b teststruct) PrintFields() map[string][]string {
	val := reflect.ValueOf(b)
	fields := []string{}
	for i := 0; i < val.Type().NumField(); i++ {
		if val.Type().Field(i).Tag.Get("json") != "-" {
			fields = append(fields, val.Type().Field(i).Name)
		}
	}
	result := make(map[string][]string)
	key := reflect.TypeOf(b).String()
	result[key] = fields
	return result
}

func main() {
	test := teststruct{
		id:   1,
		name: "testname",
		tag:  "testtag",
		text: "testText",
	}
	fmt.Println(test.PrintFields())
}
