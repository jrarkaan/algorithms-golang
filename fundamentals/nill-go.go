package main

import (
	"fmt"
)

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{
		"name": name,
	}
}

func main() {
	var person map[string]string = map[string]string{
		"name": "raka",
	}
	person2 := newMap("")
	if person2 == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(person2)
	}
	fmt.Println("person", person)
}
