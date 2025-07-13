package main

import "fmt"

func main() {
	// switch statement
	// switch expression {
	// case value1:
	//     // code to execute if expression == value1
	// case value2:
	//     // code to execute if expression == value2
	// default:
	//     // code to execute if expression doesn't match any case
	// }

	// fruit := "apple"

	// switch fruit {
	// case "apple":
	// 	fmt.Println("It's an apple!")
	// case "banana":
	// 	fmt.Println("It's a banana!")
	// case "orange":
	// 	fmt.Println("It's an orange!")
	// default:
	// 	fmt.Println("Unknown fruit!")
	// }

	// day := "Monday"

	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a weekday.")

	// case "Sunday", "Saturday":
	// 	fmt.Println("It's the weekend!")
	// default:
	// 	fmt.Println("Unknown day!")
	// }

	// number := 15
	// switch {
	// case number < 10:
	// 	fmt.Println("Number is less than 10")
	// case number >= 10 && number < 20:
	// 	fmt.Println("Number is between 10 and 19")
	// default:
	// 	fmt.Println("Number is 20 or greater")

	// }
	// // demonstrating fallthrough
	// num := 2
	
	// switch {
	// case num > 1:
	// 	fmt.Println("Greater than 1")
	// 	fallthrough // This will cause the next case to execute as well
	// case num == 2:
	// 	fmt.Println("Number is 2")
	// default:
	// 	fmt.Println("Not two")
	// }

	// demonstrating type switch
	checkType(42)
	checkType("Hello")
	checkType(true)
	checkType(3.14)
	checkType([]int{1, 2, 3})
}

func checkType (x any){
	switch x.(type) { // for type switch fallthrough is not allowed
	case int:
		fmt.Println("x is an int")
	case string:
		fmt.Println("x is a string")
	case bool:
		fmt.Println("x is a bool")
	default:
		fmt.Println("x is of a different type")
}
}