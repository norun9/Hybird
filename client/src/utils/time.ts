export const getFormattedCurrentTime = () => {
  const now = new Date()
  const formattedDate = `${now.getFullYear()}年${getMonth(now)}月${getDate(now)}日`
  let hours = now.getHours()
  const minutes = String(now.getMinutes()).padStart(2, '0')
  const ampm = hours >= 12 ? 'PM' : 'AM'
  hours = hours % 12
  hours = hours ? hours : 12 // 0 should be 12
  const formattedTime = `${String(hours).padStart(2, '0')}:${minutes} ${ampm}`
  return `${formattedDate} ${formattedTime}`
}

const getMonth = (date: Date): string => {
  return String(date.getMonth() + 1).padStart(2, '0')
}

const getDate = (date: Date): string => {
  return String(date.getDate()).padStart(2, '0')
}
