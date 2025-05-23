package main

import "fmt"

type BahanMakanan struct {
	nama, id                    string
	stok, tanggal, bulan, tahun int
}

const NMAX int = 1000

type tabMakanan [NMAX]BahanMakanan

func main() {
	var data tabMakanan
	var nData, n, pilihmenu int
	var id string
	nData = 0
	for pilihmenu != 9 {
		menu()
		fmt.Print("Pilih opsi dari nomor 1 - 9q: ")
		fmt.Scan(&pilihmenu)
		switch pilihmenu {
		case 1:
			fmt.Print("Masukkan banyaknya data bahan makanan yang akan dimasukkan: ")
			fmt.Scan(&n)
			inputBahan(&data, n, &nData)
		case 2:
			ubahData(&data, &nData, id)
		case 3:
			fmt.Print("Masukkan ID bahan makanan: ")
			fmt.Scan(&id)
			hapusData(&data, id, &nData)
		case 5:
			kadaluarsa(&data, &nData)
		case 6:
			fmt.Print("Masukkan ID bahan makanan: ")
			fmt.Scan(&id)
			cariBahan(data, nData, id)
		case 4:
			cetakBahan(data, nData)
		case 7:
			cetakBahan(InsertionSort(data, nData), nData)
		case 8:
			cetakBahan(selectionSort(data, nData), nData)
		case 9:
			fmt.Println("Log out berhasil")
		default:
			fmt.Println("Pilihan harus 1 - 9")
		}
	}
}

func menu() {
	fmt.Println("------MENU------")
	fmt.Println("1. Input data makanan")
	fmt.Println("2. Edit data makanan")
	fmt.Println("3. Hapus data makanan")
	fmt.Println("4. Daftar Bahan Makanan")
	fmt.Println("5. Cek Bahan Yang Mendekati Kadaluarsa atau Telah Kadaluarsa")
	fmt.Println("6. Cari Bahan Makanan Berdasarkan ID")
	fmt.Println("7. Daftar Bahan Makanan Urut Secara Menaik Berdasarkan Tanggal Kadaluarsa")
	fmt.Println("8. Daftar Bahan Makanan Urut Secara Menurun Berdasarkan Jumlah Stok")
	fmt.Println("9. Exit")
}

func inputBahan(T *tabMakanan, n int, nData *int) {
	var temp = *nData
	var isTrue bool
	isTrue = false
	for i := temp; i < temp+n; i++ {
		fmt.Print("Masukkan Id bahan makanan: ")
		fmt.Scan(&T[i].id)
		fmt.Print("Masukkan nama bahan (jangan gunakan spasi gunakan '_'): ")
		fmt.Scan(&T[i].nama)
		fmt.Print("Masukkan stok: ")
		fmt.Scan(&T[i].stok)
		for !isTrue {
			fmt.Print("Masukkan tanggal kadaluarsa (dd m yyyy (untuk penulisan tanggal dan bulan 1 digit tidak perlu menggunkan 0 didepannya)): ")
			fmt.Scan(&T[i].tanggal, &T[i].bulan, &T[i].tahun)
			if T[i].bulan >= 1 && T[i].bulan <= 12 {
					isTrue = true
				if T[i].bulan == 1 || T[i].bulan == 3 || T[i].bulan == 5 || T[i].bulan == 7 || T[i].bulan == 9 || T[i].bulan == 11 {
					if T[i].tanggal >= 1 && T[i].tanggal <= 31 {
						isTrue = true
					} else {
						fmt.Println("Tanggal tidak valid")
						isTrue = false
					}
				} else if T[i].bulan == 4 || T[i].bulan == 6 || T[i].bulan == 8 || T[i].bulan == 10 || T[i].bulan == 12 {
					if T[i].tanggal >= 1 && T[i].tanggal <= 30 {
						isTrue = true
					} else {
						fmt.Println("Tanggal tidak valid")
						isTrue = false
					}
				} else if T[i].bulan == 2 {
					if T[i].tanggal >= 1 && T[i].tanggal <= 28 {
						isTrue = true
					} else {
						fmt.Println("Tanggal tidak valid")
						isTrue = false
					}
				}
			} else {
				fmt.Println("Bulan tidak valid")
				isTrue = false
				
			}
		}
		*nData++
	}
}

func cetakBahan(T tabMakanan, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("---\n")
		fmt.Printf("id : %s\n", T[i].id)
		fmt.Printf("Nama : %s\n", T[i].nama)
		fmt.Printf("Stok : %d\n", T[i].stok)
		fmt.Printf("Tanggal Kadaluarsa : %02d - %02d - %02d\n", T[i].tanggal, T[i].bulan, T[i].tahun)
	}
}

