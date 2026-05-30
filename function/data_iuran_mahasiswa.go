package function

import (
	"fmt"
	"sikas/models"
)

func MenuPencatatanIuran(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int) {
	var subPilih int
	var keluar bool

	for !keluar {
		models.Clearscreen()
		models.TampilkanHeader()
		fmt.Println("-------- PENCATATAN IURAN MAHASISWA --------")
		fmt.Println("[1] Tambah Iuran")
		fmt.Println("[2] Lihat Daftar Iuran")
		fmt.Println("[3] Kembali ke Menu Utama")
		fmt.Println("============================================")
		fmt.Print("Pilihan: ")
		fmt.Scanln(&subPilih)

		if subPilih == 3 {
			keluar = true
		} else {
			switch subPilih {
			case 1:
				createIuranMahasiswa(daftarMhs, jumlahMhs)

			case 2:
				readIuranMahasiswa(daftarMhs, jumlahMhs)

			default:
				fmt.Println("\nPilihan tidak valid.")
			}

			fmt.Print("\nTekan Enter untuk melanjutkan...")
			fmt.Scanln()
		}
	}
}

func readIuranMahasiswa(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int) {
	var nimCari string
	var namaBulan [12]string
	var idx, i, j int
	var statusP string

	namaBulan = [12]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}

	fmt.Print("Masukkan NIM Mahasiswa: ")
	fmt.Scanln(&nimCari)
	idx = -1

	for i = 0; i < *jumlahMhs; i++ {
		if daftarMhs[i].NIM == nimCari {
			idx = i
		}
	}

	if idx != -1 {
		fmt.Println("--------------------------------------------")
		fmt.Println("RIWAYAT IURAN")
		fmt.Println("NIM: ", daftarMhs[idx].NIM)
		fmt.Println("Nama Lengkap: ", daftarMhs[idx].Nama)
		fmt.Println("-------------------------------------------------------------------------------")
		fmt.Printf("%-12s | %-15s | %-15s | %-20s\n", "Bulan", "Total Bayar", "Status", "Riwayat (Nominal/Tgl)")
		fmt.Println("-------------------------------------------------------------------------------")

		for i = 0; i < 12; i++ {
			statusP = "Belum Lunas"
			if daftarMhs[idx].Iuran[i].Status {
				statusP = "Lunas"
			}

			fmt.Printf("%-12s | Rp %-12d | %-15s | ", namaBulan[i], daftarMhs[idx].Iuran[i].TotalTerbayar, statusP)

			if daftarMhs[idx].Iuran[i].JumlahRiwayat > 0 {
				for j = 0; j < daftarMhs[idx].Iuran[i].JumlahRiwayat; j++ {
					if j > 0 {
						fmt.Print(", ")
					}
					fmt.Printf("RP %d (%s)", daftarMhs[idx].Iuran[i].Riwayat[j].Nominal, daftarMhs[idx].Iuran[i].Riwayat[j].Tanggal)
				}
			} else {
				fmt.Print("-")
			}
			fmt.Println()
		}
		fmt.Println("-------------------------------------------------------------------------------")
	} else {
		fmt.Println("--------------------------------------------")
		fmt.Println("\nNIM tidak ditemukan.")
	}
}

func createIuranMahasiswa(daftarMhs *[models.NMAX]models.Mahasiswa, jumlahMhs *int) {
	var nimCari, tanggal string
	var bulan, nominal, idx, i, bulanIdx int
	var transaksiBaru models.Transaksi

	fmt.Print("Masukkan NIM Mahasiswa: ")
	fmt.Scanln(&nimCari)
	idx = -1

	for i = 0; i < *jumlahMhs; i++ {
		if daftarMhs[i].NIM == nimCari {
			idx = i
		}
	}

	if idx != -1 {
		fmt.Println("--------------------------------------------")
		fmt.Printf("Nama Mahasiswa: %s\n", daftarMhs[idx].Nama)
		fmt.Print("Masukkan Bulan (1-12): ")
		fmt.Scanln(&bulan)

		if bulan < 1 || bulan > 12 {
			fmt.Println("Bulan tidak valid! Masukkan angka 1-12.")
		} else {
			fmt.Print("Masukkan Nominal Iuran: Rp ")
			fmt.Scanln(&nominal)
			fmt.Print("Masukkan Tanggal (DD-MM-YYYY): ")
			fmt.Scanln(&tanggal)
			fmt.Println("--------------------------------------------")

			bulanIdx = bulan - 1
			daftarMhs[idx].Iuran[bulanIdx].TotalTerbayar += nominal
			daftarMhs[idx].Iuran[bulanIdx].Status = daftarMhs[idx].Iuran[bulanIdx].TotalTerbayar >= models.TARGET_IURAN

			transaksiBaru = models.Transaksi{
				Nominal: nominal,
				Tanggal: tanggal,
			}

			if daftarMhs[idx].Iuran[bulanIdx].JumlahRiwayat < models.NMAX {
				daftarMhs[idx].Iuran[bulanIdx].Riwayat[daftarMhs[idx].Iuran[bulanIdx].JumlahRiwayat] = transaksiBaru
				daftarMhs[idx].Iuran[bulanIdx].JumlahRiwayat++
			}

			fmt.Println("\nPembayaran berhasil dicatat!")
		}
	} else {
		fmt.Println("--------------------------------------------")
		fmt.Println("\nNIM tidak ditemukan.")
	}
}
