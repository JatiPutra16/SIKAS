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
		fmt.Println("[1] Nama Ascending  (Insertion Sort)")
		fmt.Println("[2] Nama Descending (Selection Sort)")
		fmt.Println("[3] Tunggakan Ascending  (Insertion Sort)")
		fmt.Println("[4] Tunggakan Descending (Selection Sort)")
		fmt.Println("[5] Kembali ke Menu Utama")
		fmt.Println("============================================")
		fmt.Print("Pilihan: ")
		fmt.Scanln(&subPilih)

		if subPilih == 5 {
			keluar = true
		} else {
			switch subPilih {
			case 1:
				insertionSortNamaAscending(daftarMhs, jumlahMhs)

			case 2:
				selectionSortNamaDescending(daftarMhs, jumlahMhs)

			case 3:
				insertionSortTunggakanAscending(daftarMhs, jumlahMhs)

			case 4:
				selectionSortTunggakanDescending(daftarMhs, jumlahMhs)

			default:
				fmt.Println("\nPilihan tidak valid.")
			}

			fmt.Print("\nTekan Enter untuk melanjutkan...")
			fmt.Scanln()
		}
	}
}

func insertionSortNamaAscending(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int) {
	var sorted [models.NMAX]models.Mahasiswa
	var pass, i, j int
	var temp models.Mahasiswa
	var sudahLunas bool
	var totalTunggakan int
	var sisa int
	var status string

	for i = 0; i < *jumlahMhs; i++ {
		sorted[i] = daftarMhs[i]
	}

	pass = 1
	for pass <= *jumlahMhs-1 {
		i = pass
		temp = sorted[pass]

		for i > 0 && temp.Nama < sorted[i-1].Nama {
			sorted[i] = sorted[i-1]
			i--
		}

		sorted[i] = temp
		pass++
	}

	fmt.Println("\nData Mahasiswa berdasarkan Nama Ascending menggunakan Insertion Sort")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("%-15s | %-25s | %-15s | %-15s\n", "NIM", "Nama", "Status", "Tunggakan")
	fmt.Println("--------------------------------------------------------------------------------")

	for i = 0; i < *jumlahMhs; i++ {
		sudahLunas = true
		totalTunggakan = 0

		for j = 0; j < 12; j++ {
			if !sorted[i].Iuran[j].Status {
				sudahLunas = false

				sisa = models.TARGET_IURAN - sorted[i].Iuran[j].TotalTerbayar
				if sisa > 0 {
					totalTunggakan += sisa
				}
			}
		}

		status = "Lunas"
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

func selectionSortNamaDescending(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int) {
	var sorted [models.NMAX]models.Mahasiswa
	var pass, idx, i, j int
	var temp models.Mahasiswa
	var sudahLunas bool
	var totalTunggakan int
	var sisa int
	var status string

	for i = 0; i < *jumlahMhs; i++ {
		sorted[i] = daftarMhs[i]
	}

	pass = 1
	for pass <= *jumlahMhs-1 {
		idx = pass - 1
		i = pass

		for i < *jumlahMhs {
			if sorted[idx].Nama < sorted[i].Nama {
				idx = i
			}
			i++
		}

		temp = sorted[pass-1]
		sorted[pass-1] = sorted[idx]
		sorted[idx] = temp

		pass++
	}

	fmt.Println("\nData Mahasiswa berdasarkan Nama Descending menggunakan Selection Sort")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("%-15s | %-25s | %-15s | %-15s\n", "NIM", "Nama", "Status", "Tunggakan")
	fmt.Println("--------------------------------------------------------------------------------")

	for i = 0; i < *jumlahMhs; i++ {
		sudahLunas = true
		totalTunggakan = 0

		for j = 0; j < 12; j++ {
			if !sorted[i].Iuran[j].Status {
				sudahLunas = false

				sisa = models.TARGET_IURAN - sorted[i].Iuran[j].TotalTerbayar
				if sisa > 0 {
					totalTunggakan += sisa
				}
			}
		}

		status = "Lunas"
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

func insertionSortTunggakanAscending(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int) {
	var sorted [models.NMAX]models.Mahasiswa
	var pass, i, j, k int
	var temp models.Mahasiswa
	var lanjutGeser bool
	var totalTemp, totalGeser, sisa int
	var sudahLunas bool
	var totalTunggakan int
	var status string

	for i = 0; i < *jumlahMhs; i++ {
		sorted[i] = daftarMhs[i]
	}

	pass = 1
	for pass <= *jumlahMhs-1 {
		i = pass
		temp = sorted[pass]
		lanjutGeser = true

		totalTemp = 0
		for k = 0; k < 12; k++ {
			if !temp.Iuran[k].Status {
				sisa = models.TARGET_IURAN - temp.Iuran[k].TotalTerbayar
				if sisa > 0 {
					totalTemp += sisa
				}
			}
		}

		for i > 0 && lanjutGeser {
			totalGeser = 0
			for k = 0; k < 12; k++ {
				if !sorted[i-1].Iuran[k].Status {
					sisa = models.TARGET_IURAN - sorted[i-1].Iuran[k].TotalTerbayar
					if sisa > 0 {
						totalGeser += sisa
					}
				}
			}

			if totalTemp < totalGeser {
				sorted[i] = sorted[i-1]
				i--
			} else {
				lanjutGeser = false
			}
		}

		sorted[i] = temp
		pass++
	}

	fmt.Println("\nData Mahasiswa berdasarkan Tunggakan Ascending menggunakan Insertion Sort")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("%-15s | %-25s | %-15s | %-15s\n", "NIM", "Nama", "Status", "Tunggakan")
	fmt.Println("--------------------------------------------------------------------------------")

	for i = 0; i < *jumlahMhs; i++ {
		sudahLunas = true
		totalTunggakan = 0

		for j = 0; j < 12; j++ {
			if !sorted[i].Iuran[j].Status {
				sudahLunas = false

				sisa = models.TARGET_IURAN - sorted[i].Iuran[j].TotalTerbayar
				if sisa > 0 {
					totalTunggakan += sisa
				}
			}
		}

		status = "Lunas"
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

func selectionSortTunggakanDescending(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int) {
	var sorted [models.NMAX]models.Mahasiswa
	var pass, idx, i, j, k int
	var temp models.Mahasiswa
	var totalIdx, totalI, sisa int
	var sudahLunas bool
	var totalTunggakan int
	var status string

	for i = 0; i < *jumlahMhs; i++ {
		sorted[i] = daftarMhs[i]
	}

	pass = 1
	for pass <= *jumlahMhs-1 {
		idx = pass - 1
		i = pass

		for i < *jumlahMhs {
			totalIdx = 0
			for k = 0; k < 12; k++ {
				if !sorted[idx].Iuran[k].Status {
					sisa = models.TARGET_IURAN - sorted[idx].Iuran[k].TotalTerbayar
					if sisa > 0 {
						totalIdx += sisa
					}
				}
			}

			totalI = 0
			for k = 0; k < 12; k++ {
				if !sorted[i].Iuran[k].Status {
					sisa = models.TARGET_IURAN - sorted[i].Iuran[k].TotalTerbayar
					if sisa > 0 {
						totalI += sisa
					}
				}
			}

			if totalIdx < totalI {
				idx = i
			}
			i++
		}

		temp = sorted[pass-1]
		sorted[pass-1] = sorted[idx]
		sorted[idx] = temp

		pass++
	}

	fmt.Println("\nData Mahasiswa berdasarkan Tunggakan Descending menggunakan Selection Sort")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("%-15s | %-25s | %-15s | %-15s\n", "NIM", "Nama", "Status", "Tunggakan")
	fmt.Println("--------------------------------------------------------------------------------")

	for i = 0; i < *jumlahMhs; i++ {
		sudahLunas = true
		totalTunggakan = 0

		for j = 0; j < 12; j++ {
			if !sorted[i].Iuran[j].Status {
				sudahLunas = false

				sisa = models.TARGET_IURAN - sorted[i].Iuran[j].TotalTerbayar
				if sisa > 0 {
					totalTunggakan += sisa
				}
			}
		}

		status = "Lunas"
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
