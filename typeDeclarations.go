package main

import "fmt"

func main() {
	type TypeId string
	var userId TypeId = "121212"

	fmt.Println(userId)
	fmt.Println(TypeId("646464")) // convert string to TypeId
}
