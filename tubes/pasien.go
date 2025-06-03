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

var daftarPasien []Pasien

func inputPasien() {
	clearScreen()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Nama Pasien: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	var umur int
	for {
		fmt.Print("Masukkan Umur Pasien: ")
		fmt.Scanln(&umur)
		if umur < 0 {
			fmt.Println("Umur tidak boleh negatif")
		} else {
			break
		}
	}

	fmt.Print("Masukkan Penyakit: ")
	penyakit, _ := reader.ReadString('\n')
	penyakit = strings.TrimSpace(penyakit)

	var biaya int
	for {
		fmt.Print("Masukkan Biaya: ")
		fmt.Scanln(&biaya)
		if biaya >= 1000 {
			break
		}
		fmt.Println("Biaya terlalu kecil atau tidak valid. Harus minimal Rp1000.")
	}

	baru := Pasien{Nama: nama, Umur: umur, Penyakit: penyakit, Biaya: biaya}
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
