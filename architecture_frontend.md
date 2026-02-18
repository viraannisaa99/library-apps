# Frontend Architecture Guide (Beginner Friendly)

Dokumen ini menjelaskan arsitektur frontend `books-app` untuk pemula yang belum terlalu familiar dengan Next.js.

Tujuan:
- paham struktur project frontend
- paham alur data dari UI ke backend API
- bisa menambah halaman/fitur baru dengan pola yang konsisten

## 1. Tech Stack Frontend

- Framework: Next.js 14 (App Router)
- Language: TypeScript
- Styling: Tailwind CSS
- Form: react-hook-form
- Notifikasi: react-hot-toast
- Icons: lucide-react

## 2. Struktur Folder Penting

```text
books-app/src/
├── app/
│   ├── layout.tsx
│   ├── page.tsx
│   ├── authors/page.tsx
│   ├── books/page.tsx
│   ├── reviews/page.tsx
│   ├── explorer/page.tsx
│   └── globals.css
├── components/
│   └── ui/
│       ├── Navbar.tsx
│       └── index.tsx
├── lib/
│   └── api.ts
└── types/
    └── index.ts
```

## 3. Konsep Besar Arsitektur Frontend

Frontend ini memakai pembagian tanggung jawab sederhana:

1. `src/types/index.ts`
- semua type data (Author, Book, Review, BookExplorer, request payload)

2. `src/lib/api.ts`
- satu tempat untuk semua HTTP request ke backend
- setiap resource punya object API sendiri: `authorsApi`, `booksApi`, `reviewsApi`

3. `src/app/**/page.tsx`
- halaman fitur (Authors, Books, Reviews, Explorer)
- mengelola state UI, form, modal, loading, dan event user

4. `src/components/ui/*`
- komponen reusable seperti Navbar, Modal, ConfirmDialog, LoadingSpinner

## 4. Alur Data End-to-End di Frontend

Alur saat user klik tombol di halaman:

```text
User Action (klik submit / filter / delete)
  ->
Page Component (state + handler + validasi form)
  ->
lib/api.ts (fetch ke backend)
  ->
Backend Go API
  ->
response JSON
  ->
Page Component update state
  ->
UI re-render + toast feedback
```

Contoh nyata:
- `authors/page.tsx` memanggil `authorsApi.getAll()`
- hasil disimpan ke state `authors`
- UI menampilkan card list author

## 5. Root Layout dan Shared UI

File: `src/app/layout.tsx`

Fungsi:
- inject `Navbar` untuk semua halaman
- set area konten utama (`<main>`)
- pasang `Toaster` global untuk notifikasi

Artinya, halaman fitur tidak perlu setup navbar/toast berulang.

## 6. API Layer Pattern (`src/lib/api.ts`)

Semua call backend lewat helper `request<T>()`:

1. ambil `BASE_URL` dari `NEXT_PUBLIC_API_URL`
2. jalankan `fetch`
3. parse JSON
4. kalau status bukan `2xx`, throw error dari `json.error`
5. kembalikan data typed (`Promise<T>`)

Kenapa ini penting:
- kode request jadi terpusat
- konsisten antar halaman
- mudah di-debug

## 7. State Management Pattern per Halaman

Setiap halaman umumnya punya state berikut:

- `loading` untuk initial fetch
- `saving` untuk submit form
- `deleting` untuk hapus data
- `modalOpen` untuk modal create/edit
- `editTarget` dan `deleteTarget` untuk item aktif

Pattern ini terlihat di:
- `src/app/authors/page.tsx`
- `src/app/books/page.tsx`
- `src/app/reviews/page.tsx`

## 8. Form Handling Pattern

Pola form memakai `react-hook-form`:

1. deklarasi `useForm<T>()`
2. register field dengan validasi
3. `handleSubmit(onSubmit)`
4. saat sukses, tutup modal + refetch data + toast sukses
5. saat gagal, tampilkan toast error

Kelebihan:
- validasi ringan
- kode lebih ringkas dibanding controlled form penuh

## 9. Explorer Page sebagai Contoh Feature Lanjutan

File: `src/app/explorer/page.tsx`

Halaman ini mengajarkan:
- filter data (author dan min rating)
- query param ke backend (`booksApi.getExplorer`)
- summary data dengan `useMemo`
- tampilkan data gabungan dari tiga tabel

Ini contoh bagus untuk belajar “data transformation” di frontend.

## 10. Cara Menambah Fitur Frontend Baru (Step-by-Step)

Contoh: tambah halaman baru `/categories`.

### Step 1 - Tambah type

Edit `src/types/index.ts`:
- buat `Category`
- buat `CreateCategoryRequest` dan `UpdateCategoryRequest` jika perlu

### Step 2 - Tambah API methods

Edit `src/lib/api.ts`:
- tambah `categoriesApi.getAll`
- tambah `create`, `update`, `delete`

Gunakan pola yang sama seperti `authorsApi`.

### Step 3 - Buat halaman baru

Buat file:
- `src/app/categories/page.tsx`

Isi minimum:
- state list + loading
- `fetchCategories()`
- tombol tambah/edit/hapus
- modal form

### Step 4 - Pasang menu di Navbar

Edit:
- `src/components/ui/Navbar.tsx`

Tambah link route baru agar halaman bisa diakses.

### Step 5 - Test manual

1. buka halaman baru
2. create data
3. edit data
4. delete data
5. cek network response di browser devtools

## 11. Cara “Menghasilkan API” di Sisi Frontend

Di frontend, “menghasilkan API” berarti membuat **API client function**, bukan bikin server endpoint.

Template cepat function API baru:

```ts
export const categoriesApi = {
  getAll: () => request<{ data: Category[] }>('/categories'),
  create: (body: CreateCategoryRequest) =>
    request<{ data: Category }>('/categories', {
      method: 'POST',
      body: JSON.stringify(body),
    }),
}
```

Dengan ini, halaman tinggal memanggil:

```ts
const res = await categoriesApi.getAll()
setCategories(res.data)
```

## 12. Environment dan Menjalankan Frontend

File env:
- `books-app/.env.local`

Wajib isi:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
```

Jalankan:

```bash
cd books-app
npm install
npm run dev
```

Buka:
- `http://localhost:3000`

## 13. Troubleshooting Umum Frontend

1. `Failed to fetch`
- cek backend aktif di port yang benar
- cek `NEXT_PUBLIC_API_URL` benar
- cek CORS backend mengizinkan origin frontend

2. Data tidak update setelah create/update/delete
- pastikan setelah aksi memanggil ulang fungsi fetch list
- cek apakah promise await sudah dipakai

3. TypeScript error
- cek type di `src/types/index.ts` sudah sesuai response backend
- cek nama field snake_case/camelCase tidak tertukar

4. Route tidak bisa dibuka
- pastikan file route ada di `src/app/<route>/page.tsx`
- pastikan link di navbar sudah benar

## 14. Checklist Coding untuk Pemula

Saat nambah fitur baru, pastikan:

1. type sudah dibuat di `types`
2. method API sudah dibuat di `lib/api.ts`
3. halaman baru sudah ada di `app/<route>/page.tsx`
4. loading, error, success state ditangani
5. route ditambahkan di navbar bila perlu
6. flow create-edit-delete sudah dites manual

