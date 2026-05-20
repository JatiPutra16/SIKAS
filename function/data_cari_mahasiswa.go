package function

import (
	"fmt"
	"sikas/models"
)

func MenuCariMahasiswaBelumBayar(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int) {
	var subPilih int
	var namaBulan [12]string

	namaBulan = [12]string{
		"Januari", "Februari", "Maret", "April",
		"Mei", "Juni", "Juli", "Agustus",
		"September", "Oktober", "November", "Desember",
	}

	for {
		models.Clearscreen()
		models.TampilkanHeader()
		fmt.Println("----- CARI MAHASISWA BELUM BAYAR -----")
		fmt.Println("[1] Sequential Search berdasarkan Nama")
		fmt.Println("[2] Binary Search berdasarkan NIM")
		fmt.Println("[3] Kembali ke Menu Utama")
		fmt.Println("=======================================")
		fmt.Print("Pilihan: ")
		fmt.Scanln(&subPilih)

		if subPilih == 3 {
			break
		}

		switch subPilih {
		case 1:
			sequentialSearchBelumBayar(daftarMhs, jumlahMhs, namaBulan)

		case 2:
			binarySearchBelumBayar(daftarMhs, jumlahMhs, namaBulan)

		default:
			fmt.Println("\nPilihan tidak valid.")
		}

		fmt.Print("\nTekan Enter untuk melanjutkan...")
		fmt.Scanln()
	}
}

func sequentialSearchBelumBayar(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int, namaBulan [12]string) {
	var keyword string
	var idx int

	idx = -1

	fmt.Println("\n--- Sequential Search berdasarkan Nama ---")

	for keyword == "" {
		fmt.Print("Masukkan nama mahasiswa: ")
		fmt.Scanln(&keyword)

		if keyword == "" {
			fmt.Println("Nama tidak boleh kosong.")
		}
	}

	for i := 0; i < *jumlahMhs; i++ {
		if daftarMhs[i].Nama == keyword {
			idx = i
		}
	}

	fmt.Println("------------------------------------------------------------")

	if idx == -1 {
		fmt.Printf("Mahasiswa dengan nama %s tidak ditemukan.\n", keyword)
	} else {
		fmt.Println("NIM          :", daftarMhs[idx].NIM)
		fmt.Println("Nama Lengkap :", daftarMhs[idx].Nama)

		sudahLunas := true
		totalTunggakan := 0

		for i := 0; i < 12; i++ {
			if !daftarMhs[idx].Iuran[i].Status {
				sudahLunas = false

				sisa := models.TARGET_IURAN - daftarMhs[idx].Iuran[i].TotalTerbayar
				if sisa > 0 {
					totalTunggakan += sisa
				}
			}
		}

		if sudahLunas {
			fmt.Println("Status       : LUNAS")
		} else {
			fmt.Println("Status       : BELUM LUNAS")
			fmt.Printf("Total Tunggakan: Rp %d\n", totalTunggakan)
			fmt.Println("------------------------------------------------------------")
			fmt.Println("Detail bulan belum lunas:")

			for i := 0; i < 12; i++ {
				if !daftarMhs[idx].Iuran[i].Status {
					sisa := models.TARGET_IURAN - daftarMhs[idx].Iuran[i].TotalTerbayar

					if sisa < 0 {
						sisa = 0
					}

					fmt.Printf("- %-10s | Terbayar: Rp %-10d | Sisa: Rp %d\n",
						namaBulan[i],
						daftarMhs[idx].Iuran[i].TotalTerbayar,
						sisa,
					)
				}
			}
		}
	}

	fmt.Println("------------------------------------------------------------")
}

func binarySearchBelumBayar(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int, namaBulan [12]string) {
	var nimCari string
	var sorted [models.NMAX]models.Mahasiswa

	fmt.Println("\n--- Binary Search berdasarkan NIM ---")

	for nimCari == "" {
		fmt.Print("Masukkan NIM yang dicari: ")
		fmt.Scanln(&nimCari)

		if nimCari == "" {
			fmt.Println("NIM tidak boleh kosong.")
		}
	}

	for i := 0; i < *jumlahMhs; i++ {
		sorted[i] = daftarMhs[i]
	}

	for i := 0; i < *jumlahMhs-1; i++ {
		minIdx := i

		for j := i + 1; j < *jumlahMhs; j++ {
			if sorted[j].NIM < sorted[minIdx].NIM {
				minIdx = j
			}
		}

		sorted[i], sorted[minIdx] = sorted[minIdx], sorted[i]
	}

	kiri := 0
	kanan := *jumlahMhs - 1
	hasil := -1

	for kiri <= kanan && hasil == -1 {
		tengah := (kiri + kanan) / 2

		if sorted[tengah].NIM == nimCari {
			hasil = tengah
		} else if sorted[tengah].NIM < nimCari {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	fmt.Println("------------------------------------------------------------")

	if hasil == -1 {
		fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan.\n", nimCari)
	} else {
		fmt.Println("NIM          :", sorted[hasil].NIM)
		fmt.Println("Nama Lengkap :", sorted[hasil].Nama)

		sudahLunas := true
		totalTunggakan := 0

		for i := 0; i < 12; i++ {
			if !sorted[hasil].Iuran[i].Status {
				sudahLunas = false

				sisa := models.TARGET_IURAN - sorted[hasil].Iuran[i].TotalTerbayar
				if sisa > 0 {
					totalTunggakan += sisa
				}
			}
		}

		if sudahLunas {
			fmt.Println("Status       : LUNAS")
		} else {
			fmt.Println("Status       : BELUM LUNAS")
			fmt.Printf("Total Tunggakan: Rp %d\n", totalTunggakan)
			fmt.Println("------------------------------------------------------------")
			fmt.Println("Detail bulan belum lunas:")

			for i := 0; i < 12; i++ {
				if !sorted[hasil].Iuran[i].Status {
					sisa := models.TARGET_IURAN - sorted[hasil].Iuran[i].TotalTerbayar

					if sisa < 0 {
						sisa = 0
					}

					fmt.Printf("- %-10s | Terbayar: Rp %-10d | Sisa: Rp %d\n",
						namaBulan[i],
						sorted[hasil].Iuran[i].TotalTerbayar,
						sisa,
					)
				}
			}
		}
	}

	fmt.Println("------------------------------------------------------------")
}
