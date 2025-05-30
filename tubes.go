package main

import (
	"fmt"
	"time"
)

type BahanMakanan struct {
	nama, id                                    string
	stok, stokTergunakan, tanggal, bulan, tahun int
}

const NMAX int = 10

type tabMakanan [NMAX]BahanMakanan

func main() {
	var data tabMakanan
	var nData, n, pilihmenu int
	var id string
	nData = 0
	for pilihmenu != 11 {
		menu()
		fmt.Print("Pilih opsi dari nomor 1 - 11: ")
		fmt.Scan(&pilihmenu)
		switch pilihmenu {
		case 1:
			fmt.Print("Masukkan banyaknya data bahan makanan yang akan dimasukkan: ")
			fmt.Scan(&n)
			for n <= 0{
				fmt.Println("MASUKAN TIDAK BOLEH SAMA DENGAN 0 ATAU NEGATIF!!")
				fmt.Print("Masukkan banyaknya data bahan makanan yang akan dimasukkan: ")
				fmt.Scan(&n)
			}
			inputBahan(&data, n, &nData)

		case 2:
			ubahData(&data, &nData)
		case 3:
			fmt.Print("Masukkan ID bahan makanan: ")
			fmt.Scan(&id)
			hapusData(&data, id, &nData)
		case 5:
			kadaluarsa(&data, &nData)
		case 4:
			cetakBahan(data, nData)
		case 6:
			fmt.Print("Masukkan Nama Bahan Makanan: ")
			fmt.Scan(&id)
			cariBahan(data, nData, id)
		case 7:
			var x string
			var bin int
			var sortId tabMakanan
			fmt.Print("Masukkan ID Bahan Makanan: ")
			fmt.Scan(&x)
			sortId = sortById(data, nData)

			bin = binarySearchId(sortId, x, nData)
			if bin != -1 {
				fmt.Println("Data Ditemukan")
				fmt.Println("----------------------------------")
				fmt.Printf("Nama bahan makanan: %s\n", sortId[bin].nama)
				fmt.Printf("Id bahan makanan:: %s\n", sortId[bin].id)
				fmt.Printf("Stok Yang Tersedia: %d\n", sortId[bin].stok-sortId[bin].stokTergunakan)
				fmt.Printf("Stok Awal: %d\n", sortId[bin].stok)
				fmt.Printf("Stok Tergunakan: %d\n", sortId[bin].stokTergunakan)
				fmt.Printf("Tanggal Kadaluarsa (dd - mm - yyyy): %02d - %02d - %04d\n", sortId[bin].tanggal, sortId[bin].bulan, sortId[bin].tahun)
			} else {
				fmt.Println("DATA TIDAK DITEMUKAN")
			}
		case 8:
			cetakBahan(InsertionSort(data, nData), nData)
		case 9:
			cetakBahan(selectionSort(data, nData), nData)
		case 10:
			stokTersedikit(data, nData)
			stokTerbanyak(data, nData)
		case 11:
			fmt.Println("Log out berhasil")
		default:
			fmt.Println("Pilihan harus 1 - 11")
		}
	}
}

