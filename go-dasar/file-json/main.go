package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Orang struct {
	Nama string `json:"nama"`
	Umur int    `json:"umur"`
}

func main() {
	kyan := Orang{Nama: "Kyan", Umur: 22}

	// ini ngubah struct jadi json
	data, err := json.Marshal(kyan)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// tulis ke file (0644 = izin file standar)
	err = os.WriteFile("orang.json", data, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//baca file nya
	isi, err := os.ReadFile("orang.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//json ke struct lagi (pakai & karena unmarshal ngisi struct nya)
	var kyan2 Orang
	err = json.Unmarshal(isi, &kyan2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(kyan2)
}
