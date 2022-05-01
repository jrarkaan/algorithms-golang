package main

import (
	"fmt"
	"time"
)

// var start time.Time
func timeMeasurement(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s", elapsed)
}

func NowAsUnixMilli() int64 {
	time := time.Now().UnixNano() / 1e6
	return time
}

func linearSorting(array *[]int) []int {
	var length = len(*array)
	for i := 0; i <= length-1; i++ {
		var min int = i
		for j := i + 1; j <= length-1; j++ {
			if (*array)[j] < (*array)[min] {
				min = j
			}
		}
		var temp int = (*array)[i]
		(*array)[i] = (*array)[min]
		(*array)[min] = temp
	}
	return *array
}

func main() {
	var array = []int{
		-2, -4, 3, 1, 0, 8, 2,
	}
	fmt.Println(linearSorting(&array))
	defer timeMeasurement(time.Now())
}
