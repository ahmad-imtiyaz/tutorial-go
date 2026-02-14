package main

func endApplication() {
	message := recover() // recover akan menangkap error yang terjadi di function yang memanggilnya
	if message != nil {
		println("Aplikasi Error dengan pesan", message)
	}
	panic("Aplikasi Selesai")
}

func runApplication(error bool) {
	defer endApplication()
	if error {
		panic("APLIKASI ERROR")
	}
	println("Aplikasi Berjalan")
}

func main() {
	runApplication(true)
	println("Tes Keluar Ga")
}