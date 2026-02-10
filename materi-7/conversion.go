package main

import "fmt"

func main() {
	// jika niali melebihi batas maka akan di hitung kembali dari nilai minimum mereka
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println("nilai 32 =", nilai32)
	fmt.Println("nilai 64 =", nilai64)
	fmt.Println("nilai 8 =", nilai8)

	// Nilai konversi akan mengikuti batas maks dan min dari tipe data tujuan
	var nilai32Normal int32 = 127
	var nilai64Normal int64 = int64(nilai32Normal)
	var nilai8Normal int8 = int8(nilai32Normal)

	fmt.Println("nilai 32 =", nilai32Normal)
	fmt.Println("nilai 64 =", nilai64Normal)
	fmt.Println("nilai 8 =", nilai8Normal)

	// konversi dari nilai byte ke string
	var name = "Prince"
	var x byte = name[3]
	var eString string = string(x)

	fmt.Println(name)
	fmt.Println(x)
	fmt.Println(eString)

}