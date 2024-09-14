import { useEffect, useRef, useState } from 'react'
import ReconnectingWebSocket from 'reconnecting-websocket'
import { getFormattedCurrentTime } from '@/utils/time'
import { IMessage } from '@/types/components'

export const useWebSocket = (url: string | undefined) => {
  const socketRef = useRef<ReconnectingWebSocket | null>(null)
  const [messages, setMessages] = useState<IMessage[]>([])

  useEffect(() => {
    if (!url) {
      throw new Error('No websocket URL provided')
    }

    const ws: ReconnectingWebSocket = new ReconnectingWebSocket(url)
    socketRef.current = ws

    ws.onopen = () => {
      console.log('Websocket opened')
    }

    ws.onmessage = (event: MessageEvent<string>) => {
      console.log('Received message:', event.data)
      const dataArr = event.data.split('_')
      const content = dataArr[0]
      const timestamp = Number(dataArr[1])
      const newMessage: IMessage = {
        timestamp: timestamp,
        content: content,
        createdAt: getFormattedCurrentTime(),
      }
      setMessages((prevMessages) => {
        if (!prevMessages.some((msg) => msg.content === newMessage.content && msg.timestamp === newMessage.timestamp)) {
          return [...prevMessages, newMessage]
        }
        return prevMessages
      })
    }

    ws.onerror = (error) => {
      console.error('WebSocket error:', error)
    }

    ws.onclose = () => {
      socketRef.current = null
      ws.onmessage = null
      ws.onerror = null
      ws.onclose = null
      ws.onopen = null
    }

    return () => {
      ws.onmessage = null
      ws.onerror = null
      ws.onclose = null
      ws.onopen = null
      ws.close()
    }
  }, [url])

  return { socketRef, messages, setMessages }
}
