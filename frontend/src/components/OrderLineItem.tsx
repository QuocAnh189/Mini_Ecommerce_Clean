//interfaces
import { IOrderLine } from '@interfaces/order'

//utils
import formatNumber from '@utils/formatNumber'

interface IProps {
  orderLine: IOrderLine
}

const OrderLineItem = (props: IProps) => {
  const { orderLine } = props
  return (
    <tr>
      <td>
        <div className="flex items-center gap-3">
          <div className="avatar">
            <div className="mask mask-squircle h-12 w-12">
              <img src={orderLine.product.image_url} alt="Avatar Tailwind CSS Component" />
            </div>
          </div>
          <div>
            <div className="font-bold">{orderLine.product.name}</div>
          </div>
        </div>
      </td>
      <td>{orderLine.product.description}</td>
      <td>{orderLine.quantity}</td>
      <td>{formatNumber(orderLine.price)} VND</td>
    </tr>
  )
}

export default OrderLineItem
