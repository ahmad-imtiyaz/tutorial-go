package main

import "fmt"

// Function dengan return value tipe string
func wassupTeam(name string) string {
	return "Wassup" + name
}

// Function dengan return value tipe int
func ngitungBre(a int, b int) int {
	return a + b
}

// Main function
func main() {
	result := wassupTeam(" Prince")
	fmt.Println(result)
	itungBre := ngitungBre(10, 15)
	fmt.Println(itungBre)
}