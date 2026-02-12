package main

import "fmt"

func getFullName() (string, string, string) {
	return "Prince", "Sal", "Yazna"
}

func main() {
	fristName, middleName, lastName := getFullName()
	fmt.Println(fristName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}