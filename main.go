package main

import "fmt"

func hitung_umur(tahun int) int {
	return 2025 - tahun

}
func main() {
	fmt.Println("Starting...")
	umur := hitung_umur(1985)
	fmt.Println("Umur anda:", umur)

}
