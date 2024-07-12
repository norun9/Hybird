import React from 'react'
// import Image from 'next/image'

const SideBar: React.FC = () => {
  return (
    <div className='bg-gray-900 text-purple-lighter flex-none w-24 p-6 hidden md:block'>
      <div className='cursor-pointer mb-4 border-b border-gray-600 pb-2'>
        <div className='bg-white h-12 w-12 flex items-center justify-center text-black text-2xl font-semibold rounded-3xl mb-1 overflow-hidden'>
          <Image
            src='/assets/icon/user/free.svg' // Route of the image file
            height={144} // Desired size with correct aspect ratio
            width={144} // Desired size with correct aspect ratio
            alt='free_icon'
          />
        </div>
      </div>
      {/*<div className='cursor-pointer mb-4'>*/}
      {/*  <div className='bg-white h-12 w-12 flex items-center justify-center text-black text-2xl font-semibold rounded-xl mb-1 overflow-hidden'>*/}
      {/*    <img src='https://cdn.discordapp.com/embed/avatars/0.png' alt='' />*/}
      {/*  </div>*/}
      {/*</div>*/}
      {/*<div className='cursor-pointer mb-4'>*/}
      {/*  <div className='bg-white h-12 w-12 flex items-center justify-center text-black text-2xl font-semibold rounded-3xl mb-1 overflow-hidden'>*/}
      {/*    <img src='https://cdn.discordapp.com/embed/avatars/1.png' alt='' />*/}
      {/*  </div>*/}
      {/*</div>*/}
      {/*<div className='cursor-pointer mb-4'>*/}
      {/*  <div className='bg-white h-12 w-12 flex items-center justify-center text-black text-2xl font-semibold rounded-3xl mb-1 overflow-hidden'>*/}
      {/*    <img src='https://cdn.discordapp.com/embed/avatars/2.png' alt='' />*/}
      {/*  </div>*/}
      {/*</div>*/}
      {/*<div className='cursor-pointer mb-4'>*/}
      {/*  <div className='bg-white h-12 w-12 flex items-center justify-center text-black text-2xl font-semibold rounded-3xl mb-1 overflow-hidden'>*/}
      {/*    <img src='https://cdn.discordapp.com/embed/avatars/3.png' alt='' />*/}
      {/*  </div>*/}
      {/*</div>*/}
      {/*<div className='cursor-pointer mb-4'>*/}
      {/*  <div className='bg-white h-12 w-12 flex items-center justify-center text-black text-2xl font-semibold rounded-3xl mb-1 overflow-hidden'>*/}
      {/*    <img src='https://cdn.discordapp.com/embed/avatars/4.png' alt='' />*/}
      {/*  </div>*/}
      {/*</div>*/}
      {/*<div className='cursor-pointer'>*/}
      {/*  <div className='bg-white opacity-25 h-12 w-12 flex items-center justify-center text-black text-2xl font-semibold rounded-3xl mb-1 overflow-hidden'>*/}
      {/*    <svg className='fill-current h-10 w-10 block' xmlns='http://www.w3.org/2000/svg' viewBox='0 0 20 20'>*/}
      {/*      <path d='M16 10c0 .553-.048 1-.601 1H11v4.399c0 .552-.447.601-1 .601-.553 0-1-.049-1-.601V11H4.601C4.049 11 4 10.553 4 10c0-.553.049-1 .601-1H9V4.601C9 4.048 9.447 4 10 4c.553 0 1 .048 1 .601V9h4.399c.553 0 .601.447.601 1z' />*/}
      {/*    </svg>*/}
      {/*  </div>*/}
      {/*</div>*/}
    </div>
  )
}

export default SideBar
