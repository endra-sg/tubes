package main

import (
	"fmt"
	"strings"
)

const NMAX int = 100

// data skripsi
type skripsi struct {
	judul           string
	penulis         string
	tahunLulus      int
	pembimbing      string
	topik           string
	statusKelulusan string
}

// menyimpan array dari skripsi, dengan maksimal data yang bisa disimpan sebanyak NMAX
type tabSkripsIn [NMAX]skripsi

// untuk memilih menu dan apa yang akan dilakukan
func main() {
	var dataSkripsi tabSkripsIn
	var jumlahSkripsi int
	var menuDipilih int
	var validasi bool
	var jenis int
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
				fmt.Println("\n1. Mencari Judul Skripsi (Sequential Search)")
				fmt.Println("2. Mencari Nama Penulis (Binary Search)")
				fmt.Print("\nPilih jenis pencarian: ")
				fmt.Scan(&jenis)
				if jenis == 1 {
					sequentialSearch(&dataSkripsi, jumlahSkripsi)
				} else if jenis == 2 {
					selectionSort(&dataSkripsi, jumlahSkripsi)
					binarySearch(&dataSkripsi, jumlahSkripsi)
				} else {
					fmt.Println("Menu yang dipilih tidak valid")
				}
			}
		case 5:
			if jumlahSkripsi == 0 {
				fmt.Println("\nBelum ada data skripsi. Silakan tambah data terlebih dahulu.")
			} else {
				fmt.Println("1. Pengurutan Berdasarkan Nama")
				fmt.Println("2. Pengurutan Berdasarkan Tahun")
				fmt.Print("\nPilih jenis pengurutan: ")
				fmt.Scan(&jenis)
				if jenis == 1 {
					selectionSort(&dataSkripsi, jumlahSkripsi)
					fmt.Println("Data berhasil diurutkan! Silahkan pilih menu \"Tampilkan Data\" untuk melihat data terurut")
				} else if jenis == 2 {
					insertionSort(&dataSkripsi, jumlahSkripsi)
					fmt.Println("Data berhasil diurutkan! Silahkan pilih menu \"Tampilkan Data\" untuk melihat data terurut")
				} else {
					fmt.Println("Menu yang dipilih tidak valid")
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

// procedure untuk menampilkan list menu utama yang bisa dipilih
func daftarMenu() {
	fmt.Println("\n=======   Aplikasi SkripsIn   =======")
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

// function untuk menginputkan data skripsi dari pengguna
// mengembalikan struct skripsi dari yang diinputkan oleh pengguna
func inputData() skripsi {
	var s skripsi
	var judul, penulis, pembimbing, topik, statusKelulusan string
	fmt.Print("Judul skripsi: ")
	fmt.Scan(&judul)
	fmt.Print("Penulis skripsi: ")
	fmt.Scan(&penulis)
	fmt.Print("Tahun lulus penulis: ")
	fmt.Scan(&s.tahunLulus)
	fmt.Print("Pembimbing: ")
	fmt.Scan(&pembimbing)
	fmt.Print("Topik Penelitian: ")
	fmt.Scan(&topik)
	fmt.Print("Status Kelulusan: ")
	fmt.Scan(&statusKelulusan)
	s.judul = strings.ToLower(judul)
	s.penulis = strings.ToLower(penulis)
	s.pembimbing = strings.ToLower(pembimbing)
	s.topik = strings.ToLower(topik)
	s.statusKelulusan = strings.ToLower(statusKelulusan)
	return s
}

// function untuk memvalidasi data skripsi yang sudah diinputkan pengguna
// parameter yang digunakan adalah struct skripsi
// mengembalikan true jika data sesuai, dan false jika tidak sesuai
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

// procedure untuk menambah data skripsi yang diinputkan pengguna dan tervalidasi
// parameter yang digunakan adalah s untuk menyimpan array dan mengupdate array dengan data yang dimasukkan
// maxData untuk menghitung jumlah data yang berhasil ditambah dan mengupdate jumlah maxdData
func tambahSkripsi(s *tabSkripsIn, maxData *int) {
	var check skripsi
	if *maxData >= NMAX {
		fmt.Println("Data sudah penuh")
	}
	check = inputData()
	if validateInput(check) {
		s[*maxData] = check
		*maxData = *maxData + 1
		fmt.Println("Data berhasil ditambahkan")
	} else {
		fmt.Println("Data yang dimasukkan tidak valid!")
	}
}

// procedure untuk mengubah data skripsi yang diinputkan pengguna dan tervalidasi
// parameter yang digunakan adalah s untuk menyimpan array dan mengupdate array dengan data yang diubah
// maxData untuk menghitung jumlah data yang tersedia
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
				fmt.Println("Apakah data sudah benar? (yes/no)")
				fmt.Scan(&confirm)
				if confirm == "yes" {
					s[idx-1] = check
					fmt.Println("Data berhasil diubah!")
					validasi = false
				} else if confirm == "no" {
					fmt.Println("Data gagal diubah!")
					validasi = false
				} else {
					fmt.Println("Pilihan tidak valid!")
				}
			}
		}
	}
}

