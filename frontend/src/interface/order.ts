export interface Order {
  id: string
  code: string
  lines: OrderLine[]
  total_price: number
  status: string
  updated_at: string
}

export interface OrderLine {
  product: IProductOrder
  quantity: number
  price: number
}

export interface IProductOrder {
  id: string
  code: string
  name: string
  description: string
  price: number
}

export interface IListOrderRequest {
  code: string
  status: boolean
  page: number
  size: number
  order_by: string
  order_desc: boolean
}

export interface PlaceOrderRequest {
  user_id: string
  lines: PlaceOrderLineRequest[]
}

export interface PlaceOrderLineRequest {
  product_id: string
  quantity: number
}
