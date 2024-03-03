package main

import (
	"fmt"
	"os"
	"strconv"
)

// Struct untuk biodata
type Digitalent struct {
    Nama string
    Alamat string
    Pekerjaan string
    AlasanMemilih string
}

func main() {
	// Array berisikan data yang sudah didefinisikan pada struct
	digitalent := []Digitalent{
        {"Budi Setiawan", "Jakarta", "Software Engineer", "Saya ingin menguasai Golang untuk pengembangan backend."},
        {"Siti Rahayu", "Surabaya", "Data Analyst", "Saya tertarik dengan performa dan efisiensi Golang."},
        {"Muhammad Zhikri", "Langsa", "Back-end Web Developer", "Golang cocok untuk pengembangan aplikasi berbasis microservice."},
        {"Dewi Susanti", "Medan", "System Administrator", "Dukungan concurrency di Golang sangat menarik bagi saya."},
        {"Arief Pratama", "Semarang", "Full Stack Developer", "Saya ingin mengembangkan aplikasi web dengan Golang."},
    }

	// Memeriksa argument nomor absen
	if len(os.Args) < 2 {
		fmt.Println("Tidak ada nomor absen yang diinput")
		return
	}

	// Mengambil indeks dari argumen yang diberikan
	index, err := strconv.Atoi(os.Args[1])

	// Nomor absen invalid jika input nil, kurang dari 1, dan lebih dari jumlah data
	if err != nil || index < 1 || index > len(digitalent) {
		fmt.Println("Nomor absen invalid")
		return
	}

	// Menampilkan data sesuai indeks
	displayTalent(digitalent[index-1])
}

// Fungsi untuk menampilkan data sesuai nomor absensi
func displayTalent(digitalent Digitalent) {
	fmt.Println("Biodata Talent")
	fmt.Println("------------------------------------------------------")
	fmt.Printf("Nama: %s\n", digitalent.Nama)
	fmt.Printf("Alamat: %s\n", digitalent.Alamat)
	fmt.Printf("Pekerjaan: %s\n", digitalent.Pekerjaan)
	fmt.Printf("Alasan Memilih Golang: %s\n", digitalent.AlasanMemilih)
}