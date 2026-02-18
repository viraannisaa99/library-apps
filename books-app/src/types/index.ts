// src/types/index.ts

export interface Author {
  id: number
  name: string
  email: string
  bio: string
  created_at: string
  updated_at: string
}

export interface Book {
  id: number
  author_id: number
  title: string
  description: string
  published_year: number
  created_at: string
  updated_at: string
  author?: Author
}

export interface Review {
  id: number
  book_id: number
  reviewer: string
  rating: number
  comment: string
  created_at: string
  updated_at: string
  book?: Book
}

export interface BookExplorer {
  book_id: number
  book_title: string
  published_year: number
  author_id: number
  author_name: string
  author_email: string
  review_count: number
  avg_rating: number
  last_review_at?: string
}

// Request types
export interface CreateAuthorRequest {
  name: string
  email: string
  bio?: string
}

export interface UpdateAuthorRequest {
  name?: string
  email?: string
  bio?: string
}

export interface CreateBookRequest {
  author_id: number
  title: string
  description?: string
  published_year?: number
}

export interface UpdateBookRequest {
  title?: string
  description?: string
  published_year?: number
}

export interface CreateReviewRequest {
  book_id: number
  reviewer: string
  rating: number
  comment?: string
}

export interface UpdateReviewRequest {
  reviewer?: string
  rating?: number
  comment?: string
}

// API response wrapper
export interface ApiResponse<T> {
  data: T
  error?: string
  message?: string
}
