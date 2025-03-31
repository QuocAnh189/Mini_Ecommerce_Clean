import { EStatusOrder } from '@constants/enum'

export interface IOrder {
  id: string
  code: string
  lines: IOrderLine[]
  total_price: number
  status: EStatusOrder
  updated_at: string
}

export interface IOrderLine {
  product: IProductOrder
  quantity: number
  price: number
}

export interface IProductOrder {
  id: string
  code: string
  name: string
  image_url: string
  description: string
  price: number
}

export interface IListOrderRequest {
  user_id: string
  code: string
  status: string
  page: number
  size: number
  order_by: string
  order_desc: boolean
}

export interface IPlaceOrderRequest {
  user_id: string
  lines: IPlaceOrderLineRequest[]
}

export interface IPlaceOrderLineRequest {
  product_id: string
  quantity: number
}
