package function

import (
	"fmt"
	"sikas/models"
)

func MenuMahasiswa(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int, maxData int) {
	var subPilih int
	var keluar bool

	for !keluar {
		models.Clearscreen()
		models.TampilkanHeader()
		fmt.Println("--------- MANAJEMEN DATA MAHASISWA ---------")
		fmt.Println("[1] Tambah Mahasiswa")
		fmt.Println("[2] Ubah Data Mahasiswa")
		fmt.Println("[3] Hapus Mahasiswa")
		fmt.Println("[4] Lihat Daftar Mahasiswa")
		fmt.Println("[5] Kembali ke Menu Utama")
		fmt.Println("============================================")
		fmt.Print("Pilihan: ")
		fmt.Scanln(&subPilih)

		if subPilih == 5 {
			keluar = true
		} else {
			switch subPilih {
			case 1:
				createMahasiswa(daftarMhs, jumlahMhs, maxData)

			case 2:
				updateMahasiswa(daftarMhs, jumlahMhs)

			case 3:
				deleteMahasiswa(daftarMhs, jumlahMhs)

			case 4:
				readMahasiswa(daftarMhs, jumlahMhs)

			default:
				fmt.Println("\nPilihan tidak valid.")
			}

			fmt.Print("\nTekan Enter untuk melanjutkan...")
			fmt.Scanln()
		}
	}
}

func readMahasiswa(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int) {
	var i int

	fmt.Println("--------------------------------------------")
	fmt.Println("DAFTAR MAHASISWA:")

	if *jumlahMhs == 0 {
		fmt.Println("(Data Masih Kosong!)")
	} else {
		fmt.Println("+----+-----------------+---------------------------+-----------------+")
		fmt.Printf("| %-2s | %-15s | %-25s | %-15s |\n", "No", "NIM", "Nama Lengkap", "No HP")
		fmt.Println("+----+-----------------+---------------------------+-----------------+")
		for i = 0; i < *jumlahMhs; i++ {
			fmt.Printf("| %-2d | %-15s | %-25s | %-15s |\n", i+1, daftarMhs[i].NIM, daftarMhs[i].Nama, daftarMhs[i].NoHP)
		}
		fmt.Println("+----+-----------------+---------------------------+-----------------+")
	}
}

func createMahasiswa(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int, maxData int) {
	var newNIM string
	var exists bool
	var i int

	fmt.Println("--------------------------------------------")
	if *jumlahMhs < maxData {
		fmt.Print("Masukkan NIM		: ")
		fmt.Scanln(&newNIM)

		exists = false
		for i = 0; i < *jumlahMhs && !exists; i++ {
			if daftarMhs[i].NIM == newNIM {
				exists = true
			}
		}

		if exists {
			fmt.Println("\nError: Mahasiswa dengan NIM tersebut sudah terdaftar!")
			return
		}

		daftarMhs[*jumlahMhs].NIM = newNIM
		fmt.Print("Masukkan Nama Lengkap : ")
		fmt.Scanln(&daftarMhs[*jumlahMhs].Nama)
		fmt.Print("Masukkan Nomor HP : ")
		fmt.Scanln(&daftarMhs[*jumlahMhs].NoHP)
		*jumlahMhs++
		fmt.Println("--------------------------------------------")
		fmt.Println("\nData mahasiswa berhasil ditambahkan.")
	} else {
		fmt.Println("--------------------------------------------")
		fmt.Println("\nKapasitas penyimpanan sudah penuh!")
	}
}

func updateMahasiswa(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int) {
	var cariNIM string
	var idx, i int

	fmt.Print("Masukkan NIM yang akan diubah: ")
	fmt.Scanln(&cariNIM)
	idx = -1

	for i = 0; i < *jumlahMhs; i++ {
		if daftarMhs[i].NIM == cariNIM {
			idx = i
		}
	}

	if idx != -1 {
		fmt.Println("--------------------------------------------")
		fmt.Print("[Update] Masukkan Nama Lengkap : ")
		fmt.Scanln(&daftarMhs[idx].Nama)
		fmt.Print("[Update] Masukkan Nomor HP   : ")
		fmt.Scanln(&daftarMhs[idx].NoHP)
		fmt.Println("--------------------------------------------")
		fmt.Println("\nData berhasil diperbaharui.")
	} else {
		fmt.Println("--------------------------------------------")
		fmt.Println("\nNIM tidak ditemukan.")
	}
}

func deleteMahasiswa(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int) {
	var hapusNIM string
	var idx, i int

	fmt.Print("Masukkan NIM yang akan dihapus: ")
	fmt.Scanln(&hapusNIM)
	idx = -1

	for i = 0; i < *jumlahMhs; i++ {
		if daftarMhs[i].NIM == hapusNIM {
			idx = i
		}
	}

	if idx != -1 {
		fmt.Println("--------------------------------------------")
		for i = idx; i < *jumlahMhs-1; i++ {
			daftarMhs[i] = daftarMhs[i+1]
		}
		*jumlahMhs--
		fmt.Printf("\nData mahasiswa dengan NIM %s berhasil dihapus.\n", hapusNIM)
	} else {
		fmt.Println("--------------------------------------------")
		fmt.Println("\nNIM tidak ditemukan.")
	}
}
