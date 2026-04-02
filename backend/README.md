# Backend Go CRUD

Backend ini dibuat dengan:
- Go
- Gin
- sqlx
- PostgreSQL

Arsitektur:
- Layered (`entities -> repositories -> services -> handlers`)

## Relasi Tabel

```text
authors (1) --> books (1) --> reviews
```

## Struktur Folder

```text
backend/
|-- cmd/
|   `-- api/
|       `-- main.go
|-- config/
|   `-- database.go
|-- entities/
|   |-- author.go
|   |-- book.go
|   |-- review.go
|   `-- explorer.go
|-- repositories/
|   |-- author_repository.go
|   |-- book_repository.go
|   `-- review_repository.go
|-- services/
|   |-- author_service.go
|   |-- book_service.go
|   `-- review_service.go
|-- handlers/
|   |-- author_handler.go
|   |-- book_handler.go
|   `-- review_handler.go
`-- migrations/
    `-- init.sql
```

## Setup Cepat

### 1. Buat Database

```bash
createdb go_crud
psql go_crud < migrations/init.sql
```

### 2. Configure Environment

Backend akan mencoba membaca file `.env` secara otomatis saat startup.

Contoh file `.env`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=go_crud
DB_SSLMODE=disable
DB_SEARCH_PATH=public
CORS_ORIGIN=http://localhost:3000
```

Alternatifnya, env tetap bisa di-set manual dari shell.

Contoh PowerShell:

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

Catatan:
- File `.env` bersifat opsional.
- Environment variable dari shell tetap menimpa nilai dari `.env`.
- `DB_SEARCH_PATH` sebaiknya `public` agar sesuai migration default.

### 3. Jalankan Backend

```bash
go mod tidy
go run ./cmd/api
```

Server:
- `http://localhost:8080`
- Base API: `http://localhost:8080/api/v1`

## Endpoint API

### Authors

| Method | URL | Keterangan |
|---|---|---|
| GET | `/api/v1/authors` | List semua author |
| GET | `/api/v1/authors/:id` | Detail author |
| POST | `/api/v1/authors` | Buat author |
| PUT | `/api/v1/authors/:id` | Update author |
| DELETE | `/api/v1/authors/:id` | Hapus author |

### Books

| Method | URL | Keterangan |
|---|---|---|
| GET | `/api/v1/books` | List semua buku |
| GET | `/api/v1/books?author_id=1` | Filter buku by author |
| GET | `/api/v1/books/explorer` | JOIN authors + books + reviews |
| GET | `/api/v1/books/explorer?author_id=1&min_rating=4` | Filter explorer |
| GET | `/api/v1/books/:id` | Detail buku |
| POST | `/api/v1/books` | Buat buku |
| PUT | `/api/v1/books/:id` | Update buku |
| DELETE | `/api/v1/books/:id` | Hapus buku |

### Reviews

| Method | URL | Keterangan |
|---|---|---|
| GET | `/api/v1/reviews` | List semua review |
| GET | `/api/v1/reviews?book_id=1` | Filter review by buku |
| GET | `/api/v1/reviews/:id` | Detail review |
| POST | `/api/v1/reviews` | Buat review |
| PUT | `/api/v1/reviews/:id` | Update review |
| DELETE | `/api/v1/reviews/:id` | Hapus review |

## Contoh Request

### Create Author

```bash
curl -X POST http://localhost:8080/api/v1/authors \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Pramoedya Ananta Toer\",\"email\":\"pram@example.com\",\"bio\":\"Penulis Indonesia\"}"
```

### Create Book

```bash
curl -X POST http://localhost:8080/api/v1/books \
  -H "Content-Type: application/json" \
  -d "{\"author_id\":1,\"title\":\"Bumi Manusia\",\"description\":\"Novel sejarah\",\"published_year\":1980}"
```

### Create Review

```bash
curl -X POST http://localhost:8080/api/v1/reviews \
  -H "Content-Type: application/json" \
  -d "{\"book_id\":1,\"reviewer\":\"Budi\",\"rating\":5,\"comment\":\"Masterpiece!\"}"
```

### Explorer

```bash
curl "http://localhost:8080/api/v1/books/explorer?author_id=1&min_rating=4"
```

## Alur Data

```text
HTTP Request
  -> Handler (parse input, return HTTP response)
  -> Service (business logic, validasi relasi)
  -> Repository (SQL query)
  -> PostgreSQL
```

## Troubleshooting

1. `relation "... " does not exist`
- Jalankan migration `migrations/init.sql`.
- Pastikan `DB_SEARCH_PATH` benar (`public`).

2. Error CORS di frontend
- Pastikan `CORS_ORIGIN` memuat origin frontend.
- Restart backend setelah ubah env.

3. Endpoint baru belum muncul
- Pastikan backend direstart setelah perubahan kode.
