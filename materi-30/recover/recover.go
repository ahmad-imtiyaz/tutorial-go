package main

import "fmt"

func endApplication() {
	message := recover() // recover akan menangkap error yang terjadi di function yang memanggilnya
	if message != nil {
		fmt.Println("Aplikasi Error dengan pesan :", message)
	}
	println("Aplikasi Selesai")
}

func runApplication(error bool) {
	defer endApplication()
	if error {
		panic("APLIKASI ERROR") // panic akan menghentikan eksekusi function yang memanggilnya dan mengirimkan pesan error
	}
	println("Aplikasi Berjalan")
}

func main() {
	runApplication(true)
	println("Tes Keluar Ga")
}