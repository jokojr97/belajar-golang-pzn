package main

import "fmt"

func main() {
	// type deklaration adalah membuat alias untuk suatu tipe data
	type NoKTP string

	var NoKTPjo NoKTP = "3522122232333"
	fmt.Println(NoKTPjo)
}
