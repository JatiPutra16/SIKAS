package models

type Mahasiswa struct {
	NIM   string
	Nama  string
	NoHP  string
	Iuran [12]IuranBulan
}
