# Tes Masuk Freelance Backend - Waizly

Repositori ini berisi kode dan hasil tes untuk posisi freelance backend di Waizly. Tes ini terdiri dari dua bagian:

*   **Problem Solving:** Berisi solusi untuk soal-soal pemrograman dasar yang menguji kemampuan logika dan algoritma.
*   **Implementation Backend (Laravel):** Berisi implementasi backend menggunakan framework Laravel sesuai dengan spesifikasi yang diberikan.

## Struktur Repositori

*   `problem-solving/`: Direktori yang berisi file-file Go (`.go`) untuk setiap soal problem solving.
*   `implementation-test/`: Direktori yang berisi struktur proyek Laravel.
*   `implementation-test-2.txt`: File teks yang berisi hasil dari implementasi test 2.

## Cara Menjalankan

1.  **Problem Solving:**
    *   Pastikan Anda memiliki Go terinstal di sistem Anda.
    *   Buka terminal dan arahkan ke direktori `problem-solving`.
    *   Jalankan setiap file `.go` menggunakan perintah `go run nama_file.go`.

2.  **Implementation Backend (Laravel):**
    *   Pastikan Anda memiliki PHP dan Composer terinstal di sistem Anda.
    *   Buka terminal dan arahkan ke direktori `implementation-test`.
    *   Jalankan perintah `composer install` untuk menginstal dependensi Laravel.
    *   Jalankan perintah `php artisan key:generate` untuk membuat kunci aplikasi (app key) yang aman.
    *   **Konfigurasi Database:**
        *   Buka file `.env` di direktori `implementation-test`.
        *   Sesuaikan pengaturan database (misalnya, `DB_CONNECTION`, `DB_HOST`, `DB_PORT`, `DB_DATABASE`, `DB_USERNAME`, `DB_PASSWORD`) sesuai dengan konfigurasi database Anda.
    *   Jalankan perintah `php artisan migrate` untuk menjalankan perintah migrasi ke database anda
    *   Jalankan perintah `php artisan serve` untuk menjalankan server development Laravel.
    *   Backend akan berjalan pada alamat `http://localhost:8000`.

## Hasil Tes

Hasil dari implementasi test 2 dapat dilihat pada file `implementation-test-2.txt`.


Terima kasih telah mempertimbangkan saya untuk posisi freelance backend di Waizly.