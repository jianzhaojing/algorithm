package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func main() {
	p := Person{
		Name: "apple",
		Age:  14,
		Sex:  "log",
	}

	fmt.Printf("%#v\n", p)

	jsonByte, _ := json.Marshal(p)
	jsonStr := string(jsonByte)
	fmt.Printf("%v", jsonStr)

}
