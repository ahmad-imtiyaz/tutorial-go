package main

import "fmt"

func main() {
	var name string 

	name = "little prince"
	fmt.Println(name)

	name = "future king"
	fmt.Println(name)

	// deklrasi data tanpa tipe data

	var friendName = "yazna"
	fmt.Println(friendName)

	var age = 19
	fmt.Println(age)

	// deklarasi data tanpa var

	country := "Switzerland"
	fmt.Println(country)

	country = "Japan"
	fmt.Println(country)

	// deklarasi multiple variable
	
	var (
		firstName = "Prince"
		lastName = "Sal"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}