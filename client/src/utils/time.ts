export const getFormattedCurrentTime = () => {
  const now = new Date()
  const formattedDate = `${now.getFullYear()}年${getMonth(now)}月${getDate(now)}日`
  const formattedTime = now.toLocaleTimeString('en-US', {
    hour: 'numeric',
    minute: 'numeric',
    hour12: true,
  })
  return `${formattedDate} ${formattedTime}`
}

const getMonth = (date: Date): string => {
  return String(date.getMonth() + 1).padStart(2, '0')
}

const getDate = (date: Date): string => {
  return String(date.getDate()).padStart(2, '0')
}
