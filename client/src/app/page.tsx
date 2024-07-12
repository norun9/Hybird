'use client'
import { useFetch } from '@/hooks/useFetch'
import { Messages } from '@/types'

const Home = () => {
  const { data: messages, isLoading, isError } = useFetch<Messages[]>('/v1/messages')
  console.log(messages)

  return <div>Hello, World</div>
}

export default Home
