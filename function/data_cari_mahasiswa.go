package function

import (
	"fmt"
	"sikas/models"
	"strings"
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
			sequentialSearchBelumBayar(daftarMhs, jumlahMhs)

		case 2:
			binarySearchBelumBayar(daftarMhs, jumlahMhs, namaBulan)

		default:
			fmt.Println("\nPilihan tidak valid.")
		}

		fmt.Print("\nTekan Enter untuk melanjutkan...")
		fmt.Scanln()
	}
}

func sequentialSearchBelumBayar(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int) {
	var keyword string
	var ditemukan bool

	fmt.Println("\n--- Sequential Search berdasarkan Nama ---")
	fmt.Print("Masukkan sebagian/seluruh nama: ")
	fmt.Scanln(&keyword)

	fmt.Println("------------------------------------------------------------")
	fmt.Printf("%-15s | %-25s | %-15s\n", "NIM", "Nama", "Tunggakan")
	fmt.Println("------------------------------------------------------------")

	for i := 0; i < *jumlahMhs; i++ {
		namaLower := strings.ToLower(daftarMhs[i].Nama)
		keywordLower := strings.ToLower(keyword)

		if strings.Contains(namaLower, keywordLower) {
			sudahLunas := true
			totalTunggakan := 0

			for j := 0; j < 12; j++ {
				if !daftarMhs[i].Iuran[j].Status {
					sudahLunas = false

					sisa := models.TARGET_IURAN - daftarMhs[i].Iuran[j].TotalTerbayar
					if sisa > 0 {
						totalTunggakan += sisa
					}
				}
			}

			if !sudahLunas {
				fmt.Printf("%-15s | %-25s | Rp %-12d\n",
					daftarMhs[i].NIM,
					daftarMhs[i].Nama,
					totalTunggakan,
				)
				ditemukan = true
			}
		}
	}

	if !ditemukan {
		fmt.Println("Tidak ada mahasiswa yang cocok atau semua sudah lunas.")
	}

	fmt.Println("------------------------------------------------------------")
}

func binarySearchBelumBayar(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int, namaBulan [12]string) {
	var nimCari string
	var sorted [models.NMAX]models.Mahasiswa

	fmt.Println("\n--- Binary Search berdasarkan NIM ---")
	fmt.Print("Masukkan NIM yang dicari: ")
	fmt.Scanln(&nimCari)

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

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		if sorted[tengah].NIM == nimCari {
			hasil = tengah
			break
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
