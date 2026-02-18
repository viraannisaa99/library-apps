'use client'

import Link from 'next/link'
import { usePathname } from 'next/navigation'
import { BookOpen, Users, Star, Home, Search } from 'lucide-react'

const links = [
  { href: '/',         label: 'Home',    icon: Home },
  { href: '/authors',  label: 'Authors', icon: Users },
  { href: '/books',    label: 'Books',   icon: BookOpen },
  { href: '/reviews',  label: 'Reviews', icon: Star },
  { href: '/explorer', label: 'Explorer', icon: Search },
]

export default function Navbar() {
  const pathname = usePathname()

  return (
    <nav className="bg-white border-b border-gray-200 sticky top-0 z-10">
      <div className="max-w-6xl mx-auto px-4 flex items-center gap-1 h-14">
        <span className="font-bold text-sky-600 mr-4 flex items-center gap-1.5">
          <BookOpen className="w-5 h-5" /> Books App
        </span>
        {links.map(({ href, label, icon: Icon }) => {
          const active = pathname === href || (href !== '/' && pathname.startsWith(href))
          return (
            <Link
              key={href}
              href={href}
              className={`flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-sm font-medium transition-colors ${
                active
                  ? 'bg-sky-50 text-sky-700'
                  : 'text-gray-500 hover:text-gray-900 hover:bg-gray-100'
              }`}
            >
              <Icon className="w-4 h-4" />
              {label}
            </Link>
          )
        })}
      </div>
    </nav>
  )
}
