import React from 'react'
import Link from 'next/link'

const HeaderComponent: React.FC = () => {
  return (
    <header className='bg-blue-500 p-4 text-white'>
      <nav className='flex justify-between'>
        <div className='text-xl font-bold'>
          <Link href='/'>HyBird</Link>
        </div>
      </nav>
    </header>
  )
}

export default HeaderComponent
