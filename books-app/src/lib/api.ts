// src/lib/api.ts
// Semua request ke backend Go lewat sini

const BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1'

async function request<T>(path: string, options?: RequestInit): Promise<T> {
  const res = await fetch(`${BASE_URL}${path}`, {
    headers: { 'Content-Type': 'application/json' },
    ...options,
  })

  const json = await res.json()

  if (!res.ok) {
    throw new Error(json.error || 'Terjadi kesalahan pada server')
  }

  return json
}

// ─── AUTHORS ────────────────────────────────────────────────
export const authorsApi = {
  getAll: () =>
    request<{ data: import('@/types').Author[] }>('/authors'),

  getById: (id: number) =>
    request<{ data: import('@/types').Author }>(`/authors/${id}`),

  create: (body: import('@/types').CreateAuthorRequest) =>
    request<{ data: import('@/types').Author }>('/authors', {
      method: 'POST',
      body: JSON.stringify(body),
    }),

  update: (id: number, body: import('@/types').UpdateAuthorRequest) =>
    request<{ data: import('@/types').Author }>(`/authors/${id}`, {
      method: 'PUT',
      body: JSON.stringify(body),
    }),

  delete: (id: number) =>
    request<{ message: string }>(`/authors/${id}`, { method: 'DELETE' }),
}

// ─── BOOKS ──────────────────────────────────────────────────
export const booksApi = {
  getAll: (authorId?: number) => {
    const query = authorId ? `?author_id=${authorId}` : ''
    return request<{ data: import('@/types').Book[] }>(`/books${query}`)
  },

  getExplorer: (params?: { authorId?: number; minRating?: number }) => {
    const search = new URLSearchParams()
    if (params?.authorId) search.set('author_id', String(params.authorId))
    if (params?.minRating && params.minRating > 0) {
      search.set('min_rating', String(params.minRating))
    }
    const query = search.toString() ? `?${search.toString()}` : ''
    return request<{ data: import('@/types').BookExplorer[] }>(`/books/explorer${query}`)
  },

  getById: (id: number) =>
    request<{ data: import('@/types').Book }>(`/books/${id}`),

  create: (body: import('@/types').CreateBookRequest) =>
    request<{ data: import('@/types').Book }>('/books', {
      method: 'POST',
      body: JSON.stringify(body),
    }),

  update: (id: number, body: import('@/types').UpdateBookRequest) =>
    request<{ data: import('@/types').Book }>(`/books/${id}`, {
      method: 'PUT',
      body: JSON.stringify(body),
    }),

  delete: (id: number) =>
    request<{ message: string }>(`/books/${id}`, { method: 'DELETE' }),
}

// ─── REVIEWS ────────────────────────────────────────────────
export const reviewsApi = {
  getAll: (bookId?: number) => {
    const query = bookId ? `?book_id=${bookId}` : ''
    return request<{ data: import('@/types').Review[] }>(`/reviews${query}`)
  },

  getById: (id: number) =>
    request<{ data: import('@/types').Review }>(`/reviews/${id}`),

  create: (body: import('@/types').CreateReviewRequest) =>
    request<{ data: import('@/types').Review }>('/reviews', {
      method: 'POST',
      body: JSON.stringify(body),
    }),

  update: (id: number, body: import('@/types').UpdateReviewRequest) =>
    request<{ data: import('@/types').Review }>(`/reviews/${id}`, {
      method: 'PUT',
      body: JSON.stringify(body),
    }),

  delete: (id: number) =>
    request<{ message: string }>(`/reviews/${id}`, { method: 'DELETE' }),
}
