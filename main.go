package main

import "fmt"

// Fungsi untuk menghitung luas kubus
func hitungLuasKubus(sisi float64) float64 {
	return 6 * sisi * sisi
}

func main() {
	sisi := 3.0
	luas := hitungLuasKubus(sisi)
	fmt.Printf("Luas kubus dengan sisi %.2f adalah %.2f\n", sisi, luas)
}
