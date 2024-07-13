import React, { useEffect, useState } from 'react'
import Image from 'next/image'
import { useForm, SubmitHandler } from 'react-hook-form'
import { MessageFormValues } from '@/types/components/forms'
import { usePost } from '@/hooks'
import { IMessageReq } from '@/types/api/request'
import { IMessageRes } from '@/types/api/response'

interface Props {
  sendWsMessage: (input: string) => void
}

const MessageInput: React.FC<Props> = ({ sendWsMessage }) => {
  const [valid, setValid] = useState<boolean>(false)
  const formFieldName = 'content'
  const { postData } = usePost<IMessageReq, IMessageRes>('/v1/messages')
  const {
    register,
    handleSubmit,
    formState: { errors },
    trigger,
    watch,
    setValue,
  } = useForm<MessageFormValues>({
    mode: 'onChange',
    defaultValues: {
      content: '',
    },
  })

  useEffect(() => {
    const validate = async () => {
      const result = await trigger(formFieldName)
      setValid(result)
    }
    ;(async () => {
      await validate() // initial validation
      const subscription = watch(() => {
        validate()
      })
      return () => subscription.unsubscribe()
    })()
  }, [trigger, watch])

  const onSubmit: SubmitHandler<MessageFormValues> = async (data: MessageFormValues) => {
    try {
      sendWsMessage(data.content)
      const payload: IMessageReq = { content: data.content }
      await postData(payload)
      setValue(formFieldName, '')
    } catch (error) {
      console.error('Failed to post message', error)
    }
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <div className='bg-gray-700 flex-none pb-5 px-4'>
        <div className='flex rounded-lg overflow-hidden'>
          <input
            {...register(formFieldName, {
              required: true,
              validate: (value: string): boolean => value.trim() !== '',
            })}
            type='text'
            maxLength={50}
            className='w-full px-4 bg-gray-600 text-white'
          />
          <button
            type='submit'
            disabled={!!errors.content}
            className='border-l border-gray-600 w-[4rem] flex flex-row items-center justify-center text-3xl text-grey border-r-4 border-gray-600 bg-gray-600 p-2'>
            {valid && !errors.content ? (
              <Image
                src='/assets/icon/message/send_enabled.svg'
                height={25}
                width={25}
                alt='message_send_enabled_icon'
              />
            ) : (
              <Image
                src='/assets/icon/message/send_disabled.svg'
                height={25}
                width={25}
                alt='message_send_disabled_icon'
              />
            )}
          </button>
        </div>
      </div>
    </form>
  )
}

export default MessageInput
