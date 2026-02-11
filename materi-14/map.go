package main

import "fmt"

func main() {
	// membuat data map string dengan key string
	mahasiswa := map[string]string{
	// var mahasiswa map[string]string = map[string]string{
		"mahasiswa1" : "Prince",
		"mahasiswa2" : "Sal",
		"mahasiswa3" : "Yazna",
	}

	// membuat data map dengan key string dan value int
	npm := map[string]int{
		"npm1" :3623110001,
		"npm2" :3623110002,
		"npm3" :3623110003,
	}

	// menampilkan data map
	fmt.Println("Data Mahasiswa :",mahasiswa)
	fmt.Println("Data NPM :",npm)

	// menampilkan data spesifik pada map
	fmt.Println("Data Mahasiswa 1 :",mahasiswa["mahasiswa1"])
	fmt.Println("Data NPM 1 :",npm["npm1"])

	// mengganti data pada map
	mahasiswa["mahasiswa3"] = "Yaz"
	fmt.Println(mahasiswa["mahasiswa3"])

	// menambah data pada map
	mahasiswa["mahasiswa4"] = "Xyimsimi"
	npm["npm4"] = 3623110004
	fmt.Println("Data Mahasiswa Baru :",mahasiswa)
	fmt.Println("Data NPM Baru :",npm)

	// menghapus data pada map
	delete(mahasiswa,"mahasiswa4")
	fmt.Println("Data Mahasiswa Setelah Dihapus :",mahasiswa)

	// membuat map baru dengan make
	var mahasiswaBaru map[string]string = make(map[string]string)
	mahasiswaBaru["mahasiswa5"] = "Lorem Ipsum"
	mahasiswaBaru["mahasiswa6"] = "Dolor Sit Amet"
	fmt.Println("Data Mahasiswa Baru dengan Make :",mahasiswaBaru)
}