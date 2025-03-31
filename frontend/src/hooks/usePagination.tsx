//hook
import { useState } from 'react'

export const usePagination = (totalCount: number = 0, itemsPerPage: number = 10) => {
  const [currentPage, setCurrentPage] = useState<number>(1)
  const maxPage: number = Math.ceil(totalCount / itemsPerPage)

  const nextPage = (): void => {
    setCurrentPage((currentPage: number) => Math.min(currentPage + 1, maxPage))
  }

  const prevPage = (): void => {
    setCurrentPage((currentPage: number) => Math.max(currentPage - 1, 1))
  }

  const goToPage = (page: number): void => {
    const pageNumber: number = Math.max(1, page)
    setCurrentPage(() => Math.min(pageNumber, maxPage))
  }

  return { nextPage, prevPage, goToPage, currentPage, setCurrentPage, maxPage }
}
