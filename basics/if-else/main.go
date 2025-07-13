package main

func main() {
	temperature := 30
	if temperature < 20 {
		println("It's cold outside")
	} else if temperature < 30 {
		println("It's a nice day")
	} else {
		println("It's hot outside")
		if temperature > 40 {
			println("It's extremely hot!")
		}
	}

}
