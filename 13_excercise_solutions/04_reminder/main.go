package main

import "fmt"

func main() {
	var small int
	var large int

	fmt.Print("input small number: ")
	fmt.Scan(&small)

	fmt.Print("input large number: ")
	fmt.Scan(&large)

	fmt.Printf("the reminder is %v", (large%small))
}