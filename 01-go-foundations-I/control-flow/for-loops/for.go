package main

import (
	"fmt"
)

func main()  {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}

	uniString := "नमस्ते"
	for _, val := range uniString {
		fmt.Printf("%U\n", val)
	}

	for idx, val := range []string{"str0", "str1", "str2"} {
		fmt.Println(idx, "->", val)
	}

	for i := 1; i < 4; i ++ {
		for j := 1; j < 4; j ++ {
			fmt.Println(i*j)
		}
	}

	i := 1
	for {
		if i % 5 == 0 && i % 7 == 0 {
			break
		}
		i++
		if i > 50 {
			break
		}
	}

}