func menu() {
	fmt.Println("================")
	fmt.Println("      Menu      ")
	fmt.Println("================")
	fmt.Println("1. Input data makanan")
	fmt.Println("2. Edit data makanan")
	fmt.Println("3. Hapus data makanan")
	fmt.Println("4. Daftar Bahan Makanan")
	fmt.Println("5. Cek Bahan Yang Mendekati Kadaluarsa atau Telah Kadaluarsa")
	fmt.Println("6. Cari Bahan Makanan Berdasarkan Nama")
	fmt.Println("7. Cari Bahan Makanan Berdasarkan ID")
	fmt.Println("8. Daftar Bahan Makanan Urut Secara Menaik Berdasarkan Tanggal Kadaluarsa")
	fmt.Println("9. Daftar Bahan Makanan Urut Secara Menurun Berdasarkan Jumlah Stok")
	fmt.Println("10. Cetak Bahan Makanan dengan stok Terbanyak dan Tersedikit")
	fmt.Println("11. Exit")
	fmt.Println("=========================================================================")
}

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func inputBahan(T *tabMakanan, n int, nData *int) {
	var temp = *nData
	var isTrue bool
	var isidExist, isidBahan, isStokvalid bool
	var idBahan string
	isTrue = false
	if *nData+n > NMAX {
		fmt.Println("PERINGATAN!!, data penuh!")
	} else {
		for i := temp; i < temp+n; i++ {
			isTrue = false
			isidBahan = false
			isStokvalid = false

			for !isidBahan {
				fmt.Print("Masukkan Id bahan makanan: ")
				fmt.Scan(&idBahan)
				isidExist = false
				for j := 0; j < *nData; j++ {
					if T[j].id == idBahan {
						isidExist = true
					}
				}
				if isidExist {
					fmt.Println("ID sudah ada, silakan masukkan ID yang berbeda.")
				} else {
					T[i].id = idBahan
					isidBahan = true
				}
			}
			fmt.Print("Masukkan nama bahan (jangan gunakan spasi gunakan '_'): ")
			fmt.Scan(&T[i].nama)

			for !isStokvalid {
				fmt.Print("Masukkan stok awal: ")
				fmt.Scan(&T[i].stok)
				if T[i].stok > 0 {
					isStokvalid = true
				} else {
					fmt.Println("MASUKKAN STOK TIDAK VALID!!")
				}
			}

			for !isTrue {
				fmt.Print("Masukkan tanggal kadaluarsa (dd m yyyy (untuk penulisan tanggal dan bulan 1 digit tidak perlu menggunkan 0 didepannya)): ")
				fmt.Scan(&T[i].tanggal, &T[i].bulan, &T[i].tahun)
				if T[i].bulan >= 1 && T[i].bulan <= 12 {
					if T[i].bulan == 1 || T[i].bulan == 3 || T[i].bulan == 5 || T[i].bulan == 7 || T[i].bulan == 8 || T[i].bulan == 10 || T[i].bulan == 12 {
						if T[i].tanggal >= 1 && T[i].tanggal <= 31 {
							isTrue = true
						} else {
							fmt.Println("Tanggal tidak valid")
						}
					} else if T[i].bulan == 4 || T[i].bulan == 6 || T[i].bulan == 9 || T[i].bulan == 11 {
						if T[i].tanggal >= 1 && T[i].tanggal <= 30 {
							isTrue = true
						} else {
							fmt.Println("Tanggal tidak valid")
						}
					} else if T[i].bulan == 2 {
						if T[i].tanggal >= 1 && T[i].tanggal <= 29 {
							if isLeapYear(T[i].tahun) {
								if T[i].tanggal >= 1 && T[i].tanggal <= 29 {
									isTrue = true
								}else{
									fmt.Println("Tanggal tidak valid")
								}
							} else {
								if T[i].tanggal >= 1 && T[i].tanggal <= 28 {
									isTrue = true
								}else{
									fmt.Println("Tanggal tidak valid")
								}
							}
						}else{
							fmt.Println("Tanggal tidak valid")
						}
					}
				} else {
					fmt.Println("Bulan tidak valid")
				}
			}
			*nData++
		}
	}
}

func cetakBahan(T tabMakanan, n int) {
	fmt.Println("======================")
	fmt.Println(" Daftar Bahan Makanan ")
	fmt.Println("======================")
	for i := 0; i < n; i++ {
		fmt.Printf("----------------\n")
		fmt.Printf("id : %s\n", T[i].id)
		fmt.Printf("Nama : %s\n", T[i].nama)
		fmt.Printf("Stok Yang Tersedia: %d\n", T[i].stok-T[i].stokTergunakan)
		fmt.Printf("Stok Awal: %d\n", T[i].stok)
		fmt.Printf("Stok Tergunakan: %d\n", T[i].stokTergunakan)
		fmt.Printf("Tanggal Kadaluarsa : %02d - %02d - %02d\n", T[i].tanggal, T[i].bulan, T[i].tahun)
	}
	if n == 0 {
		fmt.Println("Tidak Ada Data !")
	}
	fmt.Println("----------------")
}

