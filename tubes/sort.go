package main

func selectionSortUmur(data []Pasien) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if data[j].Umur < data[minIdx].Umur {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
}

func gnomeSortBiaya(data []Pasien) {
	index := 0
	for index < len(data) {
		if index == 0 {
			index++
		}
		if data[index].Biaya >= data[index-1].Biaya {
			index++
		} else {
			data[index], data[index-1] = data[index-1], data[index]
			index--
		}
	}
}
