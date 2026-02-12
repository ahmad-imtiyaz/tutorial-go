package main

import "fmt"

func main() {
	// Penggunaan Continue dalam loop
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("Perulangan Genap Ke-", i)
	}
}