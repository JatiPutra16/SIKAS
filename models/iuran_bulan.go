package models

type IuranBulan struct {
	TotalTerbayar int
	Status        bool
	Riwayat       []Transaksi
}
