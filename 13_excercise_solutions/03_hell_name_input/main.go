package main

import "fmt"

func main() {
	var name string

	fmt.Print("input your name: ")
	fmt.Scan(&name)
	fmt.Printf("Hello %v", name)
}