package main

import "fmt"

func main() {
	fmt.Println("String")
	fmt.Println("Rossi")
	fmt.Println("Fakhri Rossi")

	fmt.Println("Jumlah huruf: ", len("Fakhri Rossi"))
	fmt.Println("Fakhri Rossi"[5]) // output -> 105 (byte dari i adalah 105)
}
