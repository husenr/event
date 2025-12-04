package main

import (
	"fmt"
	"time"
)

func hitung_umur(tanggal_lahir time.Time) int {
	t := time.Now()

	year := tanggal_lahir.Year()
	selisih_tahun := t.Year() - year

	return selisih_tahun
}
func pengurangan(angka1 int, angka2 int) int {

	return angka1 - angka2

}
func main() {
	var tanggal_lahir string
	var rata_umur int
	layout := "2006-01-02"

	fmt.Println("Starting...")
	fmt.Println("masukkan tanggal lahir :")
	fmt.Scanln(&tanggal_lahir)
	tanggal, err := time.Parse(layout, tanggal_lahir)
	if err != nil {
		fmt.Println("Error Parsing:", err)

	}
	umur := hitung_umur(tanggal)
	fmt.Println("Umur anda:", umur, "Tahun")
	fmt.Print("Masukkan rata-rata umur :")
	fmt.Scan(&rata_umur)
	fmt.Println("Jika rata-rata umur = ", rata_umur, "tahun, maka...")
	sisa := pengurangan(rata_umur, umur)
	fmt.Println("Sisa umur anda :", sisa, "tahun")

}
