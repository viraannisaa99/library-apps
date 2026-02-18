import Link from 'next/link'
import { BookOpen, Users, Star, Search } from 'lucide-react'

const menus = [
  {
    href: '/authors',
    icon: Users,
    title: 'Authors',
    description: 'Kelola data penulis buku',
    color: 'bg-violet-50 text-violet-600 border-violet-200',
    iconBg: 'bg-violet-100',
  },
  {
    href: '/books',
    icon: BookOpen,
    title: 'Books',
    description: 'Kelola koleksi buku',
    color: 'bg-sky-50 text-sky-600 border-sky-200',
    iconBg: 'bg-sky-100',
  },
  {
    href: '/reviews',
    icon: Star,
    title: 'Reviews',
    description: 'Kelola ulasan & rating buku',
    color: 'bg-amber-50 text-amber-600 border-amber-200',
    iconBg: 'bg-amber-100',
  },
  {
    href: '/explorer',
    icon: Search,
    title: 'Explorer',
    description: 'Lihat hasil JOIN author, book, dan review',
    color: 'bg-emerald-50 text-emerald-600 border-emerald-200',
    iconBg: 'bg-emerald-100',
  },
]

export default function HomePage() {
  return (
    <div className="max-w-3xl mx-auto">
      {/* Hero */}
      <div className="text-center mb-12">
        <div className="inline-flex items-center justify-center w-16 h-16 bg-sky-100 rounded-2xl mb-4">
          <BookOpen className="w-8 h-8 text-sky-600" />
        </div>
        <h1 className="text-3xl font-bold text-gray-900 mb-2">Books App</h1>
        <p className="text-gray-500">
          Aplikasi manajemen buku, penulis, dan review. Pilih menu di bawah untuk mulai.
        </p>
      </div>

      {/* Menu Cards */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
        {menus.map(({ href, icon: Icon, title, description, color, iconBg }) => (
          <Link key={href} href={href}>
            <div className={`card border p-6 hover:shadow-md transition-shadow cursor-pointer ${color}`}>
              <div className={`inline-flex p-3 rounded-xl mb-4 ${iconBg}`}>
                <Icon className="w-6 h-6" />
              </div>
              <h2 className="text-lg font-semibold mb-1">{title}</h2>
              <p className="text-sm opacity-75">{description}</p>
            </div>
          </Link>
        ))}
      </div>

      {/* Info */}
      <div className="mt-10 p-4 bg-gray-100 rounded-xl text-sm text-gray-500 text-center">
        Terhubung ke backend Go + Gin di{' '}
        <code className="bg-white px-1.5 py-0.5 rounded text-gray-700 border">
          {process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1'}
        </code>
      </div>
    </div>
  )
}
