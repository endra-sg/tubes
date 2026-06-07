package main

import (
	"fmt"
)

const NMAX int = 1000

// data skripsi
type skripsi struct {
	judul           string
	penulis         string
	tahunLulus      int
	pembimbing      string
	topik           string
	statusKelulusan string
}

// menyimpan array dari skripsi
type tabSkripsIn [NMAX]skripsi

// untuk memilih menu dan apa yang akan dilakukan
func main() {
	var dataSkripsi tabSkripsIn
	var jumlahSkripsi int
	var menuDipilih int
	var validasi bool
	var jenis string
	jumlahSkripsi = 0
	validasi = true

	for validasi {
		daftarMenu()
		fmt.Print("Pilih menu: ")
		fmt.Scan(&menuDipilih)

		switch menuDipilih {
		case 1:
			tambahSkripsi(&dataSkripsi, &jumlahSkripsi)
		case 2:
			if jumlahSkripsi == 0 {
				fmt.Println("\nBelum ada data skripsi. Silakan tambah data terlebih dahulu.")
			} else {
				ubahDataSkripsi(&dataSkripsi, jumlahSkripsi)
			}
		case 3:
			if jumlahSkripsi == 0 {
				fmt.Println("\nBelum ada data skripsi. Silakan tambah data terlebih dahulu.")
			} else {
				hapusDataSkripsi(&dataSkripsi, &jumlahSkripsi)
			}
		case 4:
			if jumlahSkripsi == 0 {
				fmt.Println("\nBelum ada data skripsi. Silakan tambah data terlebih dahulu.")
			} else {
				fmt.Println("1. Menggunakan nama penulis atau judul skripsi (Sequential Search)")
				fmt.Println("2. Menggunakan nama penulis (Binary Search)")
				fmt.Print("Pilih jenis searching: ")
				fmt.Scan(&jenis)
				if jenis == "1" {
					sequentialSearch(&dataSkripsi, jumlahSkripsi)
				} else if jenis == "2" {
					selectionSort(&dataSkripsi, jumlahSkripsi)
					binarySearch(&dataSkripsi, jumlahSkripsi)
				}
			}
		case 5:
			if jumlahSkripsi == 0 {
				fmt.Println("\nBelum ada data skripsi. Silakan tambah data terlebih dahulu.")
			} else {
				fmt.Println("Opsi Pengurutan Data")
				fmt.Println("1. Menggunakan Nama")
				fmt.Println("2. Menggunakan Tahun")
				fmt.Print("Pilih jenis sorting: ")
				fmt.Scan(&jenis)
				if jenis == "1" {
					selectionSort(&dataSkripsi, jumlahSkripsi)
				} else {
					insertionSort(&dataSkripsi, jumlahSkripsi)
				}
			}
		case 6:
			if jumlahSkripsi == 0 {
				fmt.Println("\nBelum ada data skripsi. Silakan tambah data terlebih dahulu.")
			} else {
				statistikSkripsi(&dataSkripsi, jumlahSkripsi)
			}
		case 7:
			if jumlahSkripsi == 0 {
				fmt.Println("\nBelum ada data skripsi. Silakan tambah data terlebih dahulu.")
			} else {
				tampilkanData(dataSkripsi, jumlahSkripsi)
			}
		case 0:
			fmt.Println("Terima kasih sudah menggunakan aplikasi")
			validasi = false
		default:
			fmt.Println("Menu yang dipilih tidak valid")
		}
	}
}

// hanya menampilkan list menu yang bisa dipilih
func daftarMenu() {
	fmt.Println("\n----- SkripsIn -----")
	fmt.Println("1. Tambah Dokumen Skripsi")
	fmt.Println("2. Ubah Dokumen Skripsi")
	fmt.Println("3. Hapus Dokumen Skripsi")
	fmt.Println("4. Cari Data Skripsi")
	fmt.Println("5. Urutkan Data Skripsi")
	fmt.Println("6. Tampilkan Statistik Skripsi")
	fmt.Println("7. Tampilkan data")
	fmt.Println("0. Keluar")
	fmt.Println()
}

