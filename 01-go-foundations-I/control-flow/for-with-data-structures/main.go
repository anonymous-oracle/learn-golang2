package main

import (
	"fmt"
	// "math"
)


func main()  {
	var intArr [5]int = [5]int{1, 2, 3, 4, 5}
	// intArr := [5]int{1, 2, 3, 4, 5}

	for _, val := range intArr {
		fmt.Println(val)
	}

	var stringSlice []string  = []string{"str0", "str1", "str2"}
	// stringSlice := []string{"str0", "str1", "str2"}
	stringSlice = append(stringSlice, "str3")
	stringSlice = append(stringSlice, "str4")

	fmt.Println(stringSlice)

	var cityMap map[string]float64 = map[string]float64 {
		"BLR": 6,
		"VNS": 0.5,
		"CAL": 2,
	}

	val, exist := cityMap["SYD"]

	fmt.Println("Sydney:", val, "Present?", exist)

	var squareSlice []int = make([]int, 0, 10) // if length of the slice is already known, make use of "make"
	// var squareSlice []int = []int{}
	for i := 1; i < 11; i++ {
		// squareSlice = append(squareSlice, int(math.Pow(float64(i), 2)))
		squareSlice = append(squareSlice, i*i)
		fmt.Println("Slice State:", squareSlice)
	}

}