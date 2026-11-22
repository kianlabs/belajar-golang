package main

import "testing"

func TestBagi(t *testing.T) {
	hasil, err := Bagi(10, 2)
	if err != nil {
		t.Fatalf("tidak boleh error: %v", err)
	}
	if hasil != 5 {
		t.Errorf("Bagi(10, 2) = %d, mau 5", hasil)
	}
}

func TestBagiNol(t *testing.T) {
	hasil, err := Bagi(10, 0)
	if err == nil {
		t.Fatalf("harusnya error, tapi tidak ada error %v", err)
	}
	if hasil != 0 {
		t.Errorf("Bagi(10, 0) = %d, tidak dapat dibagi 0", hasil)
	}
}

func TestBagiTable(t *testing.T) {
	tests := []struct {
		nama     string
		a, b     int
		mau      int
		mauError bool
	}{
		{"bagi normal", 10, 2, 5, false},
		{"bagi nol", 10, 0, 0, true},
		{"bagi habis", 10, 3, 3, false},
	}

	for _, tc := range tests {
		t.Run(tc.nama, func(t *testing.T) {
			hasil, err := Bagi(tc.a, tc.b)
			if tc.mauError {
				if err == nil {
					t.Fatalf("harusnya error")
				}
				return
			}
			if err != nil {
				t.Fatalf("tidak boleh error: %v", err)
			}
			if hasil != tc.mau {
				t.Errorf("Bagi(%d, %d) = %d, mau %d", tc.a, tc.b, hasil, tc.mau)
			}
		})
	}
}
