package main

import "fmt"

func main() {
	var nil16 int16 = 10000
	var nil64 int64 = int64(nil16)
	var nil8 int8 = int8(nil16)

	fmt.Println(nil16, nil64, nil8)

	// konversi string
	var name = "Joko"
	var e byte = name[0]
	var eString string = string(e)
	fmt.Println((eString))
}
