import React from 'react'

const Rooms: React.FC = () => {
  return (
    <div className='bg-gray-43 text-purple-lighter flex-none w-64 hidden md:block'>
      <div className='px-6 py-2 flex border-b border-gray-border-2 shadow-xl'>
        <h3 className='flex text-xl text-gray-light items-center justify-center h-9 font-semibold leading-tight truncate'></h3>
      </div>
      {/* TODO: list user (room) who are chatting with me */}
      <div className='px-6 py-2'></div>
    </div>
  )
}

export default Rooms
