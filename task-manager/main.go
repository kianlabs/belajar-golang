package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID      int    `json:"id"`
	Judul   string `json:"judul"`
	Selesai bool   `json:"selesai"`
}

func main() {
	tasks := []Task{
		{ID: 1, Judul: "Belajar Go Dasar", Selesai: true},
		{ID: 2, Judul: "Belajar Go Lanjut", Selesai: false},
		{ID: 3, Judul: "Belajar Go Web", Selesai: false},
	}
	err := Simpan(tasks)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	daftar, err := Muat()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, t := range daftar {
		fmt.Printf("ID: %d - Judul: %s - %v\n", t.ID, t.Judul, t.Selesai)
	}
}

func Simpan(tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func Muat() ([]Task, error) {
	isi, err := os.ReadFile("tasks.json")
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(isi, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
