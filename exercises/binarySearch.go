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

func bubbleSort(arrayInput *[]int8) []int8 {
	var (
		n      int  = len(*arrayInput)
		sorted bool = false
	)
	for !sorted {
		var swapped bool = false
		for i := 0; i < n-1; i++ {
			if (*arrayInput)[i] > (*arrayInput)[i+1] {
				(*arrayInput)[i+1], (*arrayInput)[i] = (*arrayInput)[i], (*arrayInput)[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
	return *arrayInput
}

func binarySearch(needleUser *int8, itemsArray *[]int8, filter func(bool) string) {
	var (
		low  int = 0
		high int = len(*itemsArray) - 1
	)
	for low <= high {
		median := (low + high) / 2
		if (*itemsArray)[median] < (*needleUser) {
			low = median + 1
		} else {
			high = median - 1
		}
		fmt.Println(low)
	}
	if low == len(*itemsArray) || (*itemsArray)[low] != (*needleUser) {
		valueRaw := filter(false)
		fmt.Println("result is", valueRaw)
	}
	valueRaw := filter(true)
	fmt.Println("result is", valueRaw)
}

func main() {
	var array = &[]int8{
		-2, -4, 3, 1, 0, 8, 2,
	}
	newArray := bubbleSort(array)
	var userNeed int8 = 1
	binarySearch(&userNeed, &newArray, func(result bool) string {
		if result {
			return "exists"
		} else {
			return "doesnt exists"
		}
	})
	defer timeMeasurement(time.Now())
}
