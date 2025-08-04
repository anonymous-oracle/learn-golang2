package main

import "fmt"

func main() {
	// var/declarative arrayName/identifier [size]dataType
	var intArr [5]int
	fmt.Println("intArr:", intArr)

	intArr[len(intArr)-1] = 10
	fmt.Println("intArr:", intArr)

	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}

	fmt.Println("fruits:", fruits)

	fmt.Println("Third fruit:", fruits[2])

	originalArr := [3]int{1, 2, 3}

	copiedArr := originalArr
	copiedArr[0] = 10
	fmt.Println("originalArr:", originalArr)
	fmt.Println("copiedArr:", copiedArr) // in go, when a variable is assigned to another variable, it creates a copy of the value and not a reference. To achieve reference-like behavior, you can use slices or pointers.

	for i := 0; i < len(fruits); i++ {
		fmt.Println("Fruit", i+1, ":", fruits[i])
	}

	for idx, val := range fruits {
		fmt.Printf("Fruit %d: %s\n", idx+1, val)
	}

	for _, val := range fruits {
		fmt.Println("Fruit:", val)
	}

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

	originalArr = [3]int{1, 2, 3}
	var copiedArrPtr *[3]int // pointer to an array of 3 elements of type int

	copiedArrPtr = &originalArr
	copiedArrPtr[0] = 10                        // modifying the first element of the array through the pointer
	fmt.Println("originalArr:", originalArr)    // originalArr is modified because copiedArrPtr points to the same array
	fmt.Println("copiedArrPtr:", *copiedArrPtr) // copiedArrPtr reflects the change made to originalArr

}
