//hook
import { useNavigate } from 'react-router-dom'

//components
import UpdateStatusModel from './UpdateStatusModel'

//util
import formatDate from '@utils/formatDate'
import formatNumber from '@utils/formatNumber'

//interfaces
import { IOrder } from '@interfaces/order'

//enum
import { EStatusOrder } from '@constants/enum'

interface IProps {
  order: IOrder
}

const OrderItem = (props: IProps) => {
  const { order } = props
  const navigate = useNavigate()

  const navigateToOrderDetails = () => {
    navigate(`/order/${order.id}`, { state: { items: order.lines } })
  }

  return (
    <>
      <tr>
        <th>{order.id}</th>
        <td>{order.code}</td>
        <td>{formatNumber(order.total_price)} VND</td>
        <td>
          {order.status === EStatusOrder.NEW && (
            <div className="flex items-center gap-2">
              <div aria-label="new" className="status status-info"></div>
              {order.status}
            </div>
          )}
          {order.status === EStatusOrder.PROGRESS && (
            <div className="flex items-center gap-2">
              <div aria-label="progress" className="status status-warning"></div>
              {order.status}
            </div>
          )}
          {order.status === EStatusOrder.DONE && (
            <div className="flex items-center gap-2">
              <div aria-label="done" className="status status-success"></div>
              {order.status}
            </div>
          )}
          {order.status === EStatusOrder.CANCELED && (
            <div className="flex items-center gap-2">
              <div aria-label="canceled" className="status status-error"></div>
              {order.status}
            </div>
          )}
        </td>
        <td>{formatDate(order.updated_at)}</td>
        <th>
          <button onClick={navigateToOrderDetails} className="btn btn-ghost btn-md">
            details
          </button>
        </th>
        <th>
          <button
            onClick={() =>
              (document?.getElementById(`update_status_modal_${order.id}`) as HTMLDialogElement).showModal()
            }
            className="btn btn-ghost btn-md"
          >
            change status
          </button>
        </th>
      </tr>
      <UpdateStatusModel orderId={order.id} currentStatus={order.status} />
    </>
  )
}

export default OrderItem