func hapusData(tab *tabMakanan, id string, n *int) {
	var i, idx int
	idx = -1
	for i = 0; i <= *n && idx == -1; i++ {
		if tab[i].id == id {
			idx = i
		}
	}
	for i = idx; i <= *n; i++ {
		tab[i] = tab[i+1]
	}
	*n -= 1
}

func ubahData(tab *tabMakanan, n *int, id string) {
	var i, idx int
	idx = -1
	fmt.Print("Silakan masukkan id bahan yang ingin diubah datanya: ")
	fmt.Scan(&id)
	for i = 0; i <= *n && idx == -1; i++ {
		if tab[i].id == id {
			idx = i
		}
	}
	if idx == -1 {
		fmt.Print("Id Data tidak ditemukan.\n")
	} else {
		fmt.Print("Masukkan nama yang baru (jika tidak ingin diubah gunakan nama yang sudah ada): ")
		fmt.Scan(&tab[idx].nama)
		fmt.Print("Ubah stok yang baru (kalau sama pakai jumlah stok lama): ")
		fmt.Scan(&tab[idx].stok)
		fmt.Print("Silakan masukkan tanggal kadaluarsa yang baru (dd mm yy): ")
		fmt.Scan(&tab[idx].tanggal, &tab[idx].bulan, &tab[idx].tahun)
	}
}

func kadaluarsa(tab *tabMakanan, n *int) {
	var i, tanggal, tahun, bulan int
	var found bool
	fmt.Print("Silakan masukkan tanggal saat ini: (dd mm yyyy (untuk penulisan tanggal dan bulan 1 digit tidak perlu menggunkan 0 didepannya)): ")
	fmt.Scan(&tanggal, &bulan, &tahun)
	found = false
	for i = 0; i <= *n; i++ {
		if ((tab[i].tanggal-tanggal <= 3 && tab[i].tanggal-tanggal >= 0) && tab[i].bulan == bulan && tab[i].tahun == tahun) || ((tab[i].bulan-bulan == 1) && tab[i].tahun == tahun && ((tab[i].tanggal <= 3 && tanggal >= 26) || (tab[i].tanggal >= 26 && tanggal <= 3))) {
			fmt.Print("PERINGATAN: ", tab[i].nama, " AKAN SEGERA KADALUARSA PADA TANGGAL ", tab[i].tanggal, "/", tab[i].bulan, "/", tab[i].tahun, "\n")
			found = true
		} else if tanggal-tab[i].tanggal > 0 && bulan-tab[i].bulan >= 0 && tahun-tab[i].tahun >= 0 && tab[i].tanggal != 0 && tanggal != 0 {
			fmt.Print("PERINGATAN: ", tab[i].nama, " TELAH KADALUARSA PADA TANGGAL ", tab[i].tanggal, "/", tab[i].bulan, "/", tab[i].tahun, "\n")
			found = true
		}
	}
	if found == false {
		fmt.Println("Tidak ada bahan makanan yang mendekati tanggal kadaluarsa atau telah kadaluarsa")
	}
}

func cariBahan(T tabMakanan, n int, id string) {
	var ketemu bool
	ketemu = false
	for i := 0; i < n; i++ {
		if T[i].id == id {
			fmt.Println("Data Ditemukan")
			fmt.Println("--------------")
			fmt.Printf("Nama bahan makanan: %s\n", T[i].nama)
			fmt.Printf("Stok: %d\n", T[i].stok)
			fmt.Printf("Tanggal Kadaluarsa (dd - mm - yyyy): %02d - %02d - %04d\n", T[i].tanggal, T[i].bulan, T[i].tahun)
			ketemu = true
		}
	}
	if ketemu == false {
		fmt.Println("Data tidak ditemukan")
	}
}

func InsertionSort(T tabMakanan, n int) tabMakanan {
	var pass, i int
	var temp BahanMakanan

	pass = 1
	for pass <= n-1 {
		i = pass
		temp.tanggal = T[pass].tanggal
		for i > 0 && temp.tanggal < T[i-1].tanggal && temp.bulan <= T[i-1].bulan && temp.tahun <= T[i-1].tahun {
			T[i] = T[i-1]
			i--
		}
		T[i].tanggal = temp.tanggal
		pass++
	}
	return T
}

func selectionSort(T tabMakanan, n int)tabMakanan{
	var i, idx, pass int
	var temp int
	
	pass = 1
	for pass < n{
		idx = pass - 1
		i = pass
		for i < n{
			if T[i].stok > T[idx].stok{
				idx = i
			}
			i++
		}
		temp = T[pass - 1].stok
		T[pass - 1].stok = T[idx].stok
		T[idx].stok = temp
		pass++
	}
	return T
}
