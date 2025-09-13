package main

import (
	"fmt"
	"log"
)

const DaysInWeek int = 7

func checkCity(city string) {
	if city == "VNS" {
		fmt.Println("That's Varanasi!")
	} else if city == "BLR" {
		fmt.Println("Namaskara Bengaluru!")
	} else {
		fmt.Println("City not recognized.")
	}
}

func demoLoops() {
	// Classic for loop
	for i := 1; i <= 3; i++ {
		fmt.Println("Count:", i)
	}

	// While-style
	j := 0
	for j < 3 {
		fmt.Println("j:", j)
		j++
	}

	j = 0
	for {
		fmt.Println("j:", j)
		j++

		if j > 3 {
			break
		}
	}
}

func greetDay(day int) {
	switch day {
	case 1:
		fmt.Println("Monday blues...")
	case 5:
		fmt.Println("Friday vibes!")
	case 7:
		fmt.Println("Sunday rest day")
		fallthrough // runs the next case's body unconditionally
	default:
		fmt.Println("Just another day")
	}
}

type UserID int

func NewUserID(id int) (UserID, error) {
	if id <= 0 { // guard clause
		return 0, fmt.Errorf("invalid id")
	}
	return UserID(id), nil
}

func describe(x any) {
	switch v := x.(type) {
	case nil:
		fmt.Println("nil interface")
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Printf("other (%T)\n", v)
	}
}

// exercise
func main() {
	fmt.Println("Hello")
	checkCity("VNS")
	greetDay(DaysInWeek)

	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Total: ", sum)

	_, err := NewUserID(-1)

	if err != nil {
		log.Println(err)
	}

	describe(err)
}
