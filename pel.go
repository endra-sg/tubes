package main

import "fmt"

const NMAX int = 1000

type skripsi struct {
	judul      string
	penulis    string
	tahunLulus string
}

type tabSkripsIn [NMAX]skripsi

func main() {
	var dataSkripsi tabSkripsIn
	var jumlahSkripsi int
	var menuDipilih string
	var validasi bool
	var jenis string
	jumlahSkripsi = 0
	validasi = true

	for validasi {
		daftarMenu()
		fmt.Print("Pilih menu: ")
		fmt.Scan(&menuDipilih)

		if menuDipilih == "1" {
			tambahSkripsi(&dataSkripsi, &jumlahSkripsi)
		} else if menuDipilih == "2" {
			if jumlahSkripsi == 0 {
				fmt.Println("\nBelum ada data skripsi. Silakan tambah data terlebih dahulu.")
			} else {
				ubahDataSkripsi(&dataSkripsi, jumlahSkripsi)
			}
		} else if menuDipilih == "3" {
			if jumlahSkripsi == 0 {
				fmt.Println("\nBelum ada data skripsi. Silakan tambah data terlebih dahulu.")
			} else {
				hapusDataSkripsi(&dataSkripsi, &jumlahSkripsi)
			}
		} else if menuDipilih == "4" {
			if jumlahSkripsi == 0 {
				fmt.Println("\nBelum ada data skripsi. Silakan tambah data terlebih dahulu.")
			} else {
				fmt.Println("1 Menggunakan Nama sequentialsearchbyNama")
				fmt.Println("1 Menggunakan Nama binarySearchBynamaS")
				fmt.Scan(&jenis)
				if jenis == "1" {
					sequentialsearchbyNama(&dataSkripsi, jumlahSkripsi)
				} else if jenis == "2" {
					selectionshortbyNama(&dataSkripsi, jumlahSkripsi)
					binarySearchBynama(&dataSkripsi, jumlahSkripsi)
				}
			}
		} else if menuDipilih == "5" {
			if jumlahSkripsi == 0 {
				fmt.Println("\nBelum ada data skripsi. Silakan tambah data terlebih dahulu.")
			} else {
				fmt.Println("Opsi Pengurutan Data")
				fmt.Println("1 Menggunakan Nama")
				fmt.Println("2 Menggunakan Tahun")
				fmt.Scan(&jenis)
				if jenis == "1" {
					selectionshortbyNama(&dataSkripsi, jumlahSkripsi)
				} else {
					selectionshortbyTahun(&dataSkripsi, jumlahSkripsi)
				}
			}
		} else if menuDipilih == "6" {
			if jumlahSkripsi == 0 {
				fmt.Println("\nBelum ada data skripsi. Silakan tambah data terlebih dahulu.")
			} else {
				statistikSkripsi(&dataSkripsi, jumlahSkripsi)
			}
		} else if menuDipilih == "0" {
			fmt.Println("Terima kasih sudah menggunakan aplikasi")
			validasi = false
		} else {
			fmt.Println("Menu yang dipilih tidak valid")
		}
	}
}

func daftarMenu() {
	fmt.Println("\n----- SkripsIn -----")
	fmt.Println("1. Tambah Dokumen Skripsi")
	fmt.Println("2. Ubah Dokumen Skripsi")
	fmt.Println("3. Hapus Dokumen Skripsi")
	fmt.Println("4. Cari Data Skripsi")
	fmt.Println("5. Urutkan Data Skripsi")
	fmt.Println("6. Tampilkan Statistik Skripsi")
	fmt.Println("0. Keluar")
	fmt.Println()
}

func tambahSkripsi(s *tabSkripsIn, maxData *int) {
	var validasi bool
	var addMore string
	validasi = true
	if *maxData >= NMAX {
		fmt.Println("Data sudah penuh")
	}
	for validasi || addMore == "yes" {
		fmt.Print("Judul skripsi: ")
		fmt.Scan(&s[*maxData].judul)
		fmt.Print("Penulis skripsi: ")
		fmt.Scan(&s[*maxData].penulis)
		fmt.Print("Tahun lulus penulis: ")
		fmt.Scan(&s[*maxData].tahunLulus)
		*maxData = *maxData + 1
		fmt.Println("\nTambah data lagi? (yes/no)")
		fmt.Scan(&addMore)
		if addMore == "yes" {
			validasi = true
		} else if addMore == "no" {
			validasi = false
		} else {
			fmt.Println("Pilihan tidak valid")
			fmt.Println("\nTambah data lagi? (yes/no)")
			fmt.Scan(&addMore)
		}
	}
}

func ubahDataSkripsi(s *tabSkripsIn, maxData int) {
	var i, idx int
	var judul, penulis, tahunLulus, confirm string
	var validasi bool = true
	fmt.Println("Silahkan pilih data yang ingin diubah")
	for i = 0; i < maxData; i++ {
		fmt.Printf("%d. %-20s %-15s %s\n", i+1, s[i].judul, s[i].penulis, s[i].tahunLulus)
	}
	for validasi {
		fmt.Print("\nMasukkan nomor data yang ingin diubah: ")
		fmt.Scan(&idx)
		if idx < 1 || idx > maxData {
			fmt.Println("Data tidak tersedia")
			validasi = false
		} else {
			fmt.Print("Judul skripsi: ")
			fmt.Scan(&judul)
			fmt.Print("Penulis skripsi: ")
			fmt.Scan(&penulis)
			fmt.Print("Tahun lulus penulis: ")
			fmt.Scan(&tahunLulus)
			fmt.Println("Data baru:")
			fmt.Printf("Judul skripsi: %s\nPenulis skripsi: %s\nTahun lulus penulis: %s\n", judul, penulis, tahunLulus)
			fmt.Println("Apakah data sudah benar?")
			fmt.Scan(&confirm)
			if confirm == "yes" {
				s[idx-1].judul = judul
				s[idx-1].penulis = penulis
				s[idx-1].tahunLulus = tahunLulus
				fmt.Println("Data berhasil diubah!")
				validasi = false
			} else if confirm == "no" {
				validasi = false
			} else {
				fmt.Println("Pilihan tidak valid!")
			}
		}
	}
}

