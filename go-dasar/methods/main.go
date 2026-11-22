package main

import "fmt"

type Orang struct {
	Nama string
	Umur int
}

type Karyawan struct {
	Orang
	Gaji int
}

func (o *Orang) Sapa() string {
	return "Halo, saya " + o.Nama
}

func (o *Orang) Dewasa() bool {
	return o.Umur >= 18
}

func (o *Orang) UlangTahun() {
	o.Umur++
}

func main() {
	kyan := Karyawan{Nama: "Kyan", Umur: 22, Gaji: 3000000}
	jamal := Orang{Nama: "Jamal", Umur: 24}
	jambul := Orang{Nama: "Jambul", Umur: 15}
	fmt.Println(kyan.Sapa())
	fmt.Println(jamal.Dewasa())
	fmt.Println(jambul.Dewasa())
	fmt.Println(kyan.Umur)
	kyan.UlangTahun()
	fmt.Println(kyan.Umur)
	fmt.Println(kyan.Nama)
	
}
