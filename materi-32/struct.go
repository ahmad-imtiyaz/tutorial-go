package main

import "fmt"

type Mahasiswa struct {
	Name,Class string
	Npm int
}


func main() {
	// pemanggilan struct yang pertama
	var Yazna Mahasiswa 
		Yazna.Name = "Yazna"
		Yazna.Class = "Sastra-Informatika"
		Yazna.Npm = 123456789
		fmt.Println(Yazna)
		// bisa di panggil satu satu
		fmt.Println(Yazna.Name)
		fmt.Println(Yazna.Class)
		fmt.Println(Yazna.Npm)

	// pemanggilan struct yang kedua
	Prince := Mahasiswa{
		Name: "Prince",
		Class: "Sastra-Informatika",
		Npm: 987654321,
	}
	fmt.Println(Prince)

	// pemanggilan struct yang ketiga
	Sal := Mahasiswa{"Sal", "Sastra-Informatika", 16655455}
	fmt.Println(Sal)
	}
