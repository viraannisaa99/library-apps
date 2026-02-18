'use client'

import { useState, useEffect, useCallback } from 'react'
import { useForm } from 'react-hook-form'
import toast from 'react-hot-toast'
import { Plus, Pencil, Trash2, Star, BookOpen, User } from 'lucide-react'
import { reviewsApi, booksApi } from '@/lib/api'
import type { Review, Book, CreateReviewRequest } from '@/types'
import {
  Modal, ConfirmDialog, LoadingSpinner, EmptyState, PageHeader, StarRating,
} from '@/components/ui'

export default function ReviewsPage() {
  const [reviews, setReviews] = useState<Review[]>([])
  const [books, setBooks]     = useState<Book[]>([])
  const [loading, setLoading] = useState(true)
  const [filterBook, setFilterBook]     = useState<number | undefined>()
  const [modalOpen, setModalOpen]       = useState(false)
  const [editTarget, setEditTarget]     = useState<Review | null>(null)
  const [deleteTarget, setDeleteTarget] = useState<Review | null>(null)
  const [deleting, setDeleting]         = useState(false)
  const [saving, setSaving]             = useState(false)
  const [hoverRating, setHoverRating]   = useState(0)

  const { register, handleSubmit, reset, watch, setValue, formState: { errors } } =
    useForm<CreateReviewRequest>()
  const currentRating = watch('rating', 0)

  const fetchReviews = useCallback(async () => {
    setLoading(true)
    try {
      const res = await reviewsApi.getAll(filterBook)
      setReviews(res.data)
    } catch (e: any) {
      toast.error(e.message)
    } finally {
      setLoading(false)
    }
  }, [filterBook])

  useEffect(() => {
    booksApi.getAll().then((r) => setBooks(r.data)).catch(() => {})
  }, [])

  useEffect(() => { fetchReviews() }, [fetchReviews])

  const openCreate = () => {
    setEditTarget(null)
    reset({ book_id: undefined as any, reviewer: '', rating: 0, comment: '' })
    setModalOpen(true)
  }

  const openEdit = (review: Review) => {
    setEditTarget(review)
    reset({
      book_id: review.book_id,
      reviewer: review.reviewer,
      rating: review.rating,
      comment: review.comment,
    })
    setModalOpen(true)
  }

  const onSubmit = async (data: CreateReviewRequest) => {
    if (!data.rating || data.rating === 0) {
      toast.error('Rating wajib dipilih!')
      return
    }
    setSaving(true)
    try {
      const payload = { ...data, book_id: Number(data.book_id), rating: Number(data.rating) }
      if (editTarget) {
        await reviewsApi.update(editTarget.id, payload)
        toast.success('Review berhasil diupdate!')
      } else {
        await reviewsApi.create(payload)
        toast.success('Review berhasil ditambahkan!')
      }
      setModalOpen(false)
      fetchReviews()
    } catch (e: any) {
      toast.error(e.message)
    } finally {
      setSaving(false)
    }
  }

  const confirmDelete = async () => {
    if (!deleteTarget) return
    setDeleting(true)
    try {
      await reviewsApi.delete(deleteTarget.id)
      toast.success('Review berhasil dihapus!')
      setDeleteTarget(null)
      fetchReviews()
    } catch (e: any) {
      toast.error(e.message)
    } finally {
      setDeleting(false)
    }
  }

  const bookTitle = (id: number) =>
    books.find((b) => b.id === id)?.title ?? `Book #${id}`

  return (
    <div>
      <PageHeader
        title="Reviews"
        description="Daftar semua ulasan buku"
        action={
          <button onClick={openCreate} className="btn-primary">
            <Plus className="w-4 h-4" /> Tambah Review
          </button>
        }
      />

      {/* Filter by book */}
      <div className="mb-5">
        <select
          className="input max-w-xs"
          value={filterBook ?? ''}
          onChange={(e) => setFilterBook(e.target.value ? Number(e.target.value) : undefined)}
        >
          <option value="">Semua Buku</option>
          {books.map((b) => (
            <option key={b.id} value={b.id}>{b.title}</option>
          ))}
        </select>
      </div>

      {loading ? (
        <LoadingSpinner />
      ) : reviews.length === 0 ? (
        <EmptyState message="Belum ada review. Tambahkan yang pertama!" />
      ) : (
        <div className="flex flex-col gap-3">
          {reviews.map((review) => (
            <div key={review.id} className="card p-5">
              <div className="flex items-start justify-between gap-4">
                <div className="flex-1 min-w-0">
                  <div className="flex items-center gap-2 flex-wrap mb-1">
                    <span className="font-semibold text-gray-900 flex items-center gap-1">
                      <User className="w-4 h-4 text-gray-400" /> {review.reviewer}
                    </span>
                    <StarRating rating={review.rating} />
                    <span className="badge bg-amber-50 text-amber-700">
                      {review.rating}/5
                    </span>
                  </div>
                  <p className="text-xs text-gray-400 flex items-center gap-1 mb-2">
                    <BookOpen className="w-3 h-3" /> {bookTitle(review.book_id)}
                  </p>
                  {review.comment && (
                    <p className="text-sm text-gray-600">{review.comment}</p>
                  )}
                </div>
                <div className="flex gap-2 flex-shrink-0">
                  <button onClick={() => openEdit(review)} className="btn-secondary">
                    <Pencil className="w-3.5 h-3.5" />
                  </button>
                  <button onClick={() => setDeleteTarget(review)} className="btn-danger">
                    <Trash2 className="w-3.5 h-3.5" />
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}

      {/* Modal Form */}
      <Modal
        title={editTarget ? 'Edit Review' : 'Tambah Review'}
        open={modalOpen}
        onClose={() => setModalOpen(false)}
      >
        <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-4">
          <div>
            <label className="label">Buku *</label>
            <select
              className="input"
              {...register('book_id', { required: 'Buku wajib dipilih' })}
            >
              <option value="">-- Pilih Buku --</option>
              {books.map((b) => (
                <option key={b.id} value={b.id}>{b.title}</option>
              ))}
            </select>
            {errors.book_id && <p className="text-red-500 text-xs mt-1">{errors.book_id.message}</p>}
          </div>

          <div>
            <label className="label">Nama Reviewer *</label>
            <input
              className="input"
              placeholder="Nama kamu"
              {...register('reviewer', { required: 'Nama reviewer wajib diisi' })}
            />
            {errors.reviewer && <p className="text-red-500 text-xs mt-1">{errors.reviewer.message}</p>}
          </div>

          {/* Star rating picker */}
          <div>
            <label className="label">Rating *</label>
            <div className="flex gap-1">
              {[1, 2, 3, 4, 5].map((star) => (
                <button
                  key={star}
                  type="button"
                  onClick={() => setValue('rating', star)}
                  onMouseEnter={() => setHoverRating(star)}
                  onMouseLeave={() => setHoverRating(0)}
                  className="focus:outline-none"
                >
                  <Star
                    className={`w-7 h-7 transition-colors ${
                      star <= (hoverRating || currentRating)
                        ? 'text-amber-400 fill-amber-400'
                        : 'text-gray-300'
                    }`}
                  />
                </button>
              ))}
              {currentRating > 0 && (
                <span className="ml-2 text-sm text-gray-500 self-center">{currentRating}/5</span>
              )}
            </div>
          </div>

          <div>
            <label className="label">Komentar</label>
            <textarea
              className="input min-h-[80px] resize-none"
              placeholder="Tulis ulasan kamu..."
              {...register('comment')}
            />
          </div>

          <div className="flex gap-2 justify-end pt-2">
            <button type="button" onClick={() => setModalOpen(false)} className="btn-secondary">
              Batal
            </button>
            <button type="submit" className="btn-primary" disabled={saving}>
              {saving ? 'Menyimpan...' : editTarget ? 'Update' : 'Tambah'}
            </button>
          </div>
        </form>
      </Modal>

      {/* Confirm Delete */}
      <ConfirmDialog
        open={!!deleteTarget}
        title="Hapus Review?"
        message={`Review dari "${deleteTarget?.reviewer}" akan dihapus. Tindakan ini tidak bisa dibatalkan.`}
        onConfirm={confirmDelete}
        onCancel={() => setDeleteTarget(null)}
        loading={deleting}
      />
    </div>
  )
}
