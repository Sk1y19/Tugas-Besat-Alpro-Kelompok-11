/*Note: program masih belum selesai masih ada fitur yang perlu di perbaiki*/

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
	var id, x string
	nData = 0
	for pilihmenu != 8 {
		menu()
		fmt.Print("Pilih opsi dari nomor 1 - 8: ")
		fmt.Scanln(&pilihmenu)
		switch pilihmenu {
		case 1:
			fmt.Print("Masukkan banyaknya data bahan makanan yang akan dimasukkan: ")
			fmt.Scan(&n)
			inputBahan(&data, n, &nData)
		case 2:
			ubahData(&data, &nData, id)
		case 3:
			fmt.Print("Masukkan id makanan: ")
			fmt.Scan(&id)
			hapusData(&data, id, &nData)
		case 4:
			kadaluarsa(&data, &nData)
		case 5:
			fmt.Scan(&x)
			cariBahan(data, nData, x)
		case 6:
			cetakBahan(data, nData)
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
		fmt.Print("Masukkan tanggal kadaluarsa (dd m yyyy (untuk penulisan bulan 1 digit tidak perlu menggunkan 0 didepannya)): ")
		fmt.Scan(&T[i].tanggal, &T[i].bulan, &T[i].tahun)
		for !isTrue {
			if T[i].bulan >= 1 && T[i].bulan <= 2 {
				if T[i].bulan == 1 || T[i].bulan == 3 || T[i].bulan == 5 || T[i].bulan == 7 || T[i].bulan == 9 || T[i].bulan == 11 {
					if T[i].tanggal >= 01 && T[i].tanggal <= 31 {
						isTrue = true
					} else {
						fmt.Println("Tanggal tidak valid")
						fmt.Print("Masukkan tanggal kadaluarsa (dd m yyyy (untuk penulisan bulan 1 digit tidak perlu menggunkan 0 didepannya)): ")
						fmt.Scan(&T[i].tanggal, &T[i].bulan, &T[i].tahun)
					}
				} else if T[i].bulan == 4 || T[i].bulan == 6 || T[i].bulan == 8 || T[i].bulan == 10 || T[i].bulan == 12 {
					if T[i].tanggal >= 1 && T[i].tanggal <= 30 {
						isTrue = true
					} else {
						fmt.Println("Tanggal tidak valid")
						fmt.Print("Masukkan tanggal kadaluarsa (dd m yyyy (untuk penulisan bulan 1 digit tidak perlu menggunkan 0 didepannya)): ")
						fmt.Scan(&T[i].tanggal, &T[i].bulan, &T[i].tahun)
					}
				} else if T[i].bulan == 2 {
					if T[i].tanggal >= 1 && T[i].tanggal <= 28 {
						isTrue = true
					} else {
						fmt.Println("Tanggal tidak valid")
						fmt.Print("Masukkan tanggal kadaluarsa (dd m yyyy (untuk penulisan bulan 1 digit tidak perlu menggunkan 0 didepannya)): ")
						fmt.Scan(&T[i].tanggal, &T[i].bulan, &T[i].tahun)
					}
				}
			} else {
				fmt.Println("Bulan tidak valid")
				fmt.Print("Masukkan tanggal kadaluarsa (dd m yyyy (untuk penulisan bulan 1 digit tidak perlu menggunkan 0 didepannya)): ")
				fmt.Scan(&T[i].tanggal, &T[i].bulan, &T[i].tahun)
			}
		}
		*nData++
	}
}

func cetakBahan(T tabMakanan, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("id : %s\n", T[i].id)
		fmt.Printf("Nama: %s\n", T[i].nama)
		fmt.Printf("Stok: %d\n", T[i].stok)
		fmt.Printf("Tanggal Kadaluarsa (dd - mm - yy): %02d - %02d - %02d\n", T[i].tanggal, T[i].bulan, T[i].tahun)
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
		fmt.Print("Id Data tidak ditemukan.")
	} else {
		fmt.Print("Masukkan nama yang baru (jika tidak ingin diubah gunakan nama yang sudah ada): ")
		fmt.Scan(&tab[idx].nama)
		fmt.Print("Ubah stok yang baru (kalau sama pakai jumlah stok lama): ")
		fmt.Scan(&tab[idx].stok)
		fmt.Print("Silakan masukkan tanggal kadaluarsa yang baru: (dd mm yy)")
		fmt.Scan(&tab[idx].tanggal, &tab[idx].bulan, &tab[idx].tahun)
	}
}

func kadaluarsa(tab *tabMakanan, n *int) {
	var i, tanggal, tahun, bulan int
	var found bool
	fmt.Print("Silakan masukkan tanggal saat ini: (dd mm yyyy): ")
	fmt.Scan(&tanggal, &bulan, &tahun)
	found = false
	for i = 0; i <= *n; i++ {
		if ((tab[i].tanggal-tanggal <= -3 && tab[i].tanggal-tanggal <= 3) && tab[i].bulan == bulan && tab[i].tahun == tahun) || ((tab[i].bulan-bulan == 1 || bulan-tab[i].bulan == 1) && tab[i].tahun == tahun && (tab[i].tanggal-tanggal >= 25 || tanggal-tab[i].tanggal-tanggal >= 25)) {
			fmt.Print("PERINGATAN: ", tab[i].nama, " AKAN SEGERA KADALUARSA PADA TANGGAL ", tab[i].tanggal, "/", tab[i].bulan, "/", tab[i].tahun)
			found = true
		}
	}
	if found == false {
		fmt.Println("Tidak ada bahan makanan yang mendekari tanggal kadaluarsa")
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

