'use client'

import { useState, useEffect, useCallback } from 'react'
import { useForm } from 'react-hook-form'
import toast from 'react-hot-toast'
import { Plus, Pencil, Trash2, BookOpen, User, Calendar } from 'lucide-react'
import { booksApi, authorsApi } from '@/lib/api'
import type { Book, Author, CreateBookRequest } from '@/types'
import {
  Modal, ConfirmDialog, LoadingSpinner, EmptyState, PageHeader,
} from '@/components/ui'

export default function BooksPage() {
  const [books, setBooks]     = useState<Book[]>([])
  const [authors, setAuthors] = useState<Author[]>([])
  const [loading, setLoading] = useState(true)
  const [filterAuthor, setFilterAuthor] = useState<number | undefined>()
  const [modalOpen, setModalOpen]   = useState(false)
  const [editTarget, setEditTarget] = useState<Book | null>(null)
  const [deleteTarget, setDeleteTarget] = useState<Book | null>(null)
  const [deleting, setDeleting] = useState(false)
  const [saving, setSaving]    = useState(false)

  const { register, handleSubmit, reset, formState: { errors } } = useForm<CreateBookRequest>()

  const fetchBooks = useCallback(async () => {
    setLoading(true)
    try {
      const res = await booksApi.getAll(filterAuthor)
      setBooks(res.data)
    } catch (e: any) {
      toast.error(e.message)
    } finally {
      setLoading(false)
    }
  }, [filterAuthor])

  useEffect(() => {
    authorsApi.getAll().then((r) => setAuthors(r.data)).catch(() => {})
  }, [])

  useEffect(() => { fetchBooks() }, [fetchBooks])

  const openCreate = () => {
    setEditTarget(null)
    reset({ author_id: undefined as any, title: '', description: '', published_year: undefined as any })
    setModalOpen(true)
  }

  const openEdit = (book: Book) => {
    setEditTarget(book)
    reset({
      author_id: book.author_id,
      title: book.title,
      description: book.description,
      published_year: book.published_year,
    })
    setModalOpen(true)
  }

  const onSubmit = async (data: CreateBookRequest) => {
    setSaving(true)
    try {
      const payload = {
        ...data,
        author_id: Number(data.author_id),
        published_year: data.published_year ? Number(data.published_year) : 0,
      }
      if (editTarget) {
        await booksApi.update(editTarget.id, payload)
        toast.success('Buku berhasil diupdate!')
      } else {
        await booksApi.create(payload)
        toast.success('Buku berhasil ditambahkan!')
      }
      setModalOpen(false)
      fetchBooks()
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
      await booksApi.delete(deleteTarget.id)
      toast.success('Buku berhasil dihapus!')
      setDeleteTarget(null)
      fetchBooks()
    } catch (e: any) {
      toast.error(e.message)
    } finally {
      setDeleting(false)
    }
  }

  const authorName = (id: number) =>
    authors.find((a) => a.id === id)?.name ?? `Author #${id}`

  return (
    <div>
      <PageHeader
        title="Books"
        description="Daftar semua buku"
        action={
          <button onClick={openCreate} className="btn-primary">
            <Plus className="w-4 h-4" /> Tambah Buku
          </button>
        }
      />

      {/* Filter */}
      <div className="mb-5">
        <select
          className="input max-w-xs"
          value={filterAuthor ?? ''}
          onChange={(e) => setFilterAuthor(e.target.value ? Number(e.target.value) : undefined)}
        >
          <option value="">Semua Author</option>
          {authors.map((a) => (
            <option key={a.id} value={a.id}>{a.name}</option>
          ))}
        </select>
      </div>

      {loading ? (
        <LoadingSpinner />
      ) : books.length === 0 ? (
        <EmptyState message="Belum ada buku. Tambahkan yang pertama!" />
      ) : (
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
          {books.map((book) => (
            <div key={book.id} className="card p-5 flex flex-col gap-3">
              <div className="flex items-start gap-3">
                <div className="w-10 h-10 bg-sky-100 rounded-xl flex items-center justify-center flex-shrink-0">
                  <BookOpen className="w-5 h-5 text-sky-600" />
                </div>
                <div className="min-w-0">
                  <p className="font-semibold text-gray-900 leading-snug">{book.title}</p>
                  <p className="text-xs text-gray-400 flex items-center gap-1 mt-0.5">
                    <User className="w-3 h-3" /> {authorName(book.author_id)}
                  </p>
                </div>
              </div>

              {book.description && (
                <p className="text-sm text-gray-500 line-clamp-2">{book.description}</p>
              )}

              {book.published_year > 0 && (
                <span className="badge bg-sky-50 text-sky-700 w-fit">
                  <Calendar className="w-3 h-3 mr-1" /> {book.published_year}
                </span>
              )}

              <div className="flex gap-2 mt-auto pt-2 border-t border-gray-100">
                <button onClick={() => openEdit(book)} className="btn-secondary flex-1 justify-center">
                  <Pencil className="w-3.5 h-3.5" /> Edit
                </button>
                <button onClick={() => setDeleteTarget(book)} className="btn-danger flex-1 justify-center">
                  <Trash2 className="w-3.5 h-3.5" /> Hapus
                </button>
              </div>
            </div>
          ))}
        </div>
      )}

      {/* Modal Form */}
      <Modal
        title={editTarget ? 'Edit Buku' : 'Tambah Buku'}
        open={modalOpen}
        onClose={() => setModalOpen(false)}
      >
        <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-4">
          <div>
            <label className="label">Author *</label>
            <select
              className="input"
              {...register('author_id', { required: 'Author wajib dipilih' })}
            >
              <option value="">-- Pilih Author --</option>
              {authors.map((a) => (
                <option key={a.id} value={a.id}>{a.name}</option>
              ))}
            </select>
            {errors.author_id && <p className="text-red-500 text-xs mt-1">{errors.author_id.message}</p>}
          </div>
          <div>
            <label className="label">Judul *</label>
            <input
              className="input"
              placeholder="Judul buku"
              {...register('title', { required: 'Judul wajib diisi' })}
            />
            {errors.title && <p className="text-red-500 text-xs mt-1">{errors.title.message}</p>}
          </div>
          <div>
            <label className="label">Deskripsi</label>
            <textarea
              className="input min-h-[80px] resize-none"
              placeholder="Sinopsis atau deskripsi buku..."
              {...register('description')}
            />
          </div>
          <div>
            <label className="label">Tahun Terbit</label>
            <input
              className="input"
              type="number"
              placeholder="cth: 2005"
              {...register('published_year')}
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
        title="Hapus Buku?"
        message={`Buku "${deleteTarget?.title}" akan dihapus beserta semua reviewnya. Tindakan ini tidak bisa dibatalkan.`}
        onConfirm={confirmDelete}
        onCancel={() => setDeleteTarget(null)}
        loading={deleting}
      />
    </div>
  )
}
