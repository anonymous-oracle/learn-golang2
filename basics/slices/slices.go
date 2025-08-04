package main

import "fmt"

func main() {
	// var sliceName []int // declare a slice

	// var numbers []int
	// var numbers1 = []int{1,2,3}

	// numbers2 := []int {4, 5, 6}

	// slice := make([]int, 5) // create a slice with length 5

	a := [5]int{1, 2, 3, 4, 5}

	slice1 := a[1:4]

	fmt.Println(slice1) // Output

	slice1 = append(slice1, 6, 7, 8)
	fmt.Println("Slice1:", slice1) // Output: [2 3 4 6 7 8]

	fmt.Println(slice1) // Output: [2 3 4 6 7 8]

	sliceCopy := make([]int, len(slice1))

	copy(sliceCopy, slice1)

	fmt.Println(sliceCopy) // Output: [2 3 4 6 7 8]
}