// procedure untuk menghapus data skripsi yang diinginkan
// parameter yang digunakan adalah s untuk menyimpan array dan mengupdate array yang berubah setelah dihapus
// maxData untuk menghitung jumlah data yang tersedia dan mengupdate jumlah data baru setelah dihapus
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

// procedure untuk mencari judul skripsi
// parameter yang digunakan s sebagai array skripsi, dan jumlah sebagai jumlah skripsi
func sequentialSearch(s *tabSkripsIn, jumlah int) {
	var found bool
	var validation bool
	var key string
	var j int

	validation = true
	for validation {
		found = false
		fmt.Print("\nJudul skripsi yang ingin di cari: ")
		fmt.Scan(&key)
		key = strings.ToLower(key)
		for j = 0; j < jumlah; j++ {
			if s[j].judul == key {
				fmt.Printf("\n=== DATA DITEMUKAN ===\n")
				fmt.Printf("\nJudul            : %s\n", s[j].judul)
				fmt.Printf("Penulis          : %s\n", s[j].penulis)
				fmt.Printf("Tahun lulus      : %d\n", s[j].tahunLulus)
				fmt.Printf("Pembimbing       : %s\n", s[j].pembimbing)
				fmt.Printf("Topik penelitian : %s\n", s[j].topik)
				fmt.Printf("Status kelulusan : %s\n", s[j].statusKelulusan)
				found = true
				validation = false
			}
		}
		if !found {
			fmt.Println("Data yang dicari tidak tersedia")
			validation = false
		}
	}
}

// procedure untuk mencari judul nama penulis
// parameter yang digunakan s sebagai array skripsi, dan n sebagai jumlah skripsi
func binarySearch(s *tabSkripsIn, n int) {
	var left, right, mid int
	var cariPenulis string
	var found, validation bool
	for validation {
		fmt.Print("\nNama penulis yang dicari: ")
		fmt.Scan(&cariPenulis)
		cariPenulis = strings.ToLower((cariPenulis))
		left = 0
		right = n - 1
		found = false

		for left <= right && !found {
			mid = (left + right) / 2
			if cariPenulis > s[mid].penulis {
				left = mid + 1
			} else if cariPenulis < s[mid].penulis {
				right = mid - 1
			} else {
				fmt.Printf("\n=== DATA DITEMUKAN ===\n")
				fmt.Printf("\nJudul            : %s\n", s[mid].judul)
				fmt.Printf("Penulis          : %s\n", s[mid].penulis)
				fmt.Printf("Tahun lulus      : %d\n", s[mid].tahunLulus)
				fmt.Printf("Pembimbing       : %s\n", s[mid].pembimbing)
				fmt.Printf("Topik penelitian : %s\n", s[mid].topik)
				fmt.Printf("Status kelulusan : %s\n", s[mid].statusKelulusan)
				found = true
			}
		}

		if !found {
			fmt.Println("Data yang dicari tidak tersedia")
			validation = false
		}
	}
}

