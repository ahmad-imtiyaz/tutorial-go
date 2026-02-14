package main

func loggig() {
	println("Selesai memanggil function")
}

func runApplication(value int) {
	// defer akan menunda eksekusi function sampai function yang memanggilnya selesai di eksekusi
	defer loggig()
	println("Run Application")
	result := 10 / value
	println("Result", result)
}
func main() {
	// defer akan dieksekusi walaupun terjadi error di function yang memanggilnya
	runApplication(0)
}