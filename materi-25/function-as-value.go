package main

func sayTesTesTis(name string) string {
	return "Tes Tes Tis " + name
}

func main() {

	princeSay := sayTesTesTis

	result := princeSay("Welcome Prince")
	println(result)
}