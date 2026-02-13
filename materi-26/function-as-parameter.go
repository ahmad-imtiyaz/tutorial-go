package main

// kode ini juga bisa di sederhanakan dengan alias dengan menggunakan seperti ini
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) { // gantiin bagian yang ini dari filter func(string) string
	nameFiltered := filter(name)
	println("Hello " + nameFiltered)
}

func spamFilter(name string) string {
	if name == "Kurang Di Untung" {
		return "****"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Prince", spamFilter)
	sayHelloWithFilter("Kurang Di Untung", spamFilter)
}