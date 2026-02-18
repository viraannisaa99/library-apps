'use client'

import { useCallback, useEffect, useMemo, useState } from 'react'
import toast from 'react-hot-toast'
import { Search } from 'lucide-react'
import { authorsApi, booksApi } from '@/lib/api'
import type { Author, BookExplorer } from '@/types'
import { EmptyState, LoadingSpinner, PageHeader, StarRating } from '@/components/ui'

const minRatingOptions = [0, 2, 3, 4, 4.5]

function formatDate(value?: string) {
  if (!value) return '-'
  const date = new Date(value)
  if (Number.isNaN(date.getTime())) return '-'
  return date.toLocaleDateString('id-ID')
}

export default function ExplorerPage() {
  const [rows, setRows] = useState<BookExplorer[]>([])
  const [authors, setAuthors] = useState<Author[]>([])
  const [loading, setLoading] = useState(true)
  const [selectedAuthor, setSelectedAuthor] = useState('')
  const [minRating, setMinRating] = useState('0')

  const fetchAuthors = useCallback(async () => {
    try {
      const res = await authorsApi.getAll()
      setAuthors(res.data)
    } catch {
      toast.error('Gagal memuat daftar author')
    }
  }, [])

  const fetchExplorer = useCallback(async () => {
    setLoading(true)
    try {
      const authorId = selectedAuthor ? Number(selectedAuthor) : undefined
      const parsedMinRating = Number(minRating)
      const res = await booksApi.getExplorer({
        authorId,
        minRating: parsedMinRating > 0 ? parsedMinRating : undefined,
      })
      setRows(res.data)
    } catch (e: any) {
      toast.error(e.message)
      setRows([])
    } finally {
      setLoading(false)
    }
  }, [selectedAuthor, minRating])

  useEffect(() => {
    fetchAuthors()
  }, [fetchAuthors])

  useEffect(() => {
    fetchExplorer()
  }, [fetchExplorer])

  const summary = useMemo(() => {
    const totalBooks = rows.length
    const totalReviews = rows.reduce((acc, item) => acc + item.review_count, 0)
    const avgRating =
      totalBooks === 0 ? 0 : rows.reduce((acc, item) => acc + item.avg_rating, 0) / totalBooks

    return { totalBooks, totalReviews, avgRating }
  }, [rows])

  return (
    <div className="max-w-6xl mx-auto">
      <PageHeader
        title="Library Explorer"
        description="Data gabungan dari authors + books + reviews (JOIN + aggregate)."
      />

      <div className="card p-4 mb-4">
        <div className="flex flex-col md:flex-row gap-3 md:items-end">
          <div className="flex-1">
            <label className="label">Filter Author</label>
            <select
              className="input"
              value={selectedAuthor}
              onChange={(e) => setSelectedAuthor(e.target.value)}
            >
              <option value="">Semua author</option>
              {authors.map((author) => (
                <option key={author.id} value={author.id}>
                  {author.name}
                </option>
              ))}
            </select>
          </div>

          <div className="w-full md:w-64">
            <label className="label">Min Rating</label>
            <select className="input" value={minRating} onChange={(e) => setMinRating(e.target.value)}>
              {minRatingOptions.map((value) => (
                <option key={value} value={value}>
                  {value === 0 ? 'Semua rating' : `>= ${value}`}
                </option>
              ))}
            </select>
          </div>

          <button className="btn btn-primary md:w-auto" onClick={fetchExplorer}>
            <Search className="w-4 h-4" />
            Refresh
          </button>
        </div>
      </div>

      <div className="grid grid-cols-1 sm:grid-cols-3 gap-3 mb-4">
        <div className="card p-4">
          <p className="text-xs text-gray-500 mb-1">Total Book</p>
          <p className="text-2xl font-semibold">{summary.totalBooks}</p>
        </div>
        <div className="card p-4">
          <p className="text-xs text-gray-500 mb-1">Total Review</p>
          <p className="text-2xl font-semibold">{summary.totalReviews}</p>
        </div>
        <div className="card p-4">
          <p className="text-xs text-gray-500 mb-1">Rata-rata Rating</p>
          <p className="text-2xl font-semibold">{summary.avgRating.toFixed(2)}</p>
        </div>
      </div>

      <div className="card overflow-x-auto">
        {loading ? (
          <LoadingSpinner text="Memuat data explorer..." />
        ) : rows.length === 0 ? (
          <EmptyState message="Belum ada data gabungan yang cocok dengan filter." />
        ) : (
          <table className="w-full text-sm">
            <thead className="bg-gray-50 text-gray-500">
              <tr>
                <th className="text-left px-4 py-3">Book</th>
                <th className="text-left px-4 py-3">Author</th>
                <th className="text-left px-4 py-3">Published</th>
                <th className="text-left px-4 py-3">Review Count</th>
                <th className="text-left px-4 py-3">Avg Rating</th>
                <th className="text-left px-4 py-3">Last Review</th>
              </tr>
            </thead>
            <tbody>
              {rows.map((row) => (
                <tr key={row.book_id} className="border-t border-gray-100">
                  <td className="px-4 py-3">
                    <p className="font-medium text-gray-900">{row.book_title}</p>
                    <p className="text-xs text-gray-500">Book ID: {row.book_id}</p>
                  </td>
                  <td className="px-4 py-3">
                    <p className="font-medium text-gray-900">{row.author_name}</p>
                    <p className="text-xs text-gray-500">{row.author_email}</p>
                  </td>
                  <td className="px-4 py-3">{row.published_year || '-'}</td>
                  <td className="px-4 py-3">{row.review_count}</td>
                  <td className="px-4 py-3">
                    <div className="flex items-center gap-2">
                      <StarRating rating={Math.round(row.avg_rating)} />
                      <span className="text-xs text-gray-500">({row.avg_rating.toFixed(2)})</span>
                    </div>
                  </td>
                  <td className="px-4 py-3">{formatDate(row.last_review_at)}</td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
    </div>
  )
}
