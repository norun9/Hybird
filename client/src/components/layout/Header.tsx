import React from 'react'
import Link from 'next/link'

const HeaderComponent: React.FC = () => {
  return (
    <header className='bg-gray-900 p-5 h-20 text-white'>
      <nav className='flex items-center h-full'>
        <div className='text-xl font-bold'>
          <Link href='/'>HyBird</Link>
        </div>
      </nav>
    </header>
  )
}

export default HeaderComponent
