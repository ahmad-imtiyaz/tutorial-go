package main

import "fmt"
func main() {
	// deklarasi array dengan tipe data string dan panjang 3
	var names [3]string

	names[0] = "Prince"
	names[1] = "Sal"
	names[2] = "Yazna"
	
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// deklarasi array dengan langsung memanggil isinya
	var values = [3]int{
		90,
		80,
		70,
	}
	fmt.Println(values)

	// model string 
	var nama_banyak = [3]string{
		"Prince",
		"Sal",
		"Yazna",
	}
	fmt.Println(nama_banyak)

	// menghitung panjang array
	var testestis = [100]string{
	}
	fmt.Println(len(values))
	fmt.Println(len(nama_banyak))
	fmt.Println(len(testestis))
}