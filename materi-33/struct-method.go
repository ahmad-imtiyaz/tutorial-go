package main

import "fmt"

type Mahasiswa struct {
	Name, Class string
	Npm         int
}

func (mahasiswa Mahasiswa) sayWassup(name string) {
	fmt.Println("Hello", name, "My name is :", mahasiswa.Name)
}

func (mhs Mahasiswa) damnIt() {
	fmt.Println("Damn it", mhs.Name)
}

func main() {
		var Yazna Mahasiswa 
		Yazna.Name = "Yazna"
		Yazna.Class = "Sastra-Informatika"
		Yazna.Npm = 123456789
	
		// memanggil method sayWassup
		Yazna.sayWassup("Prince")
		// memanggil method damnIt
		Yazna.damnIt()

}
