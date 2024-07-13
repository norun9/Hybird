import { IMessageRes } from '@/types/api/response'

export interface MessageGroup {
  [date: string]: IMessageRes
}
