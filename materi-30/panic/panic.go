package main

func endApplication() {
	panic("Aplikasi Selesai")
}

func runApplication(error bool) {
	defer endApplication()
	if error {
		println("APLIKASI ERROR")
	}
	println("Aplikasi Berjalan")
}

func main() {
	runApplication(false)
}