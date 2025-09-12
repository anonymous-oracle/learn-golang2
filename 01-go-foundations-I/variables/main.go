package main 

import (
	"fmt"
	"time"
)

func main()  {
	// explicit type
	var age int = 25

	// type inference
	name := "Mr. X"

	// zero value (not initialized, gets default)
	var score int
	var active bool

	// constant
	const Pi = 3.14159

	fmt.Println("Hello,",name)
	fmt.Println("Age,", age)
	fmt.Println("Score (zero value):", score)
	fmt.Println("Active (zero value):",active)
	fmt.Println("Pi:", Pi)
	fmt.Println("Current year:", time.Now().Year())
}