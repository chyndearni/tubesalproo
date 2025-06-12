package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pasien struct {
	Nama     string
	Umur     int
	Penyakit string
	Biaya    int
}
type Ruangan struct {
	NamaRuangan string
	Tipe        string
	Kapasitas   int
}

var daftarPasien []Pasien
var daftarRuangan []Ruangan

func inputPasien() {
	clearScreen()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Nama Pasien: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	var umur int
	valid := false
	for !valid {
		fmt.Print("Masukkan Umur Pasien: ")
		fmt.Scanln(&umur)
		if umur < 1 {
			fmt.Println("Umur tidak valid")
		} else {
			valid = true
		}
	}

	fmt.Print("Masukkan Penyakit: ")
	penyakit, _ := reader.ReadString('\n')
	penyakit = strings.TrimSpace(penyakit)

	var biaya int
	validBiaya := false
	for !validBiaya {
		fmt.Print("Masukkan Biaya: ")
		fmt.Scanln(&biaya)
		if biaya <= 1000 {
			fmt.Println("Biaya harus lebih dari 1000. Silakan masukkan lagi.")
		} else {
			validBiaya = true
		}
	}

	baru := Pasien{
		Nama:     nama,
		Umur:     umur,
		Penyakit: penyakit,
		Biaya:    biaya,
	}
	daftarPasien = append(daftarPasien, baru)
	fmt.Print("Pasien Berhasil Ditambahkan!\n")
}

func tampilkanPasien() {
	clearScreen()
	if len(daftarPasien) == 0 {
		fmt.Print("Belum ada data pasien.\n")
		return
	}
	fmt.Println("=== Daftar Pasien ===")
	for i, p := range daftarPasien {
		fmt.Printf("%d. Nama: %s | Umur: %d | Penyakit: %s | Biaya: Rp%d\n", i+1, p.Nama, p.Umur, p.Penyakit, p.Biaya)
	}
	fmt.Print()
}
func inputRuangan() {
	clearScreen()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Nama Ruangan: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan Tipe Ruangan: ")
	tipe, _ := reader.ReadString('\n')
	tipe = strings.TrimSpace(tipe)

	var kapasitas int
	for {
		fmt.Print("Masukkan Kapasitas Ruangan: ")
		fmt.Scanln(&kapasitas)
		if kapasitas <= 0 {
			fmt.Println("Kapasitas harus lebih dari 0.")
		} else {
			break
		}
	}

	ruanganBaru := Ruangan{NamaRuangan: nama, Tipe: tipe, Kapasitas: kapasitas}
	daftarRuangan = append(daftarRuangan, ruanganBaru)
	fmt.Println("Data Ruangan Berhasil Ditambahkan!")
}

func tampilkanRuangan() {
	clearScreen()
	if len(daftarRuangan) == 0 {
		fmt.Println("Belum ada data ruangan.")
		return
	}
	fmt.Println("=== Daftar Ruangan ===")
	for i, r := range daftarRuangan {
		fmt.Printf("%d. Nama: %s | Tipe: %s | Kapasitas: %d\n", i+1, r.NamaRuangan, r.Tipe, r.Kapasitas)
	}
}
