package main

import "fmt"

type Orang struct {
	Nama string
	Umur int
}

func main() {
	o := Orang{Nama: "Kyan", Umur: 22}
	fmt.Println(o.Nama)
	fmt.Println(o.Umur)
}
