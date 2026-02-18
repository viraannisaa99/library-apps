# Books App (Next.js Frontend)

Frontend untuk aplikasi manajemen:
- Authors
- Books
- Reviews
- Explorer (JOIN summary)

Stack:
- Next.js 14 (App Router)
- TypeScript
- Tailwind CSS
- react-hook-form
- react-hot-toast

## Halaman

| Halaman | URL | Fungsi |
|---|---|---|
| Home | `/` | Dashboard menu |
| Authors | `/authors` | CRUD author |
| Books | `/books` | CRUD book + filter author |
| Reviews | `/reviews` | CRUD review + filter book |
| Explorer | `/explorer` | Data gabungan authors + books + reviews |

## Prasyarat

- Node.js 18+
- Backend Go sudah berjalan di `http://localhost:8080`

## Cara Menjalankan

### 1. Masuk ke folder

```bash
cd books-app
```

### 2. Set env frontend

Buat/isi `books-app/.env.local`:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
```

### 3. Install dependency

```bash
npm install
```

Jika PowerShell memblokir `npm`:

```bash
npm.cmd install
```

### 4. Jalankan dev server

```bash
npm run dev
```

Jika perlu:

```bash
npm.cmd run dev
```

Buka:
- `http://localhost:3000`

## Struktur Folder

```text
books-app/
|-- src/
|   |-- app/
|   |   |-- layout.tsx
|   |   |-- page.tsx
|   |   |-- authors/page.tsx
|   |   |-- books/page.tsx
|   |   |-- reviews/page.tsx
|   |   `-- explorer/page.tsx
|   |-- components/ui/
|   |   |-- Navbar.tsx
|   |   `-- index.tsx
|   |-- lib/
|   |   `-- api.ts
|   `-- types/
|       `-- index.ts
|-- .env.local
`-- package.json
```

## Alur Data Frontend

```text
User action di page.tsx
  -> fungsi API di src/lib/api.ts
  -> request ke backend Go
  -> response JSON
  -> update state React
  -> render ulang UI + toast
```

## Fitur Utama

- CRUD lengkap untuk Authors, Books, Reviews
- Filter Books by `author_id`
- Filter Reviews by `book_id`
- Explorer filter `author_id` + `min_rating`
- Modal form, confirm delete, loading state, empty state

## Perintah Umum

| Perintah | Fungsi |
|---|---|
| `npm run dev` | Jalankan mode development |
| `npm run build` | Build production |
| `npm run start` | Jalankan hasil build |
| `npm run lint` | Linting |

## Troubleshooting

1. `Failed to fetch`
- Cek backend aktif di `:8080`.
- Cek `NEXT_PUBLIC_API_URL` benar.
- Cek CORS backend mengizinkan `http://localhost:3000`.

2. Data tidak muncul
- Cek tab Network di browser devtools.
- Pastikan backend sudah memiliki data dan migration sudah dijalankan.

3. Env sudah diubah tapi tidak efek
- Restart dev server Next.js setelah edit `.env.local`.

4. Port 3000 dipakai aplikasi lain

```bash
npm run dev -- -p 3001
```
