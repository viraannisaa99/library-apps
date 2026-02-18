'use client'

import { X, AlertTriangle, Loader2, Inbox } from 'lucide-react'
import { useEffect } from 'react'

// ─── Modal ───────────────────────────────────────────────────
interface ModalProps {
  title: string
  open: boolean
  onClose: () => void
  children: React.ReactNode
}

export function Modal({ title, open, onClose, children }: ModalProps) {
  // Close on Escape
  useEffect(() => {
    const handler = (e: KeyboardEvent) => e.key === 'Escape' && onClose()
    window.addEventListener('keydown', handler)
    return () => window.removeEventListener('keydown', handler)
  }, [onClose])

  if (!open) return null

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center p-4">
      {/* Overlay */}
      <div
        className="absolute inset-0 bg-black/40 backdrop-blur-sm"
        onClick={onClose}
      />
      {/* Panel */}
      <div className="relative bg-white rounded-2xl shadow-xl w-full max-w-lg max-h-[90vh] overflow-y-auto">
        <div className="flex items-center justify-between p-5 border-b">
          <h2 className="font-semibold text-lg">{title}</h2>
          <button onClick={onClose} className="p-1 rounded-lg hover:bg-gray-100 text-gray-400">
            <X className="w-5 h-5" />
          </button>
        </div>
        <div className="p-5">{children}</div>
      </div>
    </div>
  )
}

// ─── Confirm Dialog ───────────────────────────────────────────
interface ConfirmDialogProps {
  open: boolean
  title: string
  message: string
  onConfirm: () => void
  onCancel: () => void
  loading?: boolean
}

export function ConfirmDialog({
  open, title, message, onConfirm, onCancel, loading,
}: ConfirmDialogProps) {
  if (!open) return null
  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center p-4">
      <div className="absolute inset-0 bg-black/40 backdrop-blur-sm" onClick={onCancel} />
      <div className="relative bg-white rounded-2xl shadow-xl w-full max-w-sm p-6">
        <div className="flex items-center gap-3 mb-3">
          <div className="p-2 bg-red-100 rounded-xl">
            <AlertTriangle className="w-5 h-5 text-red-600" />
          </div>
          <h2 className="font-semibold text-lg">{title}</h2>
        </div>
        <p className="text-gray-500 text-sm mb-5">{message}</p>
        <div className="flex gap-2 justify-end">
          <button onClick={onCancel} className="btn-secondary" disabled={loading}>
            Batal
          </button>
          <button onClick={onConfirm} className="btn btn-danger" disabled={loading}>
            {loading ? <Loader2 className="w-4 h-4 animate-spin" /> : null}
            Hapus
          </button>
        </div>
      </div>
    </div>
  )
}

// ─── Loading Spinner ──────────────────────────────────────────
export function LoadingSpinner({ text = 'Memuat data...' }: { text?: string }) {
  return (
    <div className="flex flex-col items-center justify-center py-16 gap-3 text-gray-400">
      <Loader2 className="w-8 h-8 animate-spin" />
      <span className="text-sm">{text}</span>
    </div>
  )
}

// ─── Empty State ──────────────────────────────────────────────
export function EmptyState({ message }: { message: string }) {
  return (
    <div className="flex flex-col items-center justify-center py-16 gap-3 text-gray-400">
      <Inbox className="w-10 h-10" />
      <span className="text-sm">{message}</span>
    </div>
  )
}

// ─── Star Rating ──────────────────────────────────────────────
export function StarRating({ rating, max = 5 }: { rating: number; max?: number }) {
  return (
    <span className="flex gap-0.5">
      {Array.from({ length: max }).map((_, i) => (
        <svg
          key={i}
          className={`w-4 h-4 ${i < rating ? 'text-amber-400' : 'text-gray-200'}`}
          fill="currentColor"
          viewBox="0 0 20 20"
        >
          <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
        </svg>
      ))}
    </span>
  )
}

// ─── Page Header ─────────────────────────────────────────────
export function PageHeader({
  title,
  description,
  action,
}: {
  title: string
  description?: string
  action?: React.ReactNode
}) {
  return (
    <div className="flex items-start justify-between mb-6">
      <div>
        <h1 className="text-2xl font-bold text-gray-900">{title}</h1>
        {description && <p className="text-gray-500 text-sm mt-1">{description}</p>}
      </div>
      {action}
    </div>
  )
}
