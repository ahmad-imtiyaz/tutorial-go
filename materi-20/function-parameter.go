package main

import "fmt"

func wassup(name string, npm uint32) {
	fmt.Println("Hello", name, "Your NPM is", npm)
}

func main() {
	name := "Prince"
	wassup(name, 3623110067)
	wassup("Sal", 3623110099)

}