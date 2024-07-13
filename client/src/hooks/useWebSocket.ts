import { useEffect, useRef, useState } from 'react'
import ReconnectingWebSocket from 'reconnecting-websocket'
import { IMessageRes } from '@/types/api/response'
import { getFormattedCurrentTime } from '@/utils/time'

let connectionCnt = 0

export const useWebSocket = (url: string | undefined) => {
  const socketRef = useRef<ReconnectingWebSocket | null>(null)
  const [messages, setMessages] = useState<IMessageRes[]>([])
  const initializedRef = useRef<boolean>(false)

  useEffect(() => {
    if (!url) {
      throw new Error('No websocket URL provided')
    }

    const ws = new ReconnectingWebSocket(url)
    socketRef.current = ws
    initializedRef.current = true

    ws.onopen = () => {
      connectionCnt++
      console.log('WebSocket connection opened')
    }

    const onMessage = (event: MessageEvent<string>) => {
      console.log('number of WebSocket connections:', connectionCnt)
      const newMessage = {
        content: event.data,
        createdAt: getFormattedCurrentTime(),
      }
      setMessages((prevMessages) => [...prevMessages, newMessage])
    }

    if (initializedRef.current) {
      console.log('WebSocket initialized')
      ws.addEventListener('message', onMessage)
    }

    ws.onerror = (error) => {
      console.error('WebSocket error:', error)
    }

    ws.onclose = () => {
      connectionCnt--
      console.log('WebSocket connection closed. Current connection count:', connectionCnt)
      socketRef.current = null
      initializedRef.current = false
      ws.removeEventListener('message', onMessage)
    }

    return () => {
      ws.onmessage = null
      ws.onerror = null
      ws.onclose = null
      ws.removeEventListener('message', onMessage)
      ws.close()
    }
  }, [])

  return { socketRef, messages, setMessages }
}
