//utils
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import utc from 'dayjs/plugin/utc'
import timezone from 'dayjs/plugin/timezone'
import duration from 'dayjs/plugin/duration'

dayjs.extend(utc)
dayjs.extend(timezone)
dayjs.extend(relativeTime)
dayjs.extend(duration)

function formatDate(datetime: string, format?: string, fromNow?: boolean): string {
  const utcDatetime = dayjs?.tz(datetime)?.toISOString()
  const defaultFormat = 'DD/MM/YYYY hh:mm A'

  if (fromNow) {
    return dayjs(utcDatetime).fromNow()
  }

  return dayjs(utcDatetime).format(format || defaultFormat)
}

export default formatDate
