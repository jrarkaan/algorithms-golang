package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Febuari",
		"Maret",
		"April",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	// fmt.Println(len(slice1));
	// fmt.Println(cap(slice1));

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "raka")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)
}
