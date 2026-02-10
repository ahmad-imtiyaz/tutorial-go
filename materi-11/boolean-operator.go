package main

import "fmt"

func main() {
	// Operasi boolean AND
	var nilai_ipa = 100
	var nilai_mtk = 90

	var hasil_ipa = nilai_ipa > 80
	var hasil_mtk = nilai_mtk > 80

	fmt.Println(hasil_ipa)
	fmt.Println(hasil_mtk)

	var hasil_akhir = hasil_ipa && hasil_mtk
	fmt.Println(hasil_akhir)

	// Operasi boolean OR
	var nilai_bindo = 70
	var nilai_inggris = 900
	
	var hasil_bindo = nilai_bindo > 80
	var hasil_inggris = nilai_inggris > 80
	fmt.Println(hasil_bindo)
	fmt.Println(hasil_inggris)

	var hasil_akhir_or = hasil_bindo || hasil_inggris
	fmt.Println(hasil_akhir_or)

	// Operasi boolean Negasi
	var isLogin = true

	fmt.Println(!isLogin)

	// Operasi boolean langsung (ini yang sering digunakan)
	var nilai_sejarah = 60
	var nilai_pendidikan = 90
	fmt.Println(nilai_sejarah > 80 || nilai_pendidikan > 80)
}