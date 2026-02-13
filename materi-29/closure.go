package main

func main() {
	name := "Prince"
	counter := 0

	increment := func() {
		name = "Sal"
		println("You have called me")
		counter++

	}

	increment()
	increment()
	increment()

	println(counter)
	println(name)
}