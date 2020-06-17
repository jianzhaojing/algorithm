package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	str := `{"Name":"apple","Age":14,"Sex":"log"}`

	var s1 Person

	err := json.Unmarshal([]byte(str), &s1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v", s1)
}
