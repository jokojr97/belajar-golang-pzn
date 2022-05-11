package main

import "fmt"

func main() {
	const name string = "Joko Riyadi"
	// const firstname = "joko"
	const number = 1000
	fmt.Println(name)
	fmt.Println(number)

	// deklarasi multiple contant
	const (
		firstname = "jokori"
		lastname  = "Riyadi"
	)
	fmt.Println(firstname, lastname)
}
