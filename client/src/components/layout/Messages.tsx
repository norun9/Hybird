import React, { useState, useEffect, useRef } from 'react'
import Image from 'next/image'
import { useFetch } from '@/hooks'
import { IMessageRes } from '@/types/api/response'
import MessageInput from '@/components/ui/forms/MessageInput'
import ReconnectingWebSocket from 'reconnecting-websocket'
import { getFormattedCurrentTime } from '@/utils/time'

const Messages: React.FC = () => {
  const { data: fetchedMessages } = useFetch<IMessageRes[]>('/v1/messages')
  const socketRef = useRef<ReconnectingWebSocket>()
  const [messages, setMessages] = useState<IMessageRes[]>([])

  const groupMessagesByDate = (messages: IMessageRes[] | undefined): Record<string, IMessageRes[]> => {
    return (messages || []).reduce(
      (acc: Record<string, IMessageRes[]>, message: IMessageRes) => {
        const [date, time, timeDivision] = message.createdAt.split(' ')
        if (!acc[date]) {
          acc[date] = []
        }
        acc[date].push({
          ...message,
          createdAt: time + ' ' + timeDivision,
        })
        return acc
      },
      {} as Record<string, IMessageRes[]>,
    )
  }

  useEffect(() => {
    if (fetchedMessages) {
      setMessages(fetchedMessages)
    }
  }, [fetchedMessages])

  useEffect(() => {
    const websocket: ReconnectingWebSocket = new ReconnectingWebSocket(
      `ws://${process.env.NEXT_PUBLIC_API_HOST}/v1/messages/ws`,
    )
    socketRef.current = websocket

    const onMessage = (event: MessageEvent<string>) => {
      const newMessage: IMessageRes = {
        content: event.data,
        createdAt: getFormattedCurrentTime(),
      }
      setMessages((prevMessages) => [...prevMessages, newMessage])
    }

    websocket.addEventListener('message', onMessage)

    return (): void => {
      websocket.close()
      websocket.removeEventListener('message', onMessage)
    }
  }, [])

  const sendWsMessage = (input: string) => {
    const ws: ReconnectingWebSocket | undefined = socketRef.current
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(input)
    } else {
      throw new Error('WebSocket is not open')
    }
  }

  const groupedMessages: Record<string, IMessageRes[]> = groupMessagesByDate(messages)

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
        {Object.entries(groupedMessages).map(([date, msgs]) => (
          <>
            <div className='flex items-center w-full'>
              <div className='flex-grow border-t border-gray-600'></div>
              <span className='mx-2 text-xs font-bold text-gray-400'>{date}</span>
              <div className='flex-grow border-t border-gray-600'></div>
            </div>
            {msgs.map((msg, index) => (
              <div key={index} className='border-b border-gray-600 py-3 flex items-start mb-4 text-sm'>
                <Image src='/assets/icon/user/free.svg' height={40} width={40} alt='free_icon' className='mr-3' />
                <div className='flex-1 overflow-hidden'>
                  <div>
                    <span className='font-bold text-gray-400 cursor-pointer'>User</span>
                    <span className='font-bold text-gray-400 text-xs ml-1'>{msg.createdAt}</span>
                  </div>
                  <p className='text-white leading-normal'>{msg.content}</p>
                </div>
              </div>
            ))}
          </>
        ))}
      </div>
      {/* Message Input */}
      <MessageInput sendWsMessage={sendWsMessage} />
    </div>
  )
}

export default Messages
