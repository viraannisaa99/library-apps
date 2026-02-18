# Go CRUD Learning Project

Project ini adalah contoh **fullstack CRUD** untuk pemula.

- Backend: **Go + Gin + sqlx + PostgreSQL**
- Frontend utama: **Next.js (`books-app`)**

Fokus belajar:
- CRUD dasar (`authors`, `books`, `reviews`)
- relasi data (`author -> book -> review`)
- layered architecture di backend
- JOIN 3 tabel pada fitur `Explorer`

## Struktur Project

```text
go-crud/
|-- backend/      # API Go + Gin + PostgreSQL
|-- books-app/    # Frontend Next.js (yang dipakai)
`-- README.md
```

## Relasi Data

```text
authors (1) --> books (1) --> reviews
```

## Cara Menjalankan

### 1. Prasyarat

- Go 1.21+
- Node.js 18+
- PostgreSQL

### 2. Setup Database

```bash
createdb go_crud
psql go_crud < backend/migrations/init.sql
```

### 3. Jalankan Backend (Terminal A)

```bash
cd backend
```

Set env (PowerShell):

```powershell
$env:DB_HOST="localhost"
$env:DB_PORT="5432"
$env:DB_USER="postgres"
$env:DB_PASSWORD="postgres"
$env:DB_NAME="go_crud"
$env:DB_SSLMODE="disable"
$env:DB_SEARCH_PATH="public"
$env:CORS_ORIGIN="http://localhost:3000"
```

Jalankan backend:

```bash
go run main.go
```

Backend aktif di:
- `http://localhost:8080`
- Base API: `http://localhost:8080/api/v1`

Catatan penting:
- Backend ini **tidak** otomatis membaca file `.env` (tidak memakai `godotenv`).
- Jadi env harus di-set dari shell/terminal saat menjalankan backend.
- `DB_SEARCH_PATH` wajib sesuai schema tabel migration (default aman: `public`).

### 4. Jalankan Frontend `books-app` (Terminal B)

```bash
cd books-app
```

Buat/isi `books-app/.env.local`:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
```

Install dan jalankan:

```bash
npm install
npm run dev
```

Jika PowerShell memblokir `npm` script, gunakan:

```bash
npm.cmd install
npm.cmd run dev
```

Frontend aktif di:
- `http://localhost:3000`

### 5. Smoke Test Cepat

1. Buka `http://localhost:3000`.
2. Tambah 1 author di halaman `/authors`.
3. Tambah 1 book di `/books`.
4. Tambah 1 review di `/reviews`.
5. Cek hasil gabungan di `/explorer`.

## Alur Belajar (Disarankan)

1. Mulai dari `Authors`.
2. Lanjut ke `Books` (pakai `author_id` valid).
3. Lanjut ke `Reviews` (pakai `book_id` valid).
4. Terakhir `Explorer` untuk lihat JOIN 3 tabel.

## Alur Request End-to-End

```text
Browser (books-app)
  -> src/lib/api.ts (fetch)
  -> Gin Handler (HTTP layer)
  -> Service (business logic)
  -> Repository (SQL)
  -> PostgreSQL
```

## Fitur Utama

- CRUD Authors
- CRUD Books + filter `author_id`
- CRUD Reviews + filter `book_id`
- Explorer JOIN (`authors + books + reviews`)
- Explorer filter `author_id` dan `min_rating`
- Explorer summary `review_count`, `avg_rating`, `last_review_at`

## Endpoint Penting

- `GET /api/v1/authors`
- `GET /api/v1/books`
- `GET /api/v1/reviews`
- `GET /api/v1/books/explorer`
- `GET /api/v1/books/explorer?author_id=1&min_rating=4`

## Halaman Frontend

- `/`
- `/authors`
- `/books`
- `/reviews`
- `/explorer`

## Troubleshooting Cepat

1. `Failed to fetch` / CORS
- Pastikan backend berjalan di `:8080`.
- Pastikan `CORS_ORIGIN` memuat origin frontend (contoh `http://localhost:3000`).
- Restart backend setelah ubah env.

2. Error `relation "... " does not exist`
- Pastikan migration sudah dijalankan.
- Pastikan `DB_SEARCH_PATH` sesuai (disarankan `public`).

3. Endpoint baru belum terbaca
- Pastikan backend direstart setelah perubahan kode.

4. Frontend tidak ambil data backend
- Cek `NEXT_PUBLIC_API_URL` di `books-app/.env.local`.
- Restart Next.js dev server setelah mengubah env.

## Referensi Detail

- Backend detail: `backend/README.md`
- Frontend detail: `books-app/README.md`
- Arsitektur backend (pemula): `architecture_backend.md`
- Arsitektur frontend (pemula): `architecture_frontend.md`
