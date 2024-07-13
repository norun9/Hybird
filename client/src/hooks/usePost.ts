import { useState } from 'react'
import { apiClient } from '@/services'

export const usePost = <T, R>(url: string) => {
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [isError, setIsError] = useState<boolean>(false)
  const [data, setData] = useState<R | null>(null)

  const postData = async (payload: T) => {
    setIsLoading(true)
    setIsError(false)
    try {
      const response = await apiClient.post<R>(url, payload)
      setData(response.data)
    } catch (error) {
      setIsError(true)
      console.error('Failed to post data', error)
    } finally {
      setIsLoading(false)
    }
  }

  return { postData, data, isLoading, isError }
}
