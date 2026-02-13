package main

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		println("You are blocked ", name)
	} else {
		println("Welcome ", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Saus Tartar"
	}

	registerUser("Prince", blacklist)
	registerUser("Saus Tartar", blacklist)

	// langsung pake anonymous function tanpa di simpan ke variable
	registerUser("Kaleng Acar", func(name string) bool {
		return name == "Kaleng Acar"
	})
}