/*
	untuk menginput data yang akan dimasukkan ke array

dan mengembalikan dalam bentuk struct
*/
func inputData() skripsi {
	var s skripsi
	fmt.Print("Judul skripsi: ")
	fmt.Scan(&s.judul)
	fmt.Print("Penulis skripsi: ")
	fmt.Scan(&s.penulis)
	fmt.Print("Tahun lulus penulis: ")
	fmt.Scan(&s.tahunLulus)
	fmt.Print("Pembimbing: ")
	fmt.Scan(&s.pembimbing)
	fmt.Print("Topik Penelitian: ")
	fmt.Scan(&s.topik)
	fmt.Print("Status Kelulusan: ")
	fmt.Scan(&s.statusKelulusan)

	return s
}

/*
memvalidasi inputan agar sesuai dengan tipe data yang diminta
dan tidak menerima inputan kosong, tahun tidak negatif dan nol, serta
topik yang tersedia
*/
func validateInput(s skripsi) bool {
	if s.judul == "" || s.penulis == "" || s.pembimbing == "" {
		return false
	} else if s.tahunLulus < 0 {
		return false
	} else if s.topik != "data_science" && s.topik != "artificial_intelligence" && s.topik != "cybersecurity" && s.topik != "software_engineering" &&
		s.topik != "iot" && s.topik != "uiux" {
		return false
	} else if s.statusKelulusan != "lulus" && s.statusKelulusan != "tidak_lulus" {
		return false
	}
	return true
}

// menambah data skripsi yang sudah diinputkan dan divalidasi ke dalam array
func tambahSkripsi(s *tabSkripsIn, maxData *int) {
	var check skripsi
	if *maxData >= NMAX {
		fmt.Println("Data sudah penuh")
	}

	check = inputData()

	if validateInput(check) {
		s[*maxData] = check
		*maxData = *maxData + 1
		fmt.Println("Data berhaasil ditambahkan")
	}
}

// mengubah salah satu data yang sudah diinputkan
func ubahDataSkripsi(s *tabSkripsIn, maxData int) {
	var idx int
	var confirm string
	var validasi bool = true
	var check skripsi
	var data tabSkripsIn

	data = *s

	fmt.Println("Silahkan pilih data yang ingin diubah")
	tampilkanData(data, maxData)

	for validasi {
		fmt.Print("\nMasukkan nomor data yang ingin diubah: ")
		fmt.Scan(&idx)
		if idx < 1 || idx > maxData {
			fmt.Println("Data tidak tersedia")
			validasi = false
		} else {
			check = inputData()
			if validateInput(check) {
				fmt.Println("Data baru:")
				fmt.Printf("Judul skripsi: %s\nPenulis skripsi: %s\nTahun Lulus penulis: %d\nPembimbing: %s\nTopik penelitian: %s\nStatus kelulusan: %s\n",
					check.judul, check.penulis, check.tahunLulus, check.pembimbing, check.topik, check.statusKelulusan)
				fmt.Println("Apakah data sudah benar?")
				fmt.Scan(&confirm)
				if confirm == "yes" {
					s[idx-1] = check
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
}

func hapusDataSkripsi(s *tabSkripsIn, maxData *int) {
	var i, idx int
	var validasi bool = true
	var n int
	var data tabSkripsIn
	data = *s
	n = *maxData
	fmt.Println("Silahkan pilih data yang ingin dihapus")
	tampilkanData(data, n)

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
			fmt.Println("Data berhasil dihapus")
			validasi = false
		}
	}
}

func sequentialSearch(s *tabSkripsIn, jumlah int) {
	var found bool
	var validation bool
	var key string
	var j int

	validation = true
	for validation {
		found = false
		fmt.Println("Masukan judul skripsi yang ingin di cari")
		fmt.Scan(&key)
		for j = 0; j < jumlah; j++ {
			if s[j].judul == key {
				fmt.Printf("Data Ditemukan!\nJudul skripsi: %s\nPenulis skripsi: %s\nTahun Lulus penulis: %s\nPembimbing: %s\nTopik penelitian: %s\nStatus kelulusan: %s\n",
					s[j].judul, s[j].penulis, s[j].tahunLulus, s[j].pembimbing, s[j].topik, s[j].statusKelulusan)
				validation = false
			}
		}
		if !found {
			fmt.Println("Data yang dicari tidak tersedia")
		}
	}
}
func binarySearch(s *tabSkripsIn, n int) {
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
			fmt.Printf("Judul            : %s\n", s[mid].judul)
			fmt.Printf("Penulis          : %s\n", s[mid].penulis)
			fmt.Printf("Tahun Lulus      : %s\n", s[mid].tahunLulus)
			fmt.Printf("Pembimbing       : %s\n", s[mid].pembimbing)
			fmt.Printf("Topik Penelitian : %s\n", s[mid].topik)
			fmt.Printf("Status Kelulusan : %s\n", s[mid].statusKelulusan)
		}
	}

	if !found {
		fmt.Println("Data yang dicari tidak tersedia")
	}
}
func selectionSort(s *tabSkripsIn, n int) {
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
func insertionSort(s *tabSkripsIn, n int) {
	var pass, i int
	var temp skripsi
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = s[pass]
		for i > 0 && temp.tahunLulus < s[i-1].tahunLulus {
			s[i] = s[i-1]
			i = i - 1
		}
		s[i] = temp
		pass = pass + 1
	}
}