func hapusDataSkripsi(s *tabSkripsIn, maxData *int) {
	var i, idx int
	var confirm string
	var validasi bool = true
	fmt.Println("Silahkan pilih data yang ingin dihapus")
	for i = 0; i < *maxData; i++ {
		fmt.Printf("%d. %-20s %-15s %s\n", i+1, s[i].judul, s[i].penulis, s[i].tahunLulus)
	}
	for validasi {
		fmt.Print("\nMasukkan nomor data yang ingin dihapus: ")
		fmt.Scan(&idx)
		if idx < 1 || idx > *maxData {
			fmt.Println("Data tidak tersedia")
			validasi = false
		} else {
			for i = idx - 1; i < *maxData; i++ {
				s[i] = s[i+1]
			}
			*maxData = *maxData - 1
			fmt.Println("Ingin menghapus data lagi?(yes/no)")
			fmt.Scan(&confirm)
			if confirm == "yes" {
				validasi = true
			} else if confirm == "no" {
				validasi = false
			} else {
				fmt.Println("Pilihan tidak valid!")
				validasi = false
			}
		}
	}
}

func sequentialsearchbyNama(s *tabSkripsIn, jumlah int) {
	var found bool
	var validation bool
	var key string
	var searchData string
	var j int
	validation = true
	for validation {
		found = false
		fmt.Println("Masukan judul skripsi atau nama penulis skripsi yang ingin di cari")
		fmt.Scan(&key)
		for j = 0; j < jumlah; j++ {
			if s[j].judul == key || s[j].penulis == key {
				fmt.Printf("Data Ditemukan!\nJudul skripsi: %s\nPenulis skripsi: %s\nTahun Lulus penulis: %s\n",
					s[j].judul, s[j].penulis, s[j].tahunLulus)
				found = true
			}
		}
		if !found {
			fmt.Println("Data yang dicari tidak tersedia")
		}
		fmt.Println("Apakah anda ingin mencari data lagi (yes/no)")
		fmt.Scan(&searchData)
		if searchData == "no" {
			validation = false
		} else if searchData != "yes" {
			fmt.Println("Pilihan tidak valid, keluar dari pencarian")
			validation = false
		}
	}
}
func binarySearchBynama(s *tabSkripsIn, n int) {
	var left, right, mid int
	var penulis string
	var found bool

	fmt.Print("Masukkan nama penulis yang dicari: ")
	fmt.Scan(&penulis)

	left = 0
	right = n - 1
	found = false

	for left <= right && !found {
		mid = (left + right) / 2
		if penulis > s[mid].penulis {
			left = mid + 1
		} else if penulis < s[mid].penulis {
			right = mid - 1
		} else {
			fmt.Printf("\n=== DATA DITEMUKAN ===\n")
			fmt.Printf("Judul       : %s\n", s[mid].judul)
			fmt.Printf("Penulis     : %s\n", s[mid].penulis)
			fmt.Printf("Tahun Lulus : %s\n", s[mid].tahunLulus)
			found = true
		}
	}

	if !found {
		fmt.Println("Data yang dicari tidak tersedia")
	}
}
func selectionshortbyNama(s *tabSkripsIn, n int) {
	var pass, idx, i int
	var temp skripsi
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if s[idx].penulis < s[i].penulis {
				idx = i
			}
			i = i + 1
		}
		temp = s[pass-1]
		s[pass-1] = s[idx]
		s[idx] = temp
		pass = pass + 1
	}
}
func selectionshortbyTahun(s *tabSkripsIn, n int) {
	var pass, idx, i int
	var temp skripsi
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if s[idx].tahunLulus < s[i].tahunLulus {
				idx = i
			}
			i = i + 1
		}
		temp = s[pass-1]
		s[pass-1] = s[idx]
		s[idx] = temp
		pass = pass + 1
	}
}
func tampilkanData(s *tabSkripsIn, n int) {
	var i int
	if n == 0 {
		fmt.Println("Belum ada data")
	} else {
		fmt.Println("RESULT DATA SETELAH DI URUTKAN")
		fmt.Printf("%3s %20s %15s %10s\n", "No", "Judul", "Penulis", "Tahun")
		for i = 0; i < n; i++ {
			fmt.Printf("%3d %20s %15s %10s\n", i+1, s[i].judul, s[i].penulis, s[i].tahunLulus)
		}
	}
}
func statistikSkripsi(s *tabSkripsIn, n int) {
	var jumlah, i int
	var tahun string
	var totalSkripsi int
	totalSkripsi = 1
	jumlah = 1
	tahun = s[0].tahunLulus
	if n == 0 {
		fmt.Println("Data Belum Terinput")
	} else {
		fmt.Println("STATISTIK DATA SKRIPSI")
		fmt.Printf("%-10s %-20s\n", "Tahun", "Jumlah Skripsi")
		for i = 1; i < n; i++ {
			totalSkripsi = totalSkripsi + 1
			if s[i].tahunLulus == tahun {
				jumlah = jumlah + 1
			} else {
				fmt.Printf("%-10s %-20d\n", tahun, jumlah)
				tahun = s[i].tahunLulus
				jumlah = 1
			}
		}
		fmt.Printf("%-10s %-20d\n", tahun, jumlah)
		fmt.Println("TOTAL DOKUMEN SKRIPSI", totalSkripsi)
	}
}
