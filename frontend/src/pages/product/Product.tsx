//hooks
import { useLocation } from 'react-router-dom'
import { useEffect, useState } from 'react'
import { usePagination } from 'hooks/usePagination'

//components
import ProtectedLayout from '@layouts/protected'
import CardProduct from '@components/CardProduct'
import Skeleton from '@components/Skeleton'
import Pagination from '@components/Pagination'

//interfaces
import { IListProductRequest, IProduct } from '@interfaces/product'
import { IPagination } from '@interfaces/common'

//redux
import { useGetListProductsQuery } from '@redux/services/product'

const initParams: IListProductRequest = {
  search: '',
  page: 1,
  size: 8,
  order_by: '',
  order_desc: false,
  take_all: false,
}

const Product = () => {
  const location = useLocation()
  const [params, setParams] = useState<IListProductRequest>({ ...initParams, search: location.state?.search || '' })
  const { data: products, isLoading } = useGetListProductsQuery(params)
  const pagination: IPagination = usePagination(products?.metadata.total_count, initParams.size)

  useEffect(() => {
    setParams((prev) => ({
      ...prev,
      page: pagination.currentPage,
      search: location.state?.search || '',
    }))
  }, [pagination.currentPage, location.state?.search])

  if (isLoading) return <Skeleton />

  return (
    <ProtectedLayout>
      <div className="min-h-screen flex flex-col items-center bg-gray-100 py-10">
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 justify-start px-10 py-5">
          {products &&
            products.items.map((product: IProduct, index: number) => (
              <CardProduct key={`product-${index}`} product={product} />
            ))}
        </div>
        {pagination.maxPage > 1 && <Pagination pagination={pagination} />}
      </div>
    </ProtectedLayout>
  )
}

export default Product
