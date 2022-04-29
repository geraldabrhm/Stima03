# Tugas Besar 3 IF2211 2021/2022
### Penerapan String Matching dan Regular Expression dalam DNA Pattern Matching


Repository ini dibuat dalam rangka memenuhi Tugas Besar 3 mata kuliah IF2211 Strategi Algoritma Semester 4 Tahun 2021/2022

## Deskripsi Program
Aplikasi DNA Pattern Matching ini adalah aplikasi berbasis web yang interaktif terhadap masukan user.
Program dibuat dengan bahasa pemrograman Golang untuk bagian backend menggunakan Go Fiber Gorm untuk koneksi server dengan database, dan menggunakan framework React.js untuk bagian frontend.

Terdapat 3 fitur utama dari aplikasi ini :

* Fitur Input Disease
* Fitur DNA Test
* Fitur DNA Test Result

## Struktur Repository
Berikut adalah struktur folder dari kode program
```
root
├───dataset
├───frontend
│   ├───public
│   └───src
│       └───pages
├───server
│   ├───config
│   ├───database
│   ├───handler
│   ├───middleware
│   ├───migration
│   ├───public
│   │   ├───disease
│   │   └───patient_DNA
│   ├───route
│   ├───utils
│   └───model
│       ├───entity
│       └───request
```

## Requirements
* Telah meng-*install* MySQL
* Memiliki sebuah database kosong di MySQL 
* Membuka file ../server/config/database.go kemudian mengubah line 15 sebagai berikut
```
const MYSQL = "{username mySQL}:{password mySQL}@tcp(127.0.0.1:3306)/{nama database kosong di mySQL}?charset=utf8mb4&parseTime=True&loc=Local"
```
* Memiliki extension/plug-in browser Moesif Origin & CORS Changer serta menghidupkannya (jalan lupa untuk mematikannya setelah selesai menjalankan kode)
* Masuk ke folder frontend kemudian menjalankan command berikut di cmd
```
npm install
```
* Telah meng-*install* Golang

## How to Use
1. Clone repository ini
2. Jalankan repo hasil clone di cmd
3. Jalankan command berikut di cmd
```
cd server
cd go run main.go
```
4. Buka cmd yang baru tanpa menutup cmd yang lama kemudian jalankan repo 
5. Jalankan command berikut di cmd
```
cd frontend
cd start
```

## Anggota Kelompok
* 13520133 Jevant Jedidia Augustine
* 13520138 Gerald Abraham Sianturi
* 13520159 Atabik Muhammad Azfa Shofi
