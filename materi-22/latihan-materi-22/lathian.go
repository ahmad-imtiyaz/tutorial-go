package main

import "fmt"

func bagi(a int, b int) (int, string) {
    if b == 0 {
        return 0, "tidak bisa bagi nol"
    }
    return a / b, "berhasil"
}

func main() {
    hasil, pesan := bagi(10, 0)
    fmt.Println(hasil + 5)
	fmt.Println(pesan)
}
