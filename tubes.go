package main

import "fmt"

type Bahanmakanan struct {
	nama                        string
	stok, tanggal, bulan, tahun int
}

const NMAX int = 1000

type tabMakanan [NMAX]BahanMakanan

func main() {
	var data tabMakanan
	var nData, pilihmenu int
	var nama string
	for pilihmenu != 5 {
		menu()
		fmt.Print("Pilih opsi nomor 1 - 7: ")
		fmt.Scan(&pilihmenu)
		switch pilihmenu {
		case 1:
			inputBahan(&data, &nData)
		case 2:
			ubahData(&data, &nData)
		case 3:
			hapusData(&data, nama, &nData)
		case 4:
			kadaluarsa(data)
		case 5:
			cariBahan()
		case 6:
			kadaluarsa(data)
		case 7:
			fmt.Print("Logout")
		}
	}
}

func menu() {
	fmt.Print("------MENU------")
	fmt.Print("1. Input data makanan")
	fmt.Print("2. Edit data makanan")
	fmt.Print("3. Hapus data makanan")
	fmt.Print("4. Cek Kadaluarsa")
	fmt.Print("5. Cari bahan makanan")
	fmt.Print("6. Daftar bahan makanan")
	fmt.Print("7. Exit")
}

func inputBahan(T *tabMakanan, n *int) {
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		fmt.Print("Masukkan nama bahan:")
		fmt.Scan(&T[i].nama, "\n")
		fmt.Print("Masukkan tanggal kadaluarsa (dd mm yy):")
		fmt.Scan(&T[i].tanggal, &T[i].bulan, &T[i].tahun, "\n")
	}
}

func cetakBahan(T *tabMakanan, n *int) {
	for i := 0; i < *n; i++ {
		fmt.Print("Nama: ")
		fmt.Printf(T[i].nama, "\n")
		fmt.Print("Tanggal Kadaluarsa (dd - mm - yy):")
		fmt.Print(T[i].tanggal, T[i].bulan, T[i].tahun, "\n")
	}
}

func hapusData(tab *tabMakanan, nama string, n *int) {
	var i, idx int
	idx = -1
	for i = 0; i <= *n && idx == -1; i++ {
		if *tab[i].nama == nama {
			idx = i
		}
	}
	for i = idx; i <= *n; i++ {
		*tab[i] = *tab[i+1]
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
		if *tab[i].nama == nama {
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

	for i = 0; i <= n; i++ {
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

func binarySearch() {

}
