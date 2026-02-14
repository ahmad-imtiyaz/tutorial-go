package main

func endApplication() {
	println("Aplikasi Selesai")
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
}