package main

import (
	"context"
	"database/sql"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

type Product struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	ImageURL  string `json:"image_url"`
	HasSizes  bool   `json:"has_sizes"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Variant struct {
	ID        int     `json:"id"`
	ProductID int     `json:"product_id"`
	SizeLabel *string `json:"size_label"`
	Stock     int     `json:"stock"`
	SortOrder int     `json:"sort_order"`
}

func openStore() (*sql.DB, error) {
	err := os.MkdirAll("data", 0755)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite", "data/stok-baju.db")
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("PRAGMA journal_mode = WAL")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS products (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  price INTEGER NOT NULL DEFAULT 0,
  image_url TEXT,
  has_sizes INTEGER NOT NULL DEFAULT 1,
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL
);`)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS product_variants (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
  size_label TEXT,
  stock INTEGER NOT NULL DEFAULT 0,
  sort_order INTEGER NOT NULL DEFAULT 0
);`)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateProduct(db *sql.DB, p Product, variants []Variant) (int64, error) {
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	now := time.Now().Format(time.RFC3339)
	res, err := tx.Exec(`INSERT INTO products (name, price, image_url, has_sizes, created_at, updated_at)
	    VALUES (?, ?, ?, ?, ?, ?`, p.Name, p.Price, p.ImageURL, p.HasSizes, now, now)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	for _, v := range variants {
		_, err := tx.Exec(`INSERT INTO product_variants (product_id, size_label, stock, sort_order)
	    VALUES (?, ?, ?, ?)`, id, v.SizeLabel, v.Stock, v.SortOrder)
		if err != nil {
			return 0, err
		}
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return id, nil
}