func hapusData(tab *tabMakanan, id string, n *int) {
	var i, idx int
	idx = -1
	for i = 0; i < *n; i++ {
		if tab[i].id == id {
			idx = i
		}
	}
	if idx != -1 {
		for i = idx; i < *n-1; i++ {
			tab[i] = tab[i+1]
		}
		*n -= 1
		fmt.Println("Data telah berhasil dihapus.")
	} else {
		fmt.Println("Data tidak ditemukan!")
	}
}

func ubahData(tab *tabMakanan, n *int) {
	var i, idx, choice int
	var id string
	var isTrue bool
	isTrue = false
	choice = -1
	idx = -1
	fmt.Print("Silakan masukkan id bahan yang ingin diubah datanya: ")
	fmt.Scan(&id)
	for i = 0; i < *n && idx == -1; i++ {
		if tab[i].id == id {
			idx = i
		}
	}
	if idx == -1 {
		fmt.Print("Id Data tidak ditemukan.\n")
	} else {
		fmt.Println("================")
		fmt.Println("      Data      ")
		fmt.Println("================")
		fmt.Printf("id : %s\n", tab[idx].id)
		fmt.Printf("Nama : %s\n", tab[idx].nama)
		fmt.Printf("Stok Yang Tersedia: %d\n", tab[idx].stok-tab[idx].stokTergunakan)
		fmt.Printf("Stok Awal: %d\n", tab[idx].stok)
		fmt.Printf("Stok Tergunakan: %d\n", tab[idx].stokTergunakan)
		fmt.Printf("Tanggal Kadaluarsa : %02d - %02d - %02d\n", tab[idx].tanggal, tab[idx].bulan, tab[idx].tahun)
		fmt.Println("================")
		fmt.Println(" Menu Edit Data ")
		fmt.Println("================")
		fmt.Println("1. Ubah Nama Bahan Makanan")
		fmt.Println("2. Ubah Stok Awal")
		fmt.Println("3. Ubah Stok Tergunakan")
		fmt.Println("4. Ubah Tanggal Kadaluarsa")
		for choice < 1 || choice > 4 {
			fmt.Print("Pilih opsi 1/2/3/4: ")
			fmt.Scan(&choice)
			switch choice {
			case 1:
				fmt.Print("Masukkan nama yang baru: ")
				fmt.Scan(&tab[idx].nama)
			case 2:
				fmt.Print("Masukkan stok awal yang baru: ")
				fmt.Scan(&tab[idx].stok)
				for tab[idx].stok < tab[idx].stokTergunakan || tab[idx].stok < 0 {
					fmt.Print("STOK TIDAK VALID, SILAKAN INPUT ULANG: ")
					fmt.Scan(&tab[idx].stok)
				}
			case 3:
				fmt.Print("Masukkan stok tergunakan yang baru: ")
				fmt.Scan(&tab[idx].stokTergunakan)
				for tab[idx].stokTergunakan > tab[idx].stok || tab[idx].stokTergunakan < 0 {
					fmt.Print("STOK TERGUNAKAN TIDAK VALID, SILAKAN INPUT ULANG: ")
					fmt.Scan(&tab[idx].stokTergunakan)
				}
			case 4:
				for !isTrue {
					fmt.Print("Masukkan tanggal kadaluarsa (dd m yyyy (untuk penulisan tanggal dan bulan 1 digit tidak perlu menggunkan 0 didepannya)): ")
					fmt.Scan(&tab[idx].tanggal, &tab[idx].bulan, &tab[idx].tahun)
					if tab[idx].bulan >= 1 && tab[idx].bulan <= 12 {
						if tab[idx].bulan == 1 || tab[idx].bulan == 3 || tab[idx].bulan == 5 || tab[idx].bulan == 7 || tab[idx].bulan == 8 || tab[idx].bulan == 10 || tab[idx].bulan == 12 {
							if tab[idx].tanggal >= 1 && tab[idx].tanggal <= 31 {
								isTrue = true
							} else {
								fmt.Println("Tanggal tidak valid")
							}
						} else if tab[idx].bulan == 4 || tab[idx].bulan == 6 || tab[idx].bulan == 9 || tab[idx].bulan == 11 {
							if tab[idx].tanggal >= 1 && tab[idx].tanggal <= 30 {
								isTrue = true
							} else {
								fmt.Println("Tanggal tidak valid")
							}
						} else if tab[idx].bulan == 2 {
							if tab[idx].tanggal >= 1 && tab[idx].tanggal <= 29 {
								if isLeapYear(tab[idx].tahun) {
									if tab[idx].tanggal >= 1 && tab[idx].tanggal <= 29 {
										isTrue = true
									} else {
										fmt.Println("Tanggal tidak valid")
									}
								} else {
									if tab[idx].tanggal >= 1 && tab[idx].tanggal <= 28 {
										isTrue = true
									} else {
										fmt.Println("Tanggal tidak valid")
									}
								}
							}else{
								fmt.Println("Tanggal tidak valid")
							}
						}
					} else {
						fmt.Println("Bulan tidak valid")
					}
				}
			default:
				fmt.Println("Pilihan harus 1 - 4")
			}
		}
	}
}

