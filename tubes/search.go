package main

import (
	"fmt"
	"strings"
)

func linearSearchNama(data []Pasien, keyword string) {
	clearScreen()
	found := false
	for _, p := range data {
		if strings.EqualFold(p.Nama, keyword) {
			fmt.Print("Pasien ditemukan:")
			fmt.Printf("Nama: %s | Umur: %d | Penyakit: %s | Biaya: Rp%d\n\n", p.Nama, p.Umur, p.Penyakit, p.Biaya)
			found = true
			break
		}
	}
	if !found {
		fmt.Print("Pasien tidak ditemukan.\n")
	}
}

func binarySearchUmur(data []Pasien, target int) {
	clearScreen()
	if len(data) == 0 {
		fmt.Print("Tidak ada data pasien.\n")
		return
	}
	selectionSortUmur(data)
	low := 0
	high := len(data) - 1
	found := false

	for low <= high {
		mid := (low + high) / 2
		if data[mid].Umur == target {
			fmt.Print("Pasien dengan umur ditemukan:")
			fmt.Printf("Nama: %s | Umur: %d | Penyakit: %s | Biaya: Rp%d\n\n", data[mid].Nama, data[mid].Umur, data[mid].Penyakit, data[mid].Biaya)
			found = true
			break
		} else if data[mid].Umur < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if !found {
		fmt.Print("Tidak ada pasien dengan umur tersebut.\n")
	}
}