// procedure untuk mengurutkan nama penulis dengan selection sort secara asccending
// parameter s sebagai array skripsi, dan n sebagai jumlah array yang terisi
func selectionSort(s *tabSkripsIn, n int) {
	var pass, idx, i int
	var temp skripsi
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if s[idx].penulis > s[i].penulis {
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

// procedure untuk mengurutkan tahun dengan selection sort secara ascending
// parameter s sebagai array skripsi, dan n sebagai jumlah array yang terisi
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

// procedure yang digunakan untuk menampilakn data
// parameter s untuk array skripsi dan n untuk jumlah skripsi
func tampilkanData(s tabSkripsIn, n int) {
	var i int
	fmt.Println("+-----+----------------------+----------------------+-------------+----------------------+-------------------------+------------------+")
	fmt.Printf("| %-3s | %-20s | %-20s | %-10s | %-20s | %-23s | %-12s |\n", "NO", "Judul", "Penulis", "Tahun lulus", "Pembimbing", "Topik penelitian", "Status kelulusan")
	fmt.Println("+-----+----------------------+----------------------+-------------+----------------------+-------------------------+------------------+")
	for i = 0; i < n; i++ {
		fmt.Printf("| %-3d | %-20s | %-20s | %-11d | %-20s | %-23s | %-16s |\n", i+1, s[i].judul, s[i].penulis, s[i].tahunLulus, s[i].pembimbing, s[i].topik, s[i].statusKelulusan)
		fmt.Println("+-----+----------------------+----------------------+-------------+----------------------+-------------------------+------------------+")
	}
}

// procedure untuk menampilkan statistik yang ingin dilihat sesuai dengan data skripsi
// parameter s sebagai array skripsi dan n sebagai jumlah skripsi
func statistikSkripsi(s *tabSkripsIn, n int) {
	var i, j, menu int
	var ditemukan bool
	var jumlah int
	var lulus, tidakLulus int

	fmt.Println("============================")
	fmt.Println("   STATISTIK DATA SKRIPSI   ")
	fmt.Println("============================")
	fmt.Println("Pilih menu statistik yang ingin dipilih")
	menuStatistik()
	fmt.Print("Menu dipilih: ")
	fmt.Scan(&menu)

	switch menu {
	// persentase topik yg dibuat (baik lulus atau tidak) seluruh tahun
	case 1:
		fmt.Println("==============================")
		fmt.Println("   Persentase Topik Skripsi ")
		fmt.Println("==============================")
		fmt.Printf("\nTotal Data Skripsi : %d\n", n)
		fmt.Println("----------------------------------")
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

				fmt.Printf("%-24s : %.2f%%\n", s[i].topik, (float64(jumlah)/float64(n))*100)
			}
		}
		fmt.Println("----------------------------------")
	case 2:
		fmt.Println("=================================")
		fmt.Println("   Tingkat Kelulusan Mahasiswa")
		fmt.Println("=================================")
		fmt.Printf("\nTotal Data Skripsi : %d\n", n)
		fmt.Println("----------------------------------")

		lulus = 0
		tidakLulus = 0

		for i = 0; i < n; i++ {
			if s[i].statusKelulusan == "lulus" {
				lulus = lulus + 1
			} else {
				tidakLulus = tidakLulus + 1
			}
		}
		fmt.Printf("%-12s : %.2f%%\n", "Lulus", (float64(lulus)/float64(n))*100)
		fmt.Printf("%-12s : %.2f%%\n", "Tidak Lulus", (float64(tidakLulus)/float64(n))*100)
		fmt.Println("----------------------------------")
	case 3:
		//pengecekan apakah tahun double input agar menghindari output double
		fmt.Println("=================================")
		fmt.Println("   Jumlah Skripsi Setiap Tahun ")
		fmt.Println("=================================")
		fmt.Printf("\nTotal Data Skripsi : %d\n", n)
		fmt.Println("----------------------------------")

		for i = 0; i < n; i++ {
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

				fmt.Printf("%-8d : %d Skripsi\n", s[i].tahunLulus, jumlah)
			}
		}
		fmt.Println("----------------------------------")
	case 4:
		fmt.Println("==================================")
		fmt.Println("   Statistik Pembimbing Skripsi ")
		fmt.Println("==================================")
		fmt.Printf("\nTotal Data Skripsi : %d\n", n)
		fmt.Println("----------------------------------")

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

				fmt.Printf("%-20s : %d Skripsi\n", s[i].pembimbing, jumlah)
			}
		}
		fmt.Println("----------------------------------")

	case 5:
		fmt.Println("====================================================")
		fmt.Println("   Persentase Kelulusan Berdasarkan Topik Skripsi ")
		fmt.Println("====================================================")
		fmt.Printf("\nTotal Data Skripsi : %d\n", n)
		fmt.Println("----------------------------------")

		//cari topik nya
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
						if s[j].statusKelulusan == "lulus" {
							jumlah = jumlah + 1
						}
					}
				}
				fmt.Printf("%-24s : %.2f%%\n", s[i].topik, (float64(jumlah)/float64(n))*100)
			}
		}
		fmt.Println("----------------------------------")
	default:
		fmt.Println("Menu yang dipilih tidak valid")
	}
}

// procedure untuk menampilkan menu untuk procedure statistik
func menuStatistik() {
	fmt.Println("\n1. Persentase Topik Skripsi")
	fmt.Println("2. Tingkat Kelulusan Mahasiswa")
	fmt.Println("3. Jumlah Skripsi Setiap Tahun")
	fmt.Println("4. Statistik Pembimbing Skripsi")
	fmt.Println("5. Persentase Kelulusan Berdasarkan Topik Skripsi")
	fmt.Println()
}
