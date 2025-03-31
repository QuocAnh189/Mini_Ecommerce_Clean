import { useState, useEffect } from 'react'

/**
 * Custom hook: useDebounce
 * @param {any} value - Giá trị cần debounce
 * @param {number} delay - Độ trễ debounce (ms)
 * @returns {any} - Giá trị sau khi debounce
 */
export const useDebounce = (value: string, delay: number) => {
  const [debouncedValue, setDebouncedValue] = useState(value)

  useEffect(() => {
    const handler = setTimeout(() => {
      setDebouncedValue(value)
    }, delay)

    return () => {
      clearTimeout(handler)
    }
  }, [value, delay])

  return debouncedValue
}
