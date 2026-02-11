package main

import "fmt"

func main() {
	// membuat perulangan dengan for loop
	counter := 1

	for counter <=5 {
		fmt.Println("Perulangan Ke-", counter)
		counter++
	}
	// membuat for dengan statement
	for count :=1; count <=10; count++ {
		fmt.Println("Perulangan Ke-", count)
	}
	// membuat for range
	slice := []string{"Prince", "Sal", "Yazna"}
	for index, name := range slice {
		fmt.Println("Index ke-", index, "=", name)
	}
	// membuat for range pada map
	var person map[string]string = map[string]string{
		"name" : "Prince",
		"address" : "Indonesia",
		"job" : "Programmer",
	}
	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}