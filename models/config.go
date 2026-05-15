package models

const NMAX = 1000
const TARGET_IURAN = 100000

type Mahasiswa struct {
	NIM   string
	Nama  string
	NoHP  string
	Iuran [12]IuranBulan
}

type IuranBulan struct {
	TotalTerbayar int
	Status        bool
	Riwayat       [NMAX]Transaksi
	JumlahRiwayat int
}

type Transaksi struct {
	Nominal int
	Tanggal string
}
