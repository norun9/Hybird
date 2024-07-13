import React from 'react'
import Image from 'next/image'

const MessageInput: React.FC = () => {
  return (
    <div className='bg-gray-700 flex-none pb-5 px-4'>
      <div className='flex rounded-lg overflow-hidden'>
        <input type='text' className='w-full px-4 bg-gray-600 text-white' placeholder='Message...' />
        <button className='border-l border-gray-600 w-[4rem] flex flex-row items-center justify-center text-3xl text-grey border-r-4 border-gray-600 bg-gray-600 p-2'>
          <Image
            src='/assets/icon/message/send_disabled.svg' // Route of the image file
            height={25} // Desired size with correct aspect ratio
            width={25} // Desired size with correct aspect ratio
            alt='message_send_icon'
          />
        </button>
      </div>
    </div>
  )
}

export default MessageInput
