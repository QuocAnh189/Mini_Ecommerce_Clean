//hooks
import { useEffect, useState } from 'react'
import { usePagination } from '@hooks/usePagination'

//components
import ProtectedLayout from '@layouts/protected'
import Pagination from '@components/Pagination'
import UserItem from '@components/UserItem'
import Skeleton from '@components/Skeleton'

//interfaces
import { IListUserRequest, IUser } from '@interfaces/user'
import { IPagination } from '@interfaces/common'

//redux
import { useGetListUsersQuery } from '@redux/services/user'

const initParams: IListUserRequest = {
  search: '',
  page: 1,
  size: 10,
  order_by: '',
  order_desc: false,
  take_all: false,
}

const ListUsers = () => {
  const [params, setParams] = useState<IListUserRequest>(initParams)
  const { data: users, isLoading } = useGetListUsersQuery(params)
  const pagination: IPagination = usePagination(users?.metadata.total_count, initParams.size)

  useEffect(() => {
    setParams((prev) => ({
      ...prev,
      page: pagination.currentPage,
    }))
  }, [pagination.currentPage])

  if (isLoading) return <Skeleton />

  return (
    <ProtectedLayout>
      <div className="h-screen px-40 mt-20">
        <div className="overflow-x-auto">
          <table className="table">
            <thead>
              <tr>
                <th>Name</th>
                <th>Email</th>
                <th>Status</th>
              </tr>
            </thead>
            <tbody>
              {users &&
                users.items.map((user: IUser, index: number) => <UserItem key={`index-${index}`} user={user} />)}
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

export default ListUsers
