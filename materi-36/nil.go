package main

import "fmt"

func NewMap(name string)map[string]string{
	if name == ""{
		return nil
	} else {
		return map[string]string{
			"Name":name,
		}
	}
}
func main (){
	// kalau data ga di isi
	// var person map[string]string = NewMap("")
	// kalau data di isi
	var person map[string]string = NewMap("Prince")

	if person == nil{ 
		fmt.Println("Data ini Kosong")
	} else {
		fmt.Println(person)
	}
}