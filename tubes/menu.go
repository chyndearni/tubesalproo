package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func menu() {
	for {
		clearScreen()
		fmt.Println("====== Menu Utama ======")
		fmt.Println("1. Input Data Pasien")
		fmt.Println("2. Tampilkan Semua Pasien")
		fmt.Println("3. Urutkan Berdasarkan Umur")
		fmt.Println("4. Urutkan Berdasarkan Biaya")
		fmt.Println("5. Cari Pasien Berdasarkan Nama")
		fmt.Println("6. Cari Pasien Berdasarkan Umur")
		fmt.Println("7. Input Data Dokter")
		fmt.Println("8. Tampilkan Semua Dokter")
		fmt.Println("9. Input Data Ruangan")
		fmt.Println("10. Tampilkan Semua Ruangan")
		fmt.Println("0. Exit")
		fmt.Println("Pilih menu: ")

		var pilihan string
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			inputPasien()
			pause()
		case "2":
			tampilkanPasien()
			pause()
		case "3":
			selectionSortUmur(daftarPasien)
			fmt.Print("Diurutkan berdasarkan umur!\n")
			pause()
		case "4":
			gnomeSortBiaya(daftarPasien)
			fmt.Print("Diurutkan berdasarkan biaya!\n")
			pause()
		case "5":
			fmt.Print("Masukkan nama yang dicari: ")
			reader := bufio.NewReader(os.Stdin)
			nama, _ := reader.ReadString('\n')
			nama = strings.TrimSpace(nama)
			linearSearchNama(daftarPasien, nama)
			pause()
		case "6":
			fmt.Print("Masukkan umur yang dicari: ")
			var umurStr string
			fmt.Scanln(&umurStr)
			umur, err := strconv.Atoi(umurStr)
			if err == nil {
				binarySearchUmur(daftarPasien, umur)
			} else {
				fmt.Print("Input umur tidak valid.\n")
			}
			pause()
		case "7":
			inputDokter()
			pause()
		case "8":
			tampilkanDokter()
			pause()
		case "9":
			inputRuangan()
			pause()
		case "10":
			tampilkanRuangan()
			pause()
		case "0":
			fmt.Print("Keluar dari program. Sampai jumpa!")
			return
		default:
			fmt.Print("Pilihan tidak valid. Coba lagi.\n")
			pause()
		}
	}
}
