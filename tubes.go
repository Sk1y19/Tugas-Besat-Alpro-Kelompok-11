package main

import "fmt"

type BahanMakanan struct {
	nama                        string
	stok, tanggal, bulan, tahun int
}

const NMAX int = 1000

type tabMakanan [NMAX]BahanMakanan

func main() {
	var data tabMakanan
	var nData, pilihmenu int
	var nama, x string
	menu()
	for pilihmenu != 8 {
		fmt.Print("Pilih opsi dari nomor 1 - 8: ")
		fmt.Scanln(&pilihmenu)
		switch pilihmenu {
		case 1:
			fmt.Print("Masukkan banyaknya data bahan makanan yang akan dimasukkan: ")
			fmt.Scan(&nData)
			inputBahan(&data, &nData) 
		case 2:
			ubahData(&data, &nData)
		case 3:
			hapusData(&data, nama, &nData)
		case 4:
			kadaluarsa(&data, &nData)
		case 5:
			fmt.Scan(&x)
			cariBahan(data, nData, x)
		case 6:
			kadaluarsa(&data, &nData)
		case 7:
			InsertionSort(&data, nData)
			cetakBahan(data, nData)
		case 8:
			fmt.Println("Berhasil Log Out")
		}
	}
}

func menu() {
	fmt.Println("------MENU------")
	fmt.Println("1. Input data makanan")
	fmt.Println("2. Edit data makanan")
	fmt.Println("3. Hapus data makanan")
	fmt.Println("4. Cek Kadaluarsa")
	fmt.Println("5. Cari bahan makanan")
	fmt.Println("6. Daftar bahan makanan")
	fmt.Println("7. Cek kadaluarsa berdasarkan tanggal")
	fmt.Println("8. Exit")
}

func inputBahan(T *tabMakanan, n *int) {
	for i := 0; i < *n; i++ {
		fmt.Print("Masukkan nama bahan: ")
        fmt.Scan(&T[i].nama)
        fmt.Print("Masukkan stok: ")
        fmt.Scan(&T[i].stok)
        fmt.Print("Masukkan tanggal kadaluarsa (dd mm yy): ")
        fmt.Scan(&T[i].tanggal, &T[i].bulan, &T[i].tahun)
	}
}

func cetakBahan(T tabMakanan, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Nama: %s\n", T[i].nama)
        fmt.Printf("Stok: %d\n", T[i].stok)
        fmt.Printf("Tanggal Kadaluarsa (dd - mm - yy): %02d - %02d - %02d\n", T[i].tanggal, T[i].bulan, T[i].tahun)
	}
}

func hapusData(tab *tabMakanan, nama string, n *int) {
	var i, idx int
	idx = -1
	for i = 0; i <= *n && idx == -1; i++ {
		if tab[i].nama == nama {
			idx = i
		}
	}
	for i = idx; i <= *n; i++ {
		tab[i] = tab[i+1]
	}
	*n -= 1
}

func ubahData(tab *tabMakanan, n *int) {
	var i, idx int
	var nama string
	idx = -1
	fmt.Print("Silakan masukkan nama bahan yang ingin diubah datanya: ")
	fmt.Scan(&nama)
	for i = 0; i <= *n && idx == -1; i++ {
		if tab[i].nama == nama {
			idx = i
		}
	}
	if idx == -1 {
		fmt.Print("Data tidak ditemukan.")
	} else {
		fmt.Print("Silakan masukkan tanggal kadaluarsa yang baru: (dd mm yy)")
		fmt.Scan(&tab[idx].tanggal, &tab[idx].bulan, &tab[idx].tahun)
	}
}

func kadaluarsa(tab *tabMakanan, n *int) {
	var i, tanggal, bulan, tahun int
	fmt.Print("Silakan masukkan tanggal saat ini: (dd mm yy)")
	fmt.Scan(&tanggal, &bulan, &tahun)

	for i = 0; i <= *n; i++ {
		if ((tab[i].tanggal-tanggal <= -3 && tab[i].tanggal-tanggal <= 3) && tab[i].bulan == bulan && tab[i].tahun == tahun) || ((tab[i].bulan-bulan == 1 || bulan-tab[i].bulan == 1) && tab[i].tahun == tahun && (tab[i].tanggal-tanggal >= 25 || tanggal-tab[i].tanggal-tanggal >= 25)) {
			fmt.Print("PERINGATAN: ", tab[i].nama, " AKAN SEGERA KADALUARSA PADA TANGGAL ", tab[i].tanggal, "/", tab[i].bulan, "/", tab[i].tahun)
		}
	}
}

func cariBahan(T tabMakanan, n int, x string) {
	var ketemu bool
	var i int
	ketemu = false
	for !ketemu && i < n {
		if T[i].nama == x {
			fmt.Println("Data Ditemukan")
			fmt.Printf("Nama bahan makanan: %s", T[i].nama)
			fmt.Printf("Stok : %d", T[i].stok)
			fmt.Printf("Tanggal Kadaluarsa (dd - mm - yy): ", T[i].tanggal, T[i].bulan, T[i].tahun)
		} else {
			fmt.Println("Data tidak ditemukan")
		}
	}
}

func InsertionSort(T *tabMakanan, n int) {
	var pass, i int
	var temp BahanMakanan

	pass = 1
	for pass <= n-1 {
		i = pass
		temp.tanggal = T[pass].tanggal
		for i > 0 && temp.tanggal < T[i-1].tanggal {
			T[i] = T[i-1]
			i--
		}
		T[i].tanggal = temp.tanggal
		pass++
	}
}