func kadaluarsa(tab *tabMakanan, n *int) {
	var i int
	var found bool
	now := time.Now()
	tanggal := now.Day()
	bulan := int(now.Month())
	tahun := now.Year()
	found = false
	for i = 0; i < *n; i++ {
		if (tab[i].tanggal-now.Day() <= 3 && tab[i].tanggal-now.Day() >= 0 && tab[i].bulan == bulan && tab[i].tahun == tahun) ||
			(tab[i].bulan-bulan == 1 && tab[i].tahun == tahun && ((tab[i].tanggal <= 3 && tanggal >= 26) || (tab[i].tanggal >= 26 && tanggal <= 3))) {
			fmt.Printf("PERINGATAN: %s AKAN SEGERA KADALUARSA PADA TANGGAL %02d/%02d/%04d\n", tab[i].nama, tab[i].tanggal, tab[i].bulan, tab[i].tahun)
			found = true
		} else if (now.Day()-tab[i].tanggal > 0 && int(now.Month())-tab[i].bulan == 0 && now.Year()-tab[i].tahun >= 0 && tab[i].tanggal != 0 && now.Day() != 0) ||  (int(now.Month())-tab[i].bulan > 0 && now.Year()-tab[i].tahun == 0 && tab[i].tanggal != 0 && now.Day() != 0) || (now.Year()-tab[i].tahun > 0 && tab[i].tanggal != 0 && now.Day() != 0){
			fmt.Printf("PERINGATAN: %s TELAH KADALUARSA PADA TANGGAL %02d/%02d/%04d\n", tab[i].nama, tab[i].tanggal, tab[i].bulan, tab[i].tahun)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada bahan makanan yang mendekati tanggal kadaluarsa atau telah kadaluarsa")
	}
}

func cariBahan(T tabMakanan, n int, nama string) {
	var i int
	var ketemu bool
	ketemu = false
	for i = 0; i < n; i++ {
		if T[i].nama == nama {
			fmt.Println("Data Ditemukan")
			fmt.Println("----------------------------------")
			fmt.Printf("Id bahan makanan: %s\n", T[i].id)
			fmt.Printf("Nama bahan makanan: %s\n", T[i].nama)
			fmt.Printf("Stok Yang Tersedia: %d\n", T[i].stok-T[i].stokTergunakan)
			fmt.Printf("Stok Awal: %d\n", T[i].stok)
			fmt.Printf("Stok Tergunakan: %d\n", T[i].stokTergunakan)
			fmt.Printf("Tanggal Kadaluarsa (dd - mm - yyyy): %02d - %02d - %04d\n", T[i].tanggal, T[i].bulan, T[i].tahun)
			ketemu = true
		}
	}
	if ketemu == false {
		fmt.Println("Data tidak ditemukan")
	}
}

func InsertionSort(T tabMakanan, n int) tabMakanan { // Menaik

	var pass, i int

	var temp BahanMakanan

	for pass = 1; pass < n; pass++ {

		temp = T[pass]
		
		i = pass - 1

		for i >= 0 && (temp.tahun < T[i].tahun || (temp.tahun == T[i].tahun && temp.bulan < T[i].bulan) || (temp.tahun == T[i].tahun && temp.bulan == T[i].bulan && temp.tanggal < T[i].tanggal)) {

			T[i+1] = T[i]

			i--

		}

		T[i+1] = temp

	}

	return T

}

func selectionSort(T tabMakanan, n int) tabMakanan { //Menurun
	var i, idx, pass int
	var temp int

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if T[i].stok > T[idx].stok {
				idx = i
			}
			i++
		}
		temp = T[pass-1].stok
		T[pass-1].stok = T[idx].stok
		T[idx].stok = temp
		pass++
	}
	return T
}

