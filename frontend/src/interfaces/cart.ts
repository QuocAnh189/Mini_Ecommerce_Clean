export interface ICart {
  id: string
  user: {
    id: string
    email: string
  }
  lines: ICartLine[]
}

export interface ICartLine {
  id: string
  product: IProductCart
  quantity: number
  price: number
}

export interface IProductCart {
  id: string
  code: string
  name: string
  image_url: string
  description: string
  price: number
}

export interface IAddProductRequest {
  cart_id: string
  product_id: string
  quantity: number
}

export interface IUpdateCartLineRequest {
  id: string
  cart_id: string
  product_id: string
  quantity: number
}

export interface IRemoveProductRequest {
  cart_id: string
  product_id: string
}
