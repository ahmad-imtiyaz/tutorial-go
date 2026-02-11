package main

import "fmt"

func main() {
	// membuat if expression
	var nilai = 90

	if nilai > 99 {
		fmt.Println("LULUS")
	}

	// membuat if-else expression
	var nama string = "Prince"
	if nama == "Prince" {
		fmt.Println("Halo Prince")
	} else {
		fmt.Println("Siapa Kamu?")
	}

	// membuat else if expression
	var tinggiBadan = 170
	if tinggiBadan > 180 {
		fmt.Println("Kamu Tinggi")
	} else if tinggiBadan >= 170 {
		fmt.Println("Kamu Lumayan Tinggi")
	} else {
		fmt.Println("Kamu Pendek")
	}

	// memembuat if dengan short statement
	if umur := 20; umur >= 17 {
		fmt.Println("Kamu Sudah Dewasa")
	} else {
		fmt.Println("Kamu Masih Anak-Anak")
	}
}