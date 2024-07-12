import { apiClient } from '@/services'
import { AxiosResponse } from 'axios'
import useSWR from 'swr'

const fetcher = <T>(url: string): Promise<T> => apiClient.get<T>(url).then((res: AxiosResponse<T>) => res.data)

export const useFetch = <T>(url: string) => {
  const { data, error } = useSWR<T>(url, fetcher)

  return {
    data,
    isLoading: !error && !data,
    isError: error,
  }
}
