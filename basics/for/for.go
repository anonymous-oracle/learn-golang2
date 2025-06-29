package main

import "fmt"

func main() {
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// // iterate over a collection
	// numbers := []int{1,2,3,4,5,6}
	// for idx, number := range numbers {
	// 	fmt.Printf("Index: %d, Number: %d\n", idx, number)
	// }

	// for i:=1; i<=10; i++ {
	// 	if i%2==0 {
	// 		continue
	// 	}

	// 	fmt.Println("Odd number:", i)

	// 	if i == 5 {
	// 		break
	// 	}
	// }


	// Outer loop
	// rows := 5
	// for i:=1; i<= rows; i++ {
	// 	// inner loop for spaces before stars
	// 	for j:=1; j <= rows - i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	// inner loop for stars
	// 	for k:=1; k <= (2*i - 1); k++ { // odd number of stars
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println() // Move to the next line after each row
	// }

	for i := range 10 {
		fmt.Println(10 - i)
	}

}