func binarySearchId(T tabMakanan, x string, n int) int {
	var left, mid, right int
	var idx int

	left = 0
	right = n - 1
	idx = -1

	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if x < T[mid].id {
			right = mid - 1
		} else if x > T[mid].id {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}

func sortById(T tabMakanan, n int) tabMakanan {
	var i, j int
	var temp BahanMakanan
	for i = 1; i < n; i++ {
		temp = T[i]
		j = i - 1
		for j >= 0 && T[j].id > temp.id {
			T[j+1] = T[j]
			j--
		}
		T[j+1] = temp
	}
	return T
}

func stokTerbanyak(T tabMakanan, n int) {
	var i, idx int
	idx = 0
	for i = 1; i < n; i++ {
		if T[i].stok-T[i].stokTergunakan > T[idx].stok-T[idx].stokTergunakan && T[i].nama != "" {
			idx = i
		}
	}
	fmt.Println("---------------")
	fmt.Println("Stok Terbanyak")
	fmt.Println("--------------")
	fmt.Printf("Id bahan makanan: %s\n", T[idx].id)
	fmt.Printf("Nama bahan makanan: %s\n", T[idx].nama)
	fmt.Printf("Stok Yang Tersedia: %d\n", T[idx].stok-T[idx].stokTergunakan)
	fmt.Printf("Stok Awal: %d\n", T[idx].stok)
	fmt.Printf("Stok Tergunakan: %d\n", T[idx].stokTergunakan)
	fmt.Printf("Tanggal Kadaluarsa (dd - mm - yyyy): %02d - %02d - %04d\n", T[idx].tanggal, T[idx].bulan, T[idx].tahun)
}

func stokTersedikit(T tabMakanan, n int) {
	var i, idx int
	idx = 0
	for i = 1; i < n; i++ {
		if T[i].stok-T[i].stokTergunakan < T[idx].stok-T[idx].stokTergunakan && T[i].nama != "" {
			idx = i
		}
	}
	fmt.Println("---------------")
	fmt.Println("Stok Tersedikit")
	fmt.Println("---------------")
	fmt.Printf("Id bahan makanan: %s\n", T[idx].id)
	fmt.Printf("Nama bahan makanan: %s\n", T[idx].nama)
	fmt.Printf("Stok Yang Tersedia: %d\n", T[idx].stok-T[idx].stokTergunakan)
	fmt.Printf("Stok Awal: %d\n", T[idx].stok)
	fmt.Printf("Stok Tergunakan: %d\n", T[idx].stokTergunakan)
	fmt.Printf("Tanggal Kadaluarsa (dd - mm - yyyy): %02d - %02d - %04d\n", T[idx].tanggal, T[idx].bulan, T[idx].tahun)
}
