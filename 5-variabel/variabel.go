package main

import "fmt"

func main() {
	var name string
	name = "Joko Riyadi"
	fmt.Println(name)

	name = "Hello World!"
	fmt.Println(name)

	// akan error karena beda tipe datanya
	// name = 100
	// fmt.Println(name)

	// bisa tanpa menuliskan tipedatanya
	var friendname = "Jokori"
	fmt.Println(friendname)

	var numb = 100
	fmt.Println(numb)

	// bisa tanpa menuliskan var tapi dengan menambahkan : sebelum = menjadi :=

	country := "indonesia"
	fmt.Println(country)

	country = "nihoun"
	fmt.Println(country)

	// double declarasi variabel
	var (
		firstname = "joko"
		lastname  = "riyadi"
	)
	fmt.Println(firstname, lastname)
}
