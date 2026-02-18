package main

import "fmt"

func Wassup(i int) interface{} {
	if i == 1 {
		return 67
	} else if i == 2 {
		return false
	} else {
		return "Kasih POP"
	}
}

func main() {
	var data interface{}= Wassup(2)
	fmt.Println(data)
}