package main

import "fmt"

func main() {
	// membuat switch case
	name := "Prince"

	switch name {
		case "Prince":
			fmt.Println("Ohayoo Prince")
		case "Sal":
			fmt.Println("Ohayoo Sal")
		default :
		 fmt.Println("Who are you!")
	}

	// membuat switch dengan short statement
	switch length := len(name); length > 5 {
	case true :
		fmt.Println("Namamu Panjang")
	case false:
		fmt.Println("Namamu Pendek")
	}	
		
	// membuat switch tanpa expression(kondisi)
	var name2 string = "Sal"
	switch {
	case name2 == "Prince":
		fmt.Println("Hello Prince")
	case name2 == "Sal":
		fmt.Println("Hello Sal")
	default:
		fmt.Println("Who are you!")
	}

}