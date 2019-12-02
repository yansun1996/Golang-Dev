// Generating arbitrary JSON data using map
package main

import "fmt"
import "encoding/json"


func main() {
    keys := []string{"Segment", "Jan", "Feb", "Apr"}
    values := []interface{}{"ME",1000,2000,3000}
    fmt.Println(keys)
    fmt.Println(values)
    // map values to keys
    m := make(map[string]interface{})
    for i,v := range values {
	m[keys[i]] = v
    }
    fmt.Println(m)
    // convert map to JSON    
    data, err := json.Marshal(m)
    if err != nil {
        fmt.Println("Error!")
    }
    fmt.Printf("%s", data)

    test := map[string]map[string][]string
    map1 := map[string][]string{
		"1":[]string{"one","two"},
	}
    test["keyone"] = map1
    data2, _ := json.Marshal(test)
    fmt.Printf("%s", data2)
}
