import { IMessageRes } from '@/types/api/response'

export interface IMessage extends IMessageRes {
  timestamp?: number // Timestamp for checking message duplicates
}
