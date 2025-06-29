package main

import (
	"fmt"
	foo "net/http"
)

func main() {
	fmt.Println("The Go standard library")

	// resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Response Status:", resp.Status)
	// if resp.StatusCode == http.StatusOK {
	if resp.StatusCode == foo.StatusOK {
		fmt.Println("Successfully fetched the data!", resp.Body)
	} else {
		fmt.Println("Failed to fetch data, status code:", resp.StatusCode)
	}
}
