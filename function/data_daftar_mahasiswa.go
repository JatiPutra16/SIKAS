package function

import (
	"fmt"
	"sikas/models"
)

func MenuDaftarMahasiswa(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int) {
	var subPilih int
	var keluar bool

	for !keluar {
		models.Clearscreen()
		models.TampilkanHeader()
		fmt.Println("----- DAFTAR MAHASISWA & STATUS IURAN -----")
		fmt.Println("[1] Nama Ascending  (Selection Sort)")
		fmt.Println("[2] Nama Descending (Insertion Sort)")
		fmt.Println("[3] Tunggakan Ascending  (Selection Sort)")
		fmt.Println("[4] Tunggakan Descending (Insertion Sort)")
		fmt.Println("[5] Kembali ke Menu Utama")
		fmt.Println("============================================")
		fmt.Print("Pilihan: ")
		fmt.Scanln(&subPilih)

		if subPilih == 5 {
			keluar = true
		} else {
			switch subPilih {
			case 1:
				selectionSortNamaAscending(daftarMhs, jumlahMhs)

			case 2:
				insertionSortNamaDescending(daftarMhs, jumlahMhs)

			case 3:
				selectionSortTunggakanAscending(daftarMhs, jumlahMhs)

			case 4:
				insertionSortTunggakanDescending(daftarMhs, jumlahMhs)

			default:
				fmt.Println("\nPilihan tidak valid.")
			}

			fmt.Print("\nTekan Enter untuk melanjutkan...")
			fmt.Scanln()
		}
	}
}

func selectionSortNamaAscending(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int) {
	var sorted [models.NMAX]models.Mahasiswa

	for i := 0; i < *jumlahMhs; i++ {
		sorted[i] = daftarMhs[i]
	}

	for i := 0; i < *jumlahMhs-1; i++ {
		minIdx := i

		for j := i + 1; j < *jumlahMhs; j++ {
			if sorted[j].Nama < sorted[minIdx].Nama {
				minIdx = j
			}
		}

		sorted[i], sorted[minIdx] = sorted[minIdx], sorted[i]
	}

	fmt.Println("\nData Mahasiswa berdasarkan Nama Ascending menggunakan Selection Sort")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("%-15s | %-25s | %-15s | %-15s\n", "NIM", "Nama", "Status", "Tunggakan")
	fmt.Println("--------------------------------------------------------------------------------")

	for i := 0; i < *jumlahMhs; i++ {
		sudahLunas := true
		totalTunggakan := 0

		for j := 0; j < 12; j++ {
			if !sorted[i].Iuran[j].Status {
				sudahLunas = false

				sisa := models.TARGET_IURAN - sorted[i].Iuran[j].TotalTerbayar
				if sisa > 0 {
					totalTunggakan += sisa
				}
			}
		}

		status := "Lunas"
		if !sudahLunas {
			status = "Belum Lunas"
		}

		fmt.Printf("%-15s | %-25s | %-15s | Rp %-12d\n",
			sorted[i].NIM,
			sorted[i].Nama,
			status,
			totalTunggakan,
		)
	}

	fmt.Println("--------------------------------------------------------------------------------")
}

func insertionSortNamaDescending(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int) {
	var sorted [models.NMAX]models.Mahasiswa

	for i := 0; i < *jumlahMhs; i++ {
		sorted[i] = daftarMhs[i]
	}

	for i := 1; i < *jumlahMhs; i++ {
		kunci := sorted[i]
		j := i - 1
		lanjutGeser := true

		for j >= 0 && lanjutGeser {
			if sorted[j].Nama < kunci.Nama {
				sorted[j+1] = sorted[j]
				j--
			} else {
				lanjutGeser = false
			}
		}

		sorted[j+1] = kunci
	}

	fmt.Println("\nData Mahasiswa berdasarkan Nama Descending menggunakan Insertion Sort")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("%-15s | %-25s | %-15s | %-15s\n", "NIM", "Nama", "Status", "Tunggakan")
	fmt.Println("--------------------------------------------------------------------------------")

	for i := 0; i < *jumlahMhs; i++ {
		sudahLunas := true
		totalTunggakan := 0

		for j := 0; j < 12; j++ {
			if !sorted[i].Iuran[j].Status {
				sudahLunas = false

				sisa := models.TARGET_IURAN - sorted[i].Iuran[j].TotalTerbayar
				if sisa > 0 {
					totalTunggakan += sisa
				}
			}
		}

		status := "Lunas"
		if !sudahLunas {
			status = "Belum Lunas"
		}

		fmt.Printf("%-15s | %-25s | %-15s | Rp %-12d\n",
			sorted[i].NIM,
			sorted[i].Nama,
			status,
			totalTunggakan,
		)
	}

	fmt.Println("--------------------------------------------------------------------------------")
}

