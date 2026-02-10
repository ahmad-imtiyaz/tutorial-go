package main

import "fmt"

func main() {
	const firstName string = "Prince"
	const lastName = "Yazna"
	const value = 67

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	// deklarasi multiple constant
	const (
		country = "Indonesia" 
		age = 19
		height = 169
	 )

	 fmt.Println(country)
	 fmt.Println(age)
	 fmt.Println(height)
}