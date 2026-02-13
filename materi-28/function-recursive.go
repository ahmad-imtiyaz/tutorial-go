package main

// ini menggunakan loop biasa
func factorialLoop(value int) int {
result := 1
for i := value; i > 0; i-- {
	result *= i
}
return result
}

// ini menggunakan function recursive
func factorialRecursive(value int) int {
	if value == 1 {
		return 1 
	} else {
		return value * factorialRecursive(value-1)
	}
}
func main() {
	// 5! = 5 x 4 x 3 x 2 x 1 = 120
	// ini factorial dengan loop
	println(factorialLoop(5)) 
	// ini factorial dengan recursive
	println(factorialRecursive(5))
	// ini gambaran sederhananya
	println(5 * 4 * 3 * 2 * 1)
}