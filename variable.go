package main

import "fmt"

func main() {
	var name string = "Fakhri Rossi"
	fmt.Println(name)

	name = "Budi Pekerti"
	fmt.Println(name)

	var score int8 = 100
	fmt.Println(score)

	// declare/init variable without var keyword
	var address = "Bandung"
	fmt.Println(address)

	// declare variable without var -> u have to fill the value at first
	// most people use this method
	country := "Finland"
	fmt.Println(country)

	var (
		firstName = "Fakhri"
		lastName  = "Rossi"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
