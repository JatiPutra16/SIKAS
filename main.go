package main

import (
	"fmt"
	"sikas/models"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	var pilihan int

	for {
		clearScreen()
		models.TampilkanHeader()
		fmt.Println("[1] Manajemen Data Mahasiswa")
		fmt.Println("[2] Pencatatan Iuran Mahasiswa")
		fmt.Println("[3] Pencarian Mahasiswa")
		fmt.Println("[4] Daftar Mahasiswa & Status Iuran")
		fmt.Println("[5] Laporan Rekapitulasi & Statistik")
		fmt.Println("[6] Keluar")
		fmt.Println("============================================\n")

		fmt.Print("Masukkan Pilihan: ")
		fmt.Scanln(&pilihan)
		fmt.Println()

		switch pilihan {
		case 1:

		case 2:

		case 3:

		case 4:

		case 5:

		case 6:
			fmt.Println("Terima kasih telah menggunakan SIKAS!")
			return
		}

		models.TampilkanFooter()
		fmt.Print("\nTekan Enter untuk kembali ke menu...")
		fmt.Scanln()
	}
}
