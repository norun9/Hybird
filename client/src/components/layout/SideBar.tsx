import React from 'react'
import Image from 'next/image'

const SideBar: React.FC = () => {
  return (
    <div className='bg-gray-900 text-purple-lighter flex-none w-24 p-6 hidden md:block'>
      <div className='cursor-pointer mb-4 border-b border-gray-600 pb-2'>
        <div className='bg-white h-12 w-12 flex items-center justify-center text-black text-2xl font-semibold rounded-3xl mb-1 overflow-hidden'>
          <Image src='/assets/icon/user/free.svg' height={144} width={144} alt='free_icon' />
        </div>
      </div>
    </div>
  )
}

export default SideBar
