'use client'

import { useState, useEffect, useCallback } from 'react'
import { useForm } from 'react-hook-form'
import toast from 'react-hot-toast'
import { Plus, Pencil, Trash2, Mail, User } from 'lucide-react'
import { authorsApi } from '@/lib/api'
import type { Author, CreateAuthorRequest } from '@/types'
import {
  Modal, ConfirmDialog, LoadingSpinner, EmptyState, PageHeader,
} from '@/components/ui'

export default function AuthorsPage() {
  const [authors, setAuthors] = useState<Author[]>([])
  const [loading, setLoading] = useState(true)
  const [modalOpen, setModalOpen] = useState(false)
  const [editTarget, setEditTarget] = useState<Author | null>(null)
  const [deleteTarget, setDeleteTarget] = useState<Author | null>(null)
  const [deleting, setDeleting] = useState(false)
  const [saving, setSaving] = useState(false)

  const { register, handleSubmit, reset, formState: { errors } } = useForm<CreateAuthorRequest>()

  const fetchAuthors = useCallback(async () => {
    try {
      const res = await authorsApi.getAll()
      setAuthors(res.data)
    } catch (e: any) {
      toast.error(e.message)
    } finally {
      setLoading(false)
    }
  }, [])

  useEffect(() => { fetchAuthors() }, [fetchAuthors])

  const openCreate = () => {
    setEditTarget(null)
    reset({ name: '', email: '', bio: '' })
    setModalOpen(true)
  }

  const openEdit = (author: Author) => {
    setEditTarget(author)
    reset({ name: author.name, email: author.email, bio: author.bio })
    setModalOpen(true)
  }

  const onSubmit = async (data: CreateAuthorRequest) => {
    setSaving(true)
    try {
      if (editTarget) {
        await authorsApi.update(editTarget.id, data)
        toast.success('Author berhasil diupdate!')
      } else {
        await authorsApi.create(data)
        toast.success('Author berhasil ditambahkan!')
      }
      setModalOpen(false)
      fetchAuthors()
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
      await authorsApi.delete(deleteTarget.id)
      toast.success('Author berhasil dihapus!')
      setDeleteTarget(null)
      fetchAuthors()
    } catch (e: any) {
      toast.error(e.message)
    } finally {
      setDeleting(false)
    }
  }

  return (
    <div>
      <PageHeader
        title="Authors"
        description="Daftar semua penulis buku"
        action={
          <button onClick={openCreate} className="btn-primary">
            <Plus className="w-4 h-4" /> Tambah Author
          </button>
        }
      />

      {loading ? (
        <LoadingSpinner />
      ) : authors.length === 0 ? (
        <EmptyState message="Belum ada author. Tambahkan yang pertama!" />
      ) : (
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
          {authors.map((author) => (
            <div key={author.id} className="card p-5 flex flex-col gap-3">
              <div className="flex items-start justify-between gap-2">
                <div className="flex items-center gap-2 min-w-0">
                  <div className="w-9 h-9 bg-violet-100 rounded-full flex items-center justify-center flex-shrink-0">
                    <User className="w-4 h-4 text-violet-600" />
                  </div>
                  <div className="min-w-0">
                    <p className="font-semibold text-gray-900 truncate">{author.name}</p>
                    <p className="text-xs text-gray-400 flex items-center gap-1 truncate">
                      <Mail className="w-3 h-3" /> {author.email}
                    </p>
                  </div>
                </div>
              </div>
              {author.bio && (
                <p className="text-sm text-gray-500 line-clamp-2">{author.bio}</p>
              )}
              <div className="flex gap-2 mt-auto pt-2 border-t border-gray-100">
                <button onClick={() => openEdit(author)} className="btn-secondary flex-1 justify-center">
                  <Pencil className="w-3.5 h-3.5" /> Edit
                </button>
                <button onClick={() => setDeleteTarget(author)} className="btn-danger flex-1 justify-center">
                  <Trash2 className="w-3.5 h-3.5" /> Hapus
                </button>
              </div>
            </div>
          ))}
        </div>
      )}

      {/* Modal Form */}
      <Modal
        title={editTarget ? 'Edit Author' : 'Tambah Author'}
        open={modalOpen}
        onClose={() => setModalOpen(false)}
      >
        <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-4">
          <div>
            <label className="label">Nama *</label>
            <input
              className="input"
              placeholder="Nama penulis"
              {...register('name', { required: 'Nama wajib diisi' })}
            />
            {errors.name && <p className="text-red-500 text-xs mt-1">{errors.name.message}</p>}
          </div>
          <div>
            <label className="label">Email *</label>
            <input
              className="input"
              type="email"
              placeholder="email@example.com"
              {...register('email', {
                required: 'Email wajib diisi',
                pattern: { value: /^\S+@\S+\.\S+$/, message: 'Format email tidak valid' },
              })}
            />
            {errors.email && <p className="text-red-500 text-xs mt-1">{errors.email.message}</p>}
          </div>
          <div>
            <label className="label">Bio</label>
            <textarea
              className="input min-h-[80px] resize-none"
              placeholder="Deskripsi singkat penulis..."
              {...register('bio')}
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
        title="Hapus Author?"
        message={`Author "${deleteTarget?.name}" akan dihapus beserta semua bukunya. Tindakan ini tidak bisa dibatalkan.`}
        onConfirm={confirmDelete}
        onCancel={() => setDeleteTarget(null)}
        loading={deleting}
      />
    </div>
  )
}
