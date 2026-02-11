package main

import "fmt"

func main() {
	// Membuat slice dari array
	// Array 12 Bulan
	var months = [12]string{
		"Januari",
		"Februari", 
		"Maret", 
		"April", 
		"Mei", 
		"Juni",
		"Juli", 
		"Agustus", 
		"September", 
		"Oktober", 
		"November", 
		"Desember",
	}
	fmt.Println("ini adalah array bulan",months) // menghasilkan array 12 bulan

	var slice1 = months[4:7] // Mei, Juni, Juli
	fmt.Println(slice1)
	fmt.Println(len(slice1))// Panjang slice1 adalah 3
	fmt.Println(cap(slice1))// Kapasitas slice1 adalah 8 (dari Mei, Juni, Juli, Agustus, September, Oktober, November, Desember)

	// Mengubah elemen pada slice1
	slice1[0] = "May"
	fmt.Println(slice1) // Output: [May Juni Juli]


	// Slice dari array 7 Hari
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println(days) // Output: [Senin Selasa Rabu Kamis Jumat Sabtu Minggu]

	// Membuat slice dari array days
	daysSlice1 := days[5:] // Output: [Sabtu Minggu]
	fmt.Println(daysSlice1)	
	daysSlice1[0] = "Saturday"
	daysSlice1[1] = "Sunday"
	fmt.Println(days) // Output: [Senin Selasa Rabu Kamis Jumat Saturday Sunday]

	// membuat append slice
	daysSlice2 := append(daysSlice1, "Holiday")
	daysSlice2[0] = "Weekend"
	fmt.Println(daysSlice2) // Output: [Weekend Sunday Holiday]
	fmt.Println(days) // Output: [Senin Selasa Rabu Kamis Jumat Saturday Sunday Holiday] (TIDAK BERUBAH KARENA APPEND MEMBUAT SLICE BARU)

	// kode progam make slice
	newSlice := make([]string, 2, 5) // Membuat slice dengan panjang 3 dan kapasitas 5
	newSlice[0] = "Prince"
	newSlice[1] = "Sal"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//kode program copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice,fromSlice)

	fmt.Println(toSlice) // Output: [Senin Selasa Rabu Kamis Jumat Saturday Sunday Holiday]

	// perbedaan slice dan array

	// ini array
	iniArray := [5]int{1,2,3,4,5}
	// iniArray := [...]int{1,2,3,4,5} // ini juga boleh
	iniSlice := []int{1,2,3,4,5}

	fmt.Println("ini adalah array", iniArray)
	fmt.Println("ini adalah slice", iniSlice)
}