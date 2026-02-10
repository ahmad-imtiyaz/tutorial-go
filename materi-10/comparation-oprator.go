package main

import "fmt"

func main() {
 //operasi perbandingan string
 var nameone = "Prince"
 var nametwo = "Sal"

 var result bool = nameone == nametwo
 fmt.Println(nameone ,"==", nametwo, "=", result)
 
 // perbandingan jumlah string
 var result3 bool = nameone < nametwo
 fmt.Println(nameone, "<", nametwo, "=", result3) // karena huruf P memiliki nilai ascii lebih kecil dari S

 // oprasi perbandingan angka
 var value1 = 100
 var value2 = 200

 var result1 bool = value1 < value2
 fmt.Println(value1, "<", value2, "=", result1)

 // operasi perbandingan angka tidak sama dengan
 var result2 bool = value1 != value2
 fmt.Println(value1, "!=", value2, "=", result2)
}