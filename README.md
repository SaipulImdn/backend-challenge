# Online Store

## Deskripsi

Proyek ini adalah aplikasi back-end untuk toko online yang dibangun menggunakan Go-lang dan Fiber. Aplikasi ini menyediakan API untuk manajemen produk, keranjang belanja, dan transaksi.

## Fitur

- Melihat daftar produk berdasarkan kategori.
- Menambahkan produk ke keranjang belanja.
- Melihat daftar produk dalam keranjang belanja.
- Menghapus produk dari keranjang belanja.
- Melakukan checkout dan transaksi pembayaran.
- Login dan registrasi pelanggan.

## Struktur Proyek

- `cmd`: Direktori untuk main.go.
- `internal`: Direktori untuk kode aplikasi.
- `pkg`: Direktori untuk utilitas umum.
- `configs`: Direktori untuk konfigurasi.
- `migrations`: Direktori untuk migrasi database.
- `Dockerfile`: File untuk membangun Docker image.
- `docker-compose.yml`: File untuk mengatur Docker compose.

## Cara Menjalankan

1. Clone repositori ini.
2. Jalankan `docker-compose up` untuk memulai aplikasi.
3. Aplikasi akan tersedia di `http://localhost:8080`.

## Endpoint API

- `GET /products/:category`: Mendapatkan daftar produk berdasarkan kategori.
- `POST /cart`: Menambahkan produk ke keranjang belanja.
- `GET /cart`: Melihat daftar produk dalam keranjang belanja.
- `DELETE /cart/:id`: Menghapus produk dari keranjang belanja.
- `POST /checkout`: Melakukan checkout dan transaksi pembayaran.
- `POST /register`: Registrasi pelanggan.
- `POST /login`: Login pelanggan.
