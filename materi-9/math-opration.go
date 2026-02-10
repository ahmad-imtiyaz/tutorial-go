package main

import "fmt"
func main() {
	// operasi biasa
 var a = 10
 var b = 18
 var c = a + b

 fmt.Println(c)

 // operasi langsung
 var result = 20 - 10
 fmt.Println(result)

 //augmented operation
 var i = 100
 i += 10 // i = i + 10
 fmt.Println(i)

 // unary operation
 i++ // i = i + 1 (yang dimana niali i tadi 110)
 fmt.Println(i)

 var negatif = -100
 var positif = +100
 fmt.Println(negatif)
 fmt.Println(positif)
}