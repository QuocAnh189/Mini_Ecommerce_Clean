//interface
import { IPagination } from '@interfaces/common'

//icon
import { GrFormNext } from 'react-icons/gr'
import { GrPrevious } from 'react-icons/gr'

interface Props {
  pagination: IPagination
}

const Pagination = (props: Props) => {
  const { pagination } = props

  const groupLeftPage = [...Array(pagination.maxPage)]
    .map((_, index) => index + 1)
    .slice(
      pagination.currentPage >= 5 ? pagination.currentPage - 5 : 0,
      pagination.currentPage >= 5 ? pagination.currentPage : 5,
    )

  const groupRightPage = [...Array(pagination.maxPage)]
    .map((_, index) => index + 1)
    .slice(pagination.maxPage - 5, pagination.maxPage)

  return (
    <div className="join">
      {pagination.currentPage > 1 && (
        <button onClick={pagination.prevPage} className="join-item btn">
          <GrPrevious />
        </button>
      )}

      {pagination.maxPage <= 10 ? (
        [...Array(pagination.maxPage)].map((_, i: number) => {
          return (
            <button
              className="join-item btn"
              key={`{page}-${i + 1}`}
              onClick={() => pagination.goToPage(i + 1)}
              disabled={pagination.currentPage === i + 1}
              aria-label={`Page ${i + 1}`}
            >
              {i + 1}
            </button>
          )
        })
      ) : (
        <div className="flex flex-row gap-2">
          {(groupLeftPage[groupLeftPage.length - 1] < groupRightPage[0] || groupLeftPage.length === 0) &&
            groupLeftPage.map((page: number, index: number) => {
              return (
                <button
                  className="join-item btn"
                  key={`page-${index}`}
                  onClick={() => pagination.goToPage(page)}
                  disabled={pagination.currentPage === page}
                  aria-label={`Page ${page}`}
                >
                  {page}
                </button>
              )
            })}

          <button className="join-item btn btn-disabled">...</button>

          {groupRightPage.map((page: number, index: number) => {
            return (
              <button
                className="join-item btn"
                key={`page-${index}`}
                onClick={() => pagination.goToPage(page)}
                disabled={pagination.currentPage === page}
                aria-label={`Page ${page}`}
              >
                {page}
              </button>
            )
          })}
        </div>
      )}
      {pagination.currentPage < pagination.maxPage && (
        <button
          className="join-item btn"
          onClick={pagination.nextPage}
          disabled={pagination.currentPage === pagination.maxPage}
        >
          <GrFormNext />
        </button>
      )}
    </div>
  )
}

export default Pagination