func selectionSortTunggakanAscending(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int) {
	var sorted [models.NMAX]models.Mahasiswa

	for i := 0; i < *jumlahMhs; i++ {
		sorted[i] = daftarMhs[i]
	}

	for i := 0; i < *jumlahMhs-1; i++ {
		minIdx := i

		for j := i + 1; j < *jumlahMhs; j++ {
			totalMin := 0
			totalJ := 0

			for k := 0; k < 12; k++ {
				if !sorted[minIdx].Iuran[k].Status {
					sisa := models.TARGET_IURAN - sorted[minIdx].Iuran[k].TotalTerbayar
					if sisa > 0 {
						totalMin += sisa
					}
				}
			}

			for k := 0; k < 12; k++ {
				if !sorted[j].Iuran[k].Status {
					sisa := models.TARGET_IURAN - sorted[j].Iuran[k].TotalTerbayar
					if sisa > 0 {
						totalJ += sisa
					}
				}
			}

			if totalJ < totalMin {
				minIdx = j
			}
		}

		sorted[i], sorted[minIdx] = sorted[minIdx], sorted[i]
	}

	fmt.Println("\nData Mahasiswa berdasarkan Tunggakan Ascending menggunakan Selection Sort")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("%-15s | %-25s | %-15s | %-15s\n", "NIM", "Nama", "Status", "Tunggakan")
	fmt.Println("--------------------------------------------------------------------------------")

	for i := 0; i < *jumlahMhs; i++ {
		sudahLunas := true
		totalTunggakan := 0

		for j := 0; j < 12; j++ {
			if !sorted[i].Iuran[j].Status {
				sudahLunas = false

				sisa := models.TARGET_IURAN - sorted[i].Iuran[j].TotalTerbayar
				if sisa > 0 {
					totalTunggakan += sisa
				}
			}
		}

		status := "Lunas"
		if !sudahLunas {
			status = "Belum Lunas"
		}

		fmt.Printf("%-15s | %-25s | %-15s | Rp %-12d\n",
			sorted[i].NIM,
			sorted[i].Nama,
			status,
			totalTunggakan,
		)
	}

	fmt.Println("--------------------------------------------------------------------------------")
}

func insertionSortTunggakanDescending(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int) {
	var sorted [models.NMAX]models.Mahasiswa

	for i := 0; i < *jumlahMhs; i++ {
		sorted[i] = daftarMhs[i]
	}

	for i := 1; i < *jumlahMhs; i++ {
		kunci := sorted[i]
		j := i - 1
		lanjutGeser := true

		for j >= 0 && lanjutGeser {
			totalKunci := 0
			totalJ := 0

			for k := 0; k < 12; k++ {
				if !kunci.Iuran[k].Status {
					sisa := models.TARGET_IURAN - kunci.Iuran[k].TotalTerbayar
					if sisa > 0 {
						totalKunci += sisa
					}
				}
			}

			for k := 0; k < 12; k++ {
				if !sorted[j].Iuran[k].Status {
					sisa := models.TARGET_IURAN - sorted[j].Iuran[k].TotalTerbayar
					if sisa > 0 {
						totalJ += sisa
					}
				}
			}

			if totalJ < totalKunci {
				sorted[j+1] = sorted[j]
				j--
			} else {
				lanjutGeser = false
			}
		}

		sorted[j+1] = kunci
	}

	fmt.Println("\nData Mahasiswa berdasarkan Tunggakan Descending menggunakan Insertion Sort")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("%-15s | %-25s | %-15s | %-15s\n", "NIM", "Nama", "Status", "Tunggakan")
	fmt.Println("--------------------------------------------------------------------------------")

	for i := 0; i < *jumlahMhs; i++ {
		sudahLunas := true
		totalTunggakan := 0

		for j := 0; j < 12; j++ {
			if !sorted[i].Iuran[j].Status {
				sudahLunas = false

				sisa := models.TARGET_IURAN - sorted[i].Iuran[j].TotalTerbayar
				if sisa > 0 {
					totalTunggakan += sisa
				}
			}
		}

		status := "Lunas"
		if !sudahLunas {
			status = "Belum Lunas"
		}

		fmt.Printf("%-15s | %-25s | %-15s | Rp %-12d\n",
			sorted[i].NIM,
			sorted[i].Nama,
			status,
			totalTunggakan,
		)
	}

	fmt.Println("--------------------------------------------------------------------------------")
}
