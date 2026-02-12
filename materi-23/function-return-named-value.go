package main

func getFullMahasiswa() (fristClass string, secondClass int, thirdClass string) {
	fristClass = "Prince Sal Yazna"
	secondClass = 3623110067
	thirdClass = "Sistem & Teknologi Informasi"

	// return fristClass, secondClass, thirdClass (ga wajib ditulis class nya namun boleh kalau gabut)
	return
}

func main() {
	name, npm, major := getFullMahasiswa() // menampung return value dari function getFullMahasiswa sesuai urutan
	println("Name :", name)
	println("NPM  :", npm)
	println("Major:", major)
}