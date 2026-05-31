package function

import (
	"fmt"
	"sikas/models"
)

func MenuStatistikMahasiswa(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int) {
	var totalSaldo int
	var jumlahLunas int
	var jumlahBelumLunas int
	var totalTunggakanKelas int
	var totalTunggakan int
	var bulanLunas int
	var sisa int
	var i int
	var j int
	var sudahLunas bool

	totalSaldo = 0
	jumlahLunas = 0
	jumlahBelumLunas = 0
	totalTunggakanKelas = 0

	models.Clearscreen()
	models.TampilkanHeader()

	fmt.Println()
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("+++        STATISTIK KAS - SIKAS          +++")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++")

	for i = 0; i < *jumlahMhs; i++ {
		sudahLunas = true
		totalTunggakan = 0

		for j = 0; j < 12; j++ {
			totalSaldo += daftarMhs[i].Iuran[j].TotalTerbayar

			if !daftarMhs[i].Iuran[j].Status {
				sudahLunas = false

				sisa = models.TARGET_IURAN - daftarMhs[i].Iuran[j].TotalTerbayar
				if sisa > 0 {
					totalTunggakan += sisa
				}
			}
		}

		totalTunggakanKelas += totalTunggakan

		if sudahLunas {
			jumlahLunas++
		} else {
			jumlahBelumLunas++
		}
	}

	fmt.Println()
	fmt.Printf("  Total Mahasiswa         : %d orang\n", *jumlahMhs)
	fmt.Printf("  Mahasiswa Lunas         : %d orang\n", jumlahLunas)
	fmt.Printf("  Mahasiswa Belum Lunas   : %d orang\n", jumlahBelumLunas)
	fmt.Printf("  Total Saldo Kas         : Rp%d\n", totalSaldo)
	fmt.Printf("  Total Tunggakan Kelas   : Rp%d\n", totalTunggakanKelas)
	fmt.Printf("  Nominal Iuran/Bulan     : Rp%d\n", models.TARGET_IURAN)
	fmt.Printf("  Periode                 : %d bulan\n", 12)
	fmt.Println()
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++")

	if jumlahBelumLunas > 0 {
		fmt.Println("\nDetail mahasiswa belum lunas:")
		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Printf("%-15s | %-25s | %-12s | %-15s\n", "NIM", "Nama", "Bulan Lunas", "Tunggakan")
		fmt.Println("--------------------------------------------------------------------------------")

		for i = 0; i < *jumlahMhs; i++ {
			sudahLunas = true
			totalTunggakan = 0
			bulanLunas = 0

			for j = 0; j < 12; j++ {
				if daftarMhs[i].Iuran[j].Status {
					bulanLunas++
				} else {
					sudahLunas = false

					sisa = models.TARGET_IURAN - daftarMhs[i].Iuran[j].TotalTerbayar
					if sisa > 0 {
						totalTunggakan += sisa
					}
				}
			}

			if !sudahLunas {
				fmt.Printf("%-15s | %-25s | %2d/12 bulan | Rp %-12d\n",
					daftarMhs[i].NIM,
					daftarMhs[i].Nama,
					bulanLunas,
					totalTunggakan,
				)
			}
		}

		fmt.Println("--------------------------------------------------------------------------------")
	}

	fmt.Print("\nTekan Enter untuk melanjutkan...")
	fmt.Scanln()
}