func tampilkanData(s tabSkripsIn, n int) {
	var i int
	fmt.Printf("| %-1s | %-20s | %-15s | %-10s | %-20s | %-23s | %-12s |\n", "No", "Judul", "Penulis", "Tahun", "Pembimbing", "Topik Penelitian", "Status Kelulusan")
	fmt.Println("+----+----------------------+-----------------+------------+----------------------+-------------------------+------------------+")
	for i = 0; i < n; i++ {
		fmt.Printf("| %-1d. | %-20s | %-15s | %-10d | %-20s | %-23s | %-16s |\n", i+1, s[i].judul, s[i].penulis, s[i].tahunLulus, s[i].pembimbing, s[i].topik, s[i].statusKelulusan)
		fmt.Println("+----+----------------------+-----------------+------------+----------------------+-------------------------+------------------+")
	}
}

func statistikSkripsi(s *tabSkripsIn, n int) {
	var i, j int
	var ditemukan bool
	var jumlah int
	var lulus, belumLulus int

	fmt.Println("========================")
	fmt.Println("STATISTIK DATA SKRIPSI")
	fmt.Println("========================")

	fmt.Printf("\nTotal Skripsi           : %d\n", n)
	fmt.Println("\nBerdasarkan Tahun Lulus")

	for i = 0; i < n; i++ {\\ pengecekan apakah tahun double input agar menghindari output double
		ditemukan = false

		for j = 0; j < i; j++ {
			if s[i].tahunLulus == s[j].tahunLulus {
				ditemukan = true
			}
		}

		if !ditemukan {
			jumlah = 0

			for j = 0; j < n; j++ {
				if s[j].tahunLulus == s[i].tahunLulus {
					jumlah = jumlah + 1
				}
			}

			fmt.Printf("%-24d : %d Skripsi\n", s[i].tahunLulus, jumlah)
		}
	}
	fmt.Println("\nBerdasarkan Topik")

	for i = 0; i < n; i++ { 
		ditemukan = false

		for j = 0; j < i; j++ {
			if s[i].topik == s[j].topik {
				ditemukan = true
			}
		}

		if !ditemukan {
			jumlah = 0

			for j = 0; j < n; j++ {
				if s[j].topik == s[i].topik {
					jumlah = jumlah + 1
				}
			}

			fmt.Printf("%-24s : %d Skripsi\n", s[i].topik, jumlah)
		}
	}

	fmt.Println("\nBerdasarkan Pembimbing")

	for i = 0; i < n; i++ {
		ditemukan = false

		for j = 0; j < i; j++ {
			if s[i].pembimbing == s[j].pembimbing {
				ditemukan = true
			}
		}

		if !ditemukan {
			jumlah = 0

			for j = 0; j < n; j++ {
				if s[j].pembimbing == s[i].pembimbing {
					jumlah = jumlah + 1
				}
			}

			fmt.Printf("%-24s : %d Skripsi\n", s[i].pembimbing, jumlah)
		}
	}

	lulus = 0
	belumLulus = 0

	for i = 0; i < n; i++ {
		if s[i].statusKelulusan == "Lulus" {
			lulus = lulus + 1
		} else {
			belumLulus = belumLulus + 1
		}
	}

	fmt.Println("\nBerdasarkan Status")
	fmt.Printf("%-24s : %d Skripsi\n", "Lulus", lulus)
	fmt.Printf("%-24s : %d Skripsi\n", "Belum Lulus", belumLulus)
}
