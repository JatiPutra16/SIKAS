package main

import (
	"fmt"
	"sikas/function"
	"sikas/models"
)

var daftarMhs [models.NMAX]models.Mahasiswa
var jumlahMhs int

func main() {
	var pilihan int
	var keluar bool

	for !keluar {
		models.Clearscreen()
		models.TampilkanHeader()
		fmt.Println("[1] Manajemen Data Mahasiswa")
		fmt.Println("[2] Pencatatan Iuran Mahasiswa")
		fmt.Println("[3] Pencarian Mahasiswa")
		fmt.Println("[4] Daftar Mahasiswa & Status Iuran")
		fmt.Println("[5] Laporan Rekapitulasi & Statistik")
		fmt.Println("[6] Keluar")
		fmt.Println("============================================")

		fmt.Print("Masukkan Pilihan: ")
		fmt.Scanln(&pilihan)
		fmt.Println()

		if pilihan == 6 {
			fmt.Println("Terima kasih telah menggunakan SIKAS!")
			keluar = true
		} else {
			switch pilihan {
			case 1:
				function.MenuMahasiswa(&daftarMhs, &jumlahMhs, models.NMAX)
			case 2:
				function.MenuPencatatanIuran(&daftarMhs, &jumlahMhs)
			case 3:
				function.MenuCariMahasiswaBelumBayar(&daftarMhs, &jumlahMhs)
			case 4:
				function.MenuDaftarMahasiswa(&daftarMhs, &jumlahMhs)
			case 5:
				function.MenuStatistikMahasiswa(&daftarMhs, &jumlahMhs)

			default:
				fmt.Println("Pilihan tidak valid.")
			}

			fmt.Print("\nTekan Enter untuk kembali ke menu...")
			fmt.Scanln()
		}
	}
}
