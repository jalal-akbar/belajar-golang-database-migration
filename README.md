# belajar-golang-database-migration

## Agenda
- [Pengenalan Database Migration](#Pengenalan-Database-Migration)
- [Pengenalan Golang Migrate](#Pengenalan-Golang-Migrate)
- [Menginstall Golang Migrate](#Menginstall-Golang-Migrate)


- [Membuat File Migration](#Membuat-File-Migration)
- [Menjalankan Migration](#Menjalankan-Migration)
- [Rollback Migration](#Rollback-Migration)

# Pengenalan-Database-Migration
- (1) Pengenalan Database Migration
- Saat ini, kebanyakan aplikasi yang dibuat akan membutuhkan database
- Saat aplikasi berjalan, biasanya database sudah siap digunakan, artinya table, kolom dan semua relasinya sudah dibuat di awal sebelum aplikasi dijalankan
- Apa yang terjadi ketika misal pada saat aplikasi sudah berjalan, kita perlu menambah fitur baru, lalu butuh mengubah struktur table di database?
- Biasanya kita akan mengubahnya di database langsung, lalu melakukan perubahan kode program
- Hal ini mungkin terlihat sederhana, namun ketika skalanya sudah besar, dan anggota tim sudah  banyak, maka perubahan langsung ke database bukanlah hal sederhana lagi
- Kita harus bisa melakukan tracking apa saja yang berubah, dan memastikan semua anggota tim tahu perubahannya, sehingga bisa dilakukan hal yang sama di komputer masing-masing

- (2) Keuntungan Database Migration
- Oleh karena itu, Database Migration sangat diperlukan
- Database Migration adalah mekanisme untuk melakukan tracking perubahan struktur database, dari mulai awal dibuat sampai perubahan terakhir yang dilakukan
- Mirip seperti Git, dimana melakukan tracking semua perubahan kode program
- Dengan menggunakan Database Migration, semua tim member bisa melihat perubahan struktur database, dan bisa dengan mudah menjalankan perubahan tersebut di tiap komputer masing-masing
- Selain itu, dengan adanya Database Migration, kita bisa melakukan review terlebih dahulu, sebelum menjalankan perubahan di database, jaga-jaga ada perubahan yang salah, yang bisa berdampak berbahaya ke database

# Pengenalan-Golang-Migrate
- Golang Migrate adalah salah satu tool untuk Database Migration yang populer digunakan oleh programmer Golang
- Golang Migrate bisa diintegrasikan dengan aplikasi, atau dijalankan sebagai aplikasi standalone
- Golang Migrate mendukung banyak sekali database, seperti MySQL, PostgreSQL, Sqlite, MongoDB, Cassandra, dan lain-lain
```go
https://github.com/golang-migrate/migrate
```

# Menginstall-Golang-Migrate
- Untuk menginstall Golang Migrate, sangat mudah, kita bisa gunakan perintah berikut :
go install -tags ‘database1,database2’ github.com/golang-migrate/migrate/v4/cmd/migrate@latest
- Sesuaikan dengan database yang ingin kita gunakan, bisa lebih dari satu dengan cara menambahkan koma
- (1) Aplikasi Migrate
- Saat menginstall Golang Migrate, secara otomatis terdapat executable file di folder $GOPATH/bin/ dengan nama migrate
- File migrate tersebut adalah aplikasi Golang Migrate yang akan kita gunakan untuk membuat Database Migration

# Membuat-Project
- Sebagai contoh, kita akan menggunakan project Golang RESTful API yang pernah kita buat, lalu kita akan tambahkan Database Migration ke project tersebut
```go
https://github.com/jalal-akbar/golang-restful-api
```

# Membuat-Database-Migration
- (1) Membuat Database Migration
- Untuk membuat database migration, kita bisa gunakan perintah :
migrate create -ext sql -dir db/migrations nama_file_migration
- -ext adalah file extension, artinya kita membuat file .sql
- -dir adalah folder tempat disimpan
- Usahakan tidak menggunakan spasi pada nama file migration
- (2) File Migration
- File migration akan diawali dengan waktu ketika kita membuat file migration, lalu diikuti dengan nama migration, dan diakhiri dengan tipe migration
- misal 20220921103313_create_table_category.up.sql
- Kenapa diawali dengan waktu? Agar file migration selalu berurut sesuai waktu kita membuat file migration tersebut

# Migration-Up
- Saat kita membuat file database migration, file dengan akhiran up adalah file yang harus kita isi dengan perubahan yang ingin kita tambahkan
- Misal, sekarang kita akan tambahkan table category, sesuai dengan aplikasi RESTful API yang sudah kita buat

# Migration-Down
- Setiap file migration, selain file up, terdapat juga file down
- File down ini adalah file yang berisikan kode untuk mengembalikan perubahan yang kita lakukan di file up
- Kenapa ini diperlukan? Karena misal terjadi masalah di aplikasi, namun database migration terlanjur dijalankan, kita bisa melakukan rollback dengan cara menjalankan file down, karena berisikan kode untuk mengembalikan perubahan di file up
- Pada kasus ini, misal kita akan menghapus lagi table category

# Membuat-Database
- Sebelum menjalankan Database Migration, sekarang kita perlu membuat dulu database nya
- Hal ini karena pembuatan database tidak dilakukan di database migration, biasanya dilakukan manual diawal
- Pada kasus ini, kita menggunakan database mysql, dan kita perlu ubah juga kode koneksi database di aplikasi agar terhubung dengan database baru

# Menjalankan-Migration
- Selanjutnya, setelah selesai, kita bisa menjalan database migration menggunakan perintah :
- migrate -database "koneksidatabase" -path folder up
- -database harus berisikan koneksi database, misal untuk mysql, bisa menggunakan : mysql://user:password@tcp(host:port)/nama_database
- Untuk database lainnya, bisa lihat di halaman dokumentasinya : 
```go
https://github.com/golang-migrate/migrate#databases 
```
- -path adalah lokasi folder file database migration
- up adalah perintah untuk menjalankan database migration dengan mode up

# Migration-State
- Saat kita sudah melakukan migration, lalu kita menambah file migration baru, apa yang terjadi jika kita menjalankan migration lagi?
- Golang Migrate akan menyimpan state terakhir kita menjalankan database migration, artinya tidak akan dijalankan dari awal lagi, melainkan dari file terakhir yang sukses di migrasi
- Jadi kita tidak perlu takut file akan dijalankan lagi, jadi tidak perlu dihapus file migration lama-nya
- Semua informasi state tersebut disimpan dalam table schema_migrations

# Rollback-Migration
- Pada waktu misal terjadi masalah pada aplikasi, yang menyebabkan kita harus melakukan rollback perubahan, apa yang kita harus lakukan? 
- Fitur itu sudah ada di Golang Migrate, jadi kita bisa menjalankan mode down untuk melakukan rollback dengan perintah :
migrate -database "koneksidatabase" -path folder down

# Migrasi-Ke-Versi-Tertentu
# Dirty-State
- Saat kita membuat database migration, kadang kesalahan sering terjadi
- Misal saja, kita melakukan typo sehingga membuat perintah SQL nya salah
- Jika kita terlanjur menjalankan database migration, maka state akan berubah menjadi Dirty State
- State gimana kita tidak bisa melakukan up atau down lagi
- Pada kasus ini, kita harus perbaiki manual, kenapa harus manual? Karena tidak ada cara otomatis memperbaikinya
- (1) Permasalahan
- Permasalahan migration ini adalah, kita membuat dua table di file migration, pada pembuatan table pertama sukses, namun pada table kedua gagal
- Artinya file migration tidak sempurna, dan kita juga tidak bisa melakukan rollback, karena table kedua belum sukses dibuat
- Pada kondisi ini, terjadi yang namanya Dirty State, dimana kita tidak bisa melakukan up atau down, yang perlu kita lakukan adalah memperbaiki secara manual
- (2) Mengubah Versi
- Setelah kita memperbaiki secara manual, selanjutnya kita perlu mengubah versi migration di table schema_migrations
- Kita bisa lakukan manual, atau bisa otomatis menggunakan perintah :
migrate -database "koneksi_database" -path folder force versi
- Dimana versi adalah versi dari file database migration
- Pada kasus ini, kita akan gunakan satu versi sebelum migration yang gagal

# Mencoba-Aplikasi
- Sekarang sebelum kita menjalankan aplikasi, selalu jalankan database migration terlebih dahulu
- Dan selanjutnya kita bisa mencoba menjalankan aplikasinya




# Membuat-File-Migration
# Menjalankan-Migration
# Rollback-Migration
