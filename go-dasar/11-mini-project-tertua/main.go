package main

import "fmt"

type Orang struct {
	Nama string
	Umur int
}

func main() {
	daftar := []Orang{
		{Nama: "Kyan", Umur: 22},
		{Nama: "Jamal", Umur: 24},
		{Nama: "Eman", Umur: 25},
	}
	fmt.Println(tertua(daftar))
}

func tertua(daftar []Orang) Orang {
	palingTua := daftar[0]
	for _, d := range daftar {
		fmt.Printf("%s - %d Tahun\n", d.Nama, d.Umur)
		if d.Umur > palingTua.Umur {
			palingTua = d
		}
	}
	return palingTua
}
