package main

import "fmt"

type Buku struct {
	Judul string
	Harga int
}

func main() {
	// slice of struct: daftar buku
	buku := []Buku{
		{Judul: "Go Dasar", Harga: 50000},
		{Judul: "Go Lanjut", Harga: 75000},
	}

	// tambah satu pakai append
	buku = append(buku, Buku{Judul: "Go Web", Harga: 100000})

	// muter semua isi slice
	for i, b := range buku {
		fmt.Println(i, b.Judul, b.Harga)
	}
	fmt.Println("jumlah:", len(buku))
}
