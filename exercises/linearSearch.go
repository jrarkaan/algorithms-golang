package main

import (
	"fmt"
	"time"
)

type responseSchema struct {
	nameResponse string

	valueResponseSchema struct {
		valueResponse int8
	}
}

// var start time.Time
func timeMeasurement(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s", elapsed)
}

func linearSearch(dataArray *[]int8, searchValue *int8) *responseSchema {
	var findResult bool
	var result responseSchema
	for i, _ := range *dataArray {
		findResult = false
		if (*searchValue) == (*dataArray)[i] {
			findResult = true
		}
	}

	if findResult {
		result = responseSchema{
			nameResponse:        "status",
			valueResponseSchema: struct{ valueResponse int8 }{valueResponse: 1},
		}
	} else {
		result = responseSchema{
			nameResponse:        "status",
			valueResponseSchema: struct{ valueResponse int8 }{valueResponse: 0},
		}
	}
	return &result
}

func main() {
	var originalArray = []int8{90, 60, 40, 50, 34, 49, 30}
	var searchUser int8 = 8
	fmt.Println(linearSearch(&originalArray, &searchUser))
	defer timeMeasurement(time.Now())
}
