//hooks
import { useState, useEffect } from 'react'
import { usePagination } from '@hooks/usePagination'

//components
import OrderItem from '@components/OrderItem'
import Pagination from '@components/Pagination'
import Skeleton from '@components/Skeleton'

//interfaces
import { IPagination } from '@interfaces/common'
import { IListOrderRequest, IOrder } from '@interfaces/order'

//redux
import { useGetListMyOrdersQuery } from '@redux/services/order'

const initParams: IListOrderRequest = {
  user_id: JSON.parse(localStorage.getItem('user') || '{}').id,
  code: '',
  status: '',
  page: 1,
  size: 10,
  order_by: '',
  order_desc: false,
}

const Order = () => {
  const [params, setParams] = useState<IListOrderRequest>(initParams)
  const { data: orders, isLoading } = useGetListMyOrdersQuery(params)
  const pagination: IPagination = usePagination(orders?.metadata.total_count, initParams.size)

  useEffect(() => {
    setParams((prev) => ({
      ...prev,
      page: pagination.currentPage,
    }))
  }, [pagination.currentPage])

  if (isLoading) return <Skeleton />

  return (
    <div className="h-screen flex flex-col items-center mt-20 gap-4">
      <div className="overflow-x-auto px-40">
        <table className="table table-zebra">
          <thead>
            <tr>
              <th>ID</th>
              <th>Code</th>
              <th>Total Price</th>
              <th>Status</th>
              <th>Created</th>
            </tr>
          </thead>
          <tbody>
            {orders &&
              orders.items.map((order: IOrder, index: number) => <OrderItem key={`order-${index}`} order={order} />)}
          </tbody>
        </table>
      </div>
      {pagination.maxPage > 1 && <Pagination pagination={pagination} />}
    </div>
  )
}

export default Order
