package main

import (
	"fmt"
)
// Value object using type 
type UserID int

func NewUserID(id int) (UserID, error) { // defining the return types of functions; functions are Always declared outside
	if id <= 0 {
		return 0, fmt.Errorf("invalid id")
	}
	return UserID(id), nil // explicit type casting since types are not inferred implicitly
}


func main()  {
	city := "VNS"
	var name string = "Mr. X"

	fmt.Println("Hello, ", name, "City:", city)

	// constant
	const DaysInWeek = 7
	fmt.Println("Number of days in a week, ", DaysInWeek)

	// zero value variable
	var zero int
	fmt.Println("Zero value variable: ", zero)

	
	NewUserID(7)
	NewUserID(0)
	
}