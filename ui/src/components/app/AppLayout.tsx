'use client'

import { ReactNode, useState } from 'react'
import { usePathname } from 'next/navigation'
import { Badge, Navbar, Sidebar } from 'flowbite-react'
import { LogoIcon } from 'components/ui/LogoIcon'
import { siteConfig } from 'lib/siteConfig'
import { HiViewGrid, HiMenu, HiOutlineBookOpen } from 'react-icons/hi'
import 'styles/app.css'

type AppLayoutProps = {
  children?: ReactNode
}

export default function AppLayout({ children }: AppLayoutProps) {
  const [isSidebarOpen, setIsSidebarOpen] = useState(false)
  const pathname = usePathname()

  const navigation = [
    {
      name: 'Surveys',
      href: `/app`,
      icon: HiViewGrid,
    },
  ]

  return (
    <main className="mb-auto flex-grow dark">
      <Navbar
        fluid
        className="fixed w-full z-20 top-0 dark:bg-gray-950 border-b border-slate-6 px-5 sm:px-5"
      >
        <button
          type="button"
          className="inline-flex items-center p-2 text-sm text-gray-500 rounded-lg lg:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
          onClick={() => setIsSidebarOpen(!isSidebarOpen)}
        >
          <HiMenu className="w-6 h-6" />
        </button>

        <Navbar.Brand href="/app">
          <div className="flex gap-4">
            <LogoIcon />
          </div>
        </Navbar.Brand>
      </Navbar>

      <Sidebar
        className={`sidebar fixed z-10 top-0 h-screen dark dark:bg-gray-800 ${
          isSidebarOpen ? '' : ' hidden lg:block'
        }`}
      >
        <Sidebar.Items>
          <Sidebar.ItemGroup>
            {navigation.map((item) => (
              <Sidebar.Item
                key={item.name}
                href={item.href}
                icon={item.icon}
                className={
                  pathname && pathname.startsWith(item.href)
                    ? 'bg-gray-700'
                    : ''
                }
              >
                {item.name}
              </Sidebar.Item>
            ))}
          </Sidebar.ItemGroup>
          <Sidebar.ItemGroup>
            <Sidebar.Item
              href="https://github.com/plutov/formulosity"
              icon={HiOutlineBookOpen}
            >
              Documentation
            </Sidebar.Item>
          </Sidebar.ItemGroup>
        </Sidebar.Items>
        <Sidebar.CTA>
          <div className="mb-3 flex items-center">
            <Badge color="warning">Beta</Badge>
          </div>
          <div className="mb-3 text-sm text-cyan-900 dark:text-gray-400">
            <p>{siteConfig.name} is currently in beta.</p>
          </div>
        </Sidebar.CTA>
      </Sidebar>

      <div className="p-4 lg:ml-64 mt-16">{children}</div>
    </main>
  )
}
