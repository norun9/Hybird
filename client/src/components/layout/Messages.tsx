import React, { useEffect, useRef } from 'react'
import Image from 'next/image'
import { useFetch, useWebSocket } from '@/hooks'
import { IMessageRes } from '@/types/api/response'
import MessageInput from '@/components/ui/forms/MessageInput'

const Messages: React.FC = React.memo(() => {
  const { data: fetchedMessages } = useFetch<IMessageRes[]>('/v1/messages')
  const lastEventRef = useRef<HTMLDivElement>(null)
  const webSocketURL = process.env.NEXT_PUBLIC_WEB_SOCKET_URL

  const scrollToLastEvent = () => {
    if (lastEventRef.current) {
      lastEventRef.current.scrollIntoView({ behavior: 'smooth' })
    }
  }

  const { socketRef, messages, setMessages } = useWebSocket(webSocketURL)

  const groupMessagesByDate = (messages: IMessageRes[] | undefined): Record<string, IMessageRes[]> => {
    return (messages || []).reduce(
      (acc: Record<string, IMessageRes[]>, message: IMessageRes) => {
        const [date, time, ampm] = message.createdAt.split(' ')
        if (!acc[date]) {
          acc[date] = []
        }
        acc[date].push({
          ...message,
          createdAt: `${time} ${ampm}`,
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

  const sendWsMessage = (input: string) => {
    const ws = socketRef.current
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(input)
    } else {
      throw new Error('WebSocket is not open')
    }
  }

  const groupedMessages: Record<string, IMessageRes[]> = groupMessagesByDate(messages)

  useEffect(() => {
    scrollToLastEvent()
  }, [groupedMessages])

  return (
    <div className='flex-1 flex flex-col bg-gray-50 overflow-hidden'>
      <div className='border-b border-gray-border-3 flex px-6 py-2 items-center flex-none shadow-xl'>
        <div className='flex flex-col justify-center h-9'>
          <h3 className='text-white font-bold text-xl text-gray-100'>
            {/*<span className='text-gray-400'>#</span> general*/}
            {/* Other Side Username */}
          </h3>
        </div>
      </div>
      {/* Chat messages */}
      <div className='flex-1 overflow-y-auto py-4'>
        {Object.entries(groupedMessages).map(([date, msgs], dateIndex, dateArray) => (
          <div key={date}>
            <div className='flex items-center w-full mb-1'>
              <div className='flex-grow border-t border-gray-border-3'></div>
              <span className='mx-2 text-xs font-bold text-gray-400'>{date}</span>
              <div className='flex-grow border-t border-gray-border-3'></div>
            </div>
            {msgs.map((msg: IMessageRes, index: number) => {
              const isLastEvent: boolean = dateIndex === dateArray.length - 1 && index === msgs.length - 1
              return (
                <div
                  ref={isLastEvent ? lastEventRef : null}
                  key={index}
                  className='hover:bg-gray-47 px-6 pt-3 pb-7 flex items-start text-sm'>
                  <Image src='/assets/icon/user/free.svg' height={40} width={40} alt='free_icon' className='mr-3' />
                  <div className='flex-1 overflow-hidden'>
                    <div className='flex items-center'>
                      <span className='font-bold text-orange-400 cursor-pointer'>User</span>
                      <span className='font-bold text-gray-400 text-xs ml-1'>{msg.createdAt}</span>
                    </div>
                    <p className='text-gray-light leading-normal'>{msg.content}</p>
                  </div>
                </div>
              )
            })}
          </div>
        ))}
      </div>
      <MessageInput sendWsMessage={sendWsMessage} />
    </div>
  )
})

export default Messages
