package models

func InisialisasiDummyData(daftarMhs *[NMAX]Mahasiswa, jumlahMhs *int) {
	*jumlahMhs = 5

	// Data 1: Jati Bintang (Bayar 1 bulan lunas)
	daftarMhs[0].NIM = "1301220001"
	daftarMhs[0].Nama = "Jat_Bintang"
	daftarMhs[0].NoHP = "081234567890"
	daftarMhs[0].Iuran[0] = IuranBulan{
		TotalTerbayar: 100000,
		Status:        true,
		JumlahRiwayat: 1,
	}
	daftarMhs[0].Iuran[0].Riwayat[0] = Transaksi{Nominal: 100000, Tanggal: "05-01-2024"}

	// Data 2: Nyoman Deka (Lunas Semua - 12 Bulan)
	daftarMhs[1].NIM = "1301220002"
	daftarMhs[1].Nama = "Nyoman_Deka"
	daftarMhs[1].NoHP = "081345678901"
	for i := 0; i < 12; i++ {
		daftarMhs[1].Iuran[i] = IuranBulan{
			TotalTerbayar: 100000,
			Status:        true,
			JumlahRiwayat: 1,
		}
		daftarMhs[1].Iuran[i].Riwayat[0] = Transaksi{Nominal: 100000, Tanggal: "10-01-2024"}
	}

	// Data 3: Citra Lestari (Belum bayar sama sekali)
	daftarMhs[2].NIM = "1301220003"
	daftarMhs[2].Nama = "Citra_Lestari"
	daftarMhs[2].NoHP = "081456789012"

	// Data 4: Dedi Kurniawan (Cicilan - Belum Lunas)
	daftarMhs[3].NIM = "1301220004"
	daftarMhs[3].Nama = "Dedi_Kurniawan"
	daftarMhs[3].NoHP = "081567890123"
	daftarMhs[3].Iuran[0] = IuranBulan{
		TotalTerbayar: 50000,
		Status:        false,
		JumlahRiwayat: 1,
	}
	daftarMhs[3].Iuran[0].Riwayat[0] = Transaksi{Nominal: 50000, Tanggal: "15-01-2024"}

	// Data 5: Eka Putri (Bayar 1 bulan lunas)
	daftarMhs[4].NIM = "1301220005"
	daftarMhs[4].Nama = "Eka_Putri"
	daftarMhs[4].NoHP = "081678901234"
	daftarMhs[4].Iuran[0] = IuranBulan{
		TotalTerbayar: 100000,
		Status:        true,
		JumlahRiwayat: 1,
	}
	daftarMhs[4].Iuran[0].Riwayat[0] = Transaksi{Nominal: 100000, Tanggal: "20-01-2024"}
}
