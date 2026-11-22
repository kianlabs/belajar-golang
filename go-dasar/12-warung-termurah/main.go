package main

import "fmt"

type Barang struct {
	Nama  string
	Harga int
}

func main() {
	daftar := []Barang{
		{Nama: "Kopi", Harga: 5000},
		{Nama: "Es Teh", Harga: 3000},
		{Nama: "Indomie", Harga: 7000},
	}
	fmt.Println(termurah(daftar))
}

func termurah(daftar []Barang) Barang {
	palingMurah := daftar[0]
	for _, d := range daftar {
		fmt.Printf("%s - %d\n", d.Nama, d.Harga)
		if palingMurah.Harga > d.Harga {
			palingMurah = d
		}
	}
	return palingMurah
}
