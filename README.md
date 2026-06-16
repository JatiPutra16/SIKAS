# 💰 SIKAS
## Sistem Informasi Kas Mahasiswa

![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-Academic-blue?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-Development-success?style=for-the-badge)

SIKAS (**Sistem Informasi Kas Mahasiswa**) adalah aplikasi berbasis **Golang** yang dirancang untuk membantu bendahara kelas dalam mengelola administrasi kas mahasiswa secara digital.

Aplikasi ini dikembangkan sebagai **Tugas Besar Mata Kuliah Algoritma Pemrograman 2** dengan mengimplementasikan konsep **CRUD**, **Searching**, dan **Sorting** menggunakan algoritma dasar.

---

## ✨ Features

- 👨‍🎓 Manajemen Data Mahasiswa (Create, Read, Update, Delete)
- 💵 Pencatatan Pembayaran Kas
- 📅 Penyimpanan Riwayat Tanggal Pembayaran
- 🔍 Sequential Search
- 🔎 Binary Search
- 📊 Selection Sort
- 📈 Insertion Sort
- 💰 Statistik Total Saldo Kas
- ✅ Statistik Mahasiswa yang Sudah Lunas
- 📋 Tampilan Data yang Terstruktur

---

## 🧠 Algoritma yang Digunakan

| Algoritma | Fungsi |
|------------|----------------------------|
| Sequential Search | Mencari mahasiswa yang belum membayar |
| Binary Search | Mencari data mahasiswa pada data yang telah diurutkan |
| Selection Sort | Mengurutkan data berdasarkan nama atau tunggakan |
| Insertion Sort | Mengurutkan data secara efisien pada data kecil |

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Paradigm:** Procedural Programming
- **Data Structure:** Struct & Slice
- **CLI Application**
- **Standard Library Only**

---

## 📂 Project Structure

```
SIKAS/
│
├── main.go
├── models/
│   ├── mahasiswa.go
│   └── transaksi.go
│
├── services/
│   ├── crud.go
│   ├── search.go
│   ├── sorting.go
│   └── statistik.go
│
├── utils/
│   └── helper.go
│
├── data/
│   └── mahasiswa.json
│
└── README.md
```

> Struktur dapat berubah sesuai perkembangan proyek.

---

## 🚀 Installation

Clone repository

```bash
git clone https://github.com/JatiPutra16/SIKAS.git
```

Masuk ke folder project

```bash
cd SIKAS
```

Jalankan aplikasi

```bash
go run .
```

atau

```bash
go run main.go
```

Build executable

```bash
go build
```

Jalankan hasil build

```bash
./SIKAS
```

---

## 📊 Studi Kasus

SIKAS dibuat untuk mempermudah bendahara kelas dalam:

- Mengelola data mahasiswa
- Mencatat pembayaran kas
- Mengetahui mahasiswa yang belum membayar
- Mengurutkan data berdasarkan nama atau tunggakan
- Menghitung total saldo kas secara otomatis

Dengan digitalisasi proses tersebut, pengelolaan kas menjadi lebih **efisien**, **akurat**, dan **mudah dipantau**.

---

## 🎯 Tujuan Pembelajaran

Project ini mengimplementasikan materi:

- Algoritma Pemrograman
- Searching Algorithm
- Sorting Algorithm
- Struct dan Slice pada Golang
- Pengolahan Data
- Modular Programming

---


## 👥 Team

**Kelompok Tugas Besar**

Mata Kuliah **Algoritma Pemrograman 2**

Program Studi Teknik Informatika

---

## 📄 License

Project ini dibuat untuk keperluan akademik dan pembelajaran.

Silakan digunakan sebagai referensi dengan tetap mencantumkan sumber yang sesuai.

---

<div align="center">

# 💰 SIKAS

### Sistem Informasi Kas Mahasiswa

*Aplikasi sederhana berbasis Golang untuk pengelolaan kas mahasiswa yang cepat, efisien, dan terstruktur.*

</div>
