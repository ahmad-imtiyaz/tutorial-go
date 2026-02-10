package main

import "fmt"
func main() {
	// mengubah tipe data atau alias dari uint32 ke alias NPM
	type NPM uint32

	var prince NPM = 3623110067
	fmt.Println("NPM Prince :", prince)

	// mengubah tipe data atau alias dari bool ke alias status
	type status bool

	var statusMahasiswa status = true
	fmt.Println("Status Mahasiswa Prince :", statusMahasiswa)
}