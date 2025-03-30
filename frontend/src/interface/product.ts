export interface IProduct {
  id: string
  code: string
  name: string
  image_url: string
  description: string
  price: number
  active: boolean
  created_at: string
  updated_at: string
}

export interface IListProductRequest {
  search: string
  page: number
  size: number
  order_by: string
  order_desc: boolean
  take_all: boolean
}

export interface ICreateProductRequest {
  name: string
  description: string
  image: any
  price: number
}

export interface IUpdateProductRequest {
  name: string
  description: string
  image: any
  price: number
}
