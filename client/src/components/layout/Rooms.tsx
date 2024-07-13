import React from 'react'

const Rooms: React.FC = () => {
  return (
    <div className='bg-gray-800 text-purple-lighter flex-none w-64 hidden md:block'>
      <div className='text-white px-6 py-2 flex border-b border-gray-600 shadow-xl'>
        <h3 className='flex items-center justify-center h-9 font-semibold text-xl leading-tight truncate'>
          メッセージ
        </h3>
      </div>
      {/* TODO: list user (room) who are chatting with me */}
      <div className='px-6 py-2'></div>
    </div>
  )
}

export default Rooms
