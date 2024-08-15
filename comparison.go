package main

import "fmt"

func main() {
	name1 := "Rossi"
	name2 := "Rossi"
	name3 := "rossi"

	fmt.Println(name1 == name2)
	fmt.Println(name1 == name3)

	a := "a"
	b := "b"

	fmt.Println(a > b)
}
