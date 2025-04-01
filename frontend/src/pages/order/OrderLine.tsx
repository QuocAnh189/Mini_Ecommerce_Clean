//hooks
import { useLocation, useNavigate } from 'react-router-dom'
import { usePagination } from '@hooks/usePagination'

//components
import ProtectedLayout from '@layouts/protected'
import Pagination from '@components/Pagination'
import OrderLineItem from '@components/OrderLineItem'

//interfaces
import { IOrderLine } from '@interfaces/order'
import { IPagination } from '@interfaces/common'

const OrderLine = () => {
  const navigate = useNavigate()
  const location = useLocation()
  const { items }: { items: IOrderLine[] } = location.state

  const pagination: IPagination = usePagination(items.length, 4)

  return (
    <ProtectedLayout>
      <div className="h-screen px-40 mt-20">
        <button
          onClick={() => {
            navigate(-1)
          }}
          className="btn btn-outline btn-info"
        >
          Back
        </button>
        <div className="overflow-x-auto">
          <table className="table">
            <thead>
              <tr>
                <th>Product</th>
                <th>Description</th>
                <th>Quantity</th>
                <th>Price</th>
              </tr>
            </thead>
            <tbody>
              {items &&
                items.map((orderLine: IOrderLine, index: number) => (
                  <OrderLineItem key={`order-line-${index}`} orderLine={orderLine} />
                ))}
            </tbody>
          </table>
        </div>

        <div className="flex justify-center mt-10">
          {pagination.maxPage > 1 && <Pagination pagination={pagination} />}
        </div>
      </div>
    </ProtectedLayout>
  )
}

export default OrderLine
