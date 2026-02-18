# Backend Architecture Guide (Beginner Friendly)

Dokumen ini menjelaskan cara kerja backend di project ini untuk pemula yang belum terlalu familiar dengan Go.

Tujuan utama:
- paham struktur kode backend
- paham alur request dari client sampai database
- bisa menambah endpoint API baru dengan pola yang konsisten

## 1. Tech Stack

- Language: Go
- HTTP framework: Gin
- DB access: sqlx
- Database: PostgreSQL

File acuan:
- `backend/main.go`
- `backend/config/database.go`
- `backend/entities/*`
- `backend/repositories/*`
- `backend/services/*`
- `backend/handlers/*`

## 2. Gambaran Arsitektur Layered

Project backend memakai pola layered:

```text
Client Request
  -> Handler
  -> Service
  -> Repository
  -> PostgreSQL
  -> Response ke Client
```

Peran setiap layer:

1. `Handler`
- menerima HTTP request (path, query, body)
- validasi input dasar
- panggil service
- bentuk HTTP response (`200`, `400`, `404`, `500`)

2. `Service`
- tempat business logic
- validasi aturan domain (misal relasi harus ada)
- ubah error teknis jadi error yang lebih manusiawi

3. `Repository`
- fokus ke SQL
- query `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `JOIN`
- mapping hasil query ke struct `entities`

4. `Entities`
- model data dan DTO request
- jembatan antara JSON request/response dan data DB

## 3. Struktur Folder Backend

```text
backend/
├── main.go
├── config/
│   └── database.go
├── entities/
│   ├── author.go
│   ├── book.go
│   ├── review.go
│   └── explorer.go
├── repositories/
│   ├── author_repository.go
│   ├── book_repository.go
│   └── review_repository.go
├── services/
│   ├── author_service.go
│   ├── book_service.go
│   └── review_service.go
├── handlers/
│   ├── author_handler.go
│   ├── book_handler.go
│   └── review_handler.go
└── migrations/
    └── init.sql
```

## 4. Startup Flow (Apa yang terjadi saat server dinyalakan)

Lihat `backend/main.go`.

Urutan startup:

1. Buka koneksi DB lewat `config.NewDB()`.
2. Buat repository (`authorRepo`, `bookRepo`, `reviewRepo`).
3. Buat service dan inject dependency repository.
4. Buat handler dan inject dependency service.
5. Setup router Gin + middleware CORS.
6. Daftarkan route `/api/v1/...`.
7. Jalankan server di `:8080`.

Kenapa dependency di-inject?
- supaya setiap layer tidak membuat dependency sendiri
- lebih mudah dites dan di-maintain

## 5. Request Lifecycle (Contoh konkret)

Contoh request:

`GET /api/v1/books?author_id=1`

Alurnya:

1. Masuk ke `book_handler.GetAll`.
2. Handler baca query `author_id`.
3. Handler panggil `bookService.GetByAuthorID`.
4. Service validasi author harus ada.
5. Service panggil `bookRepository.FindByAuthorID`.
6. Repository jalankan SQL ke Postgres.
7. Hasil query kembali ke service -> handler -> client JSON.

## 6. Cara Menambah API Baru (Template Praktis)

Bagian ini penting untuk pemula. Ikuti urutan ini setiap kali menambah fitur API.

### Step 1 - Tentukan kebutuhan endpoint

Contoh:
- Endpoint: `GET /api/v1/books/explorer`
- Tujuan: tampilkan hasil JOIN `authors + books + reviews`
- Query optional: `author_id`, `min_rating`

### Step 2 - Siapkan model response di `entities`

Tambah struct baru di `entities`, misalnya:
- `backend/entities/explorer.go`

Kenapa?
- supaya hasil SQL punya struktur yang jelas
- response JSON konsisten

### Step 3 - Tambah kontrak di `Repository` interface

File:
- `backend/repositories/book_repository.go`

Tambahkan method baru di interface, contoh:
- `FindExplorer(authorID int, minRating float64) ([]entities.BookExplorer, error)`

Lalu implementasikan SQL-nya di struct repository.

Tips SQL:
- pakai `JOIN` untuk relasi wajib
- pakai `LEFT JOIN` kalau data child bisa kosong
- pakai `COALESCE` untuk nilai default

### Step 4 - Tambah method di `Service`

File:
- `backend/services/book_service.go`

Tambahkan method `GetExplorer(...)` untuk:
- validasi `author_id` (kalau diisi, pastikan ada)
- validasi `min_rating` (contoh range 0..5)
- panggil repository method

### Step 5 - Tambah endpoint di `Handler`

File:
- `backend/handlers/book_handler.go`

Di handler:
- parse query (`author_id`, `min_rating`)
- jika parsing gagal, return `400`
- panggil service
- map error ke status code (`404`, `400`, `500`)

### Step 6 - Daftarkan route di `main.go`

File:
- `backend/main.go`

Tambahkan route sebelum `/:id`:
- `books.GET("/explorer", bookHandler.GetExplorer)`

Catatan:
- route statis seperti `/explorer` sebaiknya ditaruh sebelum `/:id` agar tidak tertangkap sebagai parameter id.

### Step 7 - Test manual

Contoh:

```bash
curl "http://localhost:8080/api/v1/books/explorer"
curl "http://localhost:8080/api/v1/books/explorer?author_id=1&min_rating=4"
```

### Step 8 - Verifikasi compile

```bash
cd backend
go test ./...
```

Kalau lulus, endpoint siap dipakai frontend.

## 7. Cara Menghasilkan API dari Nol (Ringkas)

Jika kamu mau buat resource baru (misal `categories`), pattern-nya sama:

1. Buat tabel di migration.
2. Buat entity + request DTO.
3. Buat repository interface + SQL implementation.
4. Buat service (logic dan validasi).
5. Buat handler HTTP.
6. Wire di `main.go`.
7. Test dengan curl/Postman.
8. Hubungkan ke frontend.

## 8. Error Handling Pattern di Project Ini

Pattern yang dipakai:

- Repository mengembalikan error teknis (`sql.ErrNoRows`, dll).
- Service menerjemahkan error penting jadi pesan domain:
  - `"author not found"`
  - `"book not found"`
  - `"review not found"`
- Handler menentukan status code dari pesan error:
  - `404` untuk not found
  - `400` untuk invalid input
  - `500` untuk error internal

## 9. CORS dan Konfigurasi Environment

Backend sudah memakai CORS middleware di `backend/main.go`.

Env penting:
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`
- `DB_SSLMODE`, `DB_SEARCH_PATH`
- `CORS_ORIGIN` (bisa lebih dari satu, pisahkan dengan koma)

Contoh:

```text
CORS_ORIGIN=http://localhost:3000,http://localhost:3001
```

## 10. Checklist Saat API Baru Tidak Jalan

1. Route sudah didaftarkan di `main.go`?
2. Handler method sudah dipanggil route yang benar?
3. Service method sudah masuk interface?
4. Repository method sudah masuk interface dan ada implementasi?
5. SQL sudah sesuai nama tabel/kolom di migration?
6. Backend sudah direstart setelah perubahan kode?
7. Cek log backend untuk error detail.

## 11. Next Step untuk Pemula

Setelah paham struktur ini, lanjut belajar:

1. Unit test service dengan mock repository.
2. Pagination di list endpoint.
3. Sorting dan search query param.
4. Centralized error response format.
5. Auth sederhana (JWT) di layer middleware.

