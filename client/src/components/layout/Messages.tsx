import React from 'react'
import Image from 'next/image'
import { useFetch } from '@/hooks'
import { IMessage } from '@/types'
import MessageInput from '@/components/ui/MessageInput'

const Messages: React.FC = () => {
  const { data: messages } = useFetch<IMessage[]>('/v1/messages')

  return (
    <div className='flex-1 flex flex-col bg-gray-700 overflow-hidden'>
      {/* Top bar */}
      <div className='border-b border-gray-600 flex px-6 py-2 items-center flex-none shadow-xl'>
        <div className='flex flex-col justify-center h-9'>
          <h3 className='text-white font-bold text-xl text-gray-100'>
            {/*<span className='text-gray-400'>#</span> general*/}
            {/* Other Side Username */}
          </h3>
        </div>
      </div>
      {/* Chat messages */}
      <div className='flex-1 overflow-y-auto px-6 py-4'>
        {/* A message */}
        {messages?.map((message, index) => (
          <div key={index} className='border-b border-gray-600 py-3 flex items-start mb-4 text-sm'>
            <Image src='/assets/icon/user/free.svg' height={40} width={40} alt='free_icon' className='mr-3' />
            <div className='flex-1 overflow-hidden'>
              <div>
                <span className='font-bold text-red-300 cursor-pointer'>User</span>
                <span className='font-bold text-gray-400 text-xs ml-1'>{message.createdAt}</span>
              </div>
              <p className='text-white leading-normal'>{message.content}</p>
            </div>
          </div>
        ))}
      </div>
      {/* Message input */}
      <MessageInput />
    </div>
  )
}

export default Messages
