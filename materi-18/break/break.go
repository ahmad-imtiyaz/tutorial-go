package main

import "fmt"

func main() {
	// Penggunaan Break dalam loop
	for i := 1; i <= 20; i++ {
		if i > 10 {
			break
		}
		fmt.Println("Perulangan Ke-", i)
	}
}