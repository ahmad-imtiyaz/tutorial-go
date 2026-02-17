package main

import "fmt"

// ini interface
type HasName interface {
	GetName() string
}

// ini deklarasi hello
func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// contoh pertama
type Person struct {
	Name string
}

func(person Person)GetName()string{
	return person.Name
}

// contoh ke dua
type Animal struct{
	Name string
}

func(animal Animal) GetName()string {
 return animal.Name
}

func main() {
	var FP Person
	FP.Name = "FP"

	SayHello(FP)

	hamster := Animal{
		Name: "Bubble",
	}
	SayHello(hamster)

}
