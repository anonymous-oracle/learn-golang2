package main

import "fmt"

func main() {
	
	// sum := 0
	// for {
	// 	sum += 10
	// 	fmt.Println("Sum: ", sum)
	// 	if sum > 100 {
	// 		break
	// 	}
	// }

	num := 1
	for num <= 10 {
		if num % 2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd Number: ", num)
		num++ // ++ increment operator
	}
	

}
