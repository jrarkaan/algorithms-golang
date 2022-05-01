// algorithms buble sort
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array = [...]string{
		"januari",
		"febuari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	var slice1 = array[4:7]
	fmt.Println("slice1", slice1)
	fmt.Println("Length slice1", len(slice1))
	fmt.Println("cap slice1", cap(slice1))

	array[4] = "diubah"
	fmt.Println(array)
	//fmt.Println(array[6])

	var slice2 = array[10:]
	fmt.Println((slice2))
	var slice3 = append(slice2, "raka")
	fmt.Println("slice3", slice3)
	slice3[1] = "Bukan-Desember"
	fmt.Println(slice3)
	fmt.Println(slice2)

	var iniSlice = [...]int8{
		1, 3, 4, 5, 6,
	}
	fmt.Println("type", reflect.TypeOf(iniSlice).Kind())
}
