package main

import "fmt"

// Definisikan tipe data untuk jenis ban
type Ban string

const (
	Karet Ban = "karet"
	Kayu  Ban = "kayu"
	Besi  Ban = "besi"
)

// Definisikan struct untuk mobil
type Mobil struct {
	Roda                  [4]Ban
	PintuKanan, PintuKiri string
}

// Method untuk mengganti ban pada mobil
func (m *Mobil) GantiBan(posisi int, jenis Ban) {
	m.Roda[posisi] = jenis
}

// Method untuk mengetuk pintu
func (m Mobil) KetukPintu() {
	fmt.Printf("Ketuk Pintu Kanan: %s\n", m.PintuKanan)
	fmt.Printf("Ketuk Pintu Kiri: %s\n", m.PintuKiri)
}

// Method untuk membuka pintu
func (m Mobil) BukaPintu() {
	fmt.Printf("Buka Pintu Kanan: %s\n", m.PintuKanan)
	fmt.Printf("Buka Pintu Kiri: %s\n", m.PintuKiri)
}

func main() {
	// Inisialisasi mobil
	mobil := Mobil{
		Roda:       [4]Ban{Karet, Karet, Karet, Karet},
		PintuKanan: "Knock",
		PintuKiri:  "Beep",
	}

	// Tampilkan status roda sebelum penggantian
	fmt.Println("Status Roda Mobil Sebelum Penggantian:")
	for i, ban := range mobil.Roda {
		fmt.Printf("Roda %d: %s\n", i+1, ban)
	}

	// Ganti roda pada posisi ke-2 dengan ban kayu
	mobil.GantiBan(1, Kayu)
	mobil.GantiBan(3, Besi)

	// Tampilkan status roda setelah penggantian
	fmt.Println("\nStatus Roda Mobil Setelah Penggantian:")
	for i, ban := range mobil.Roda {
		fmt.Printf("Roda %d: %s\n", i+1, ban)
	}

	// Ketuk dan buka pintu
	fmt.Println("\nSuara Pintu:")
	mobil.KetukPintu()
	mobil.BukaPintu()
}
