package main

import "fmt"

type Apapun interface {
}

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func myfun(a interface{}) {
	var value, ok interface{} = a.(float64)
	fmt.Println(value, ok)
}

func main() {
	var data interface{} = Ups(9)
	fmt.Println(data)

	var testvariabel interface{} = "lo jeke"
	myfun(testvariabel)
}
