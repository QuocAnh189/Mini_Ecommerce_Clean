//components
import toast from 'react-hot-toast'

//icon
import { SlHandbag } from 'react-icons/sl'

//interfaces
import { IProduct } from '@interfaces/product'
import { IAddProductRequest } from '@interfaces/cart'

//util
import formatNumber from '@utils/formatNumber'

//redux
import { useAddProductToCartMutation } from '@redux/services/cart'
import { useAppSelector } from '@redux/hook'

interface IProps {
  product: IProduct
}

const CardProduct = (props: IProps) => {
  const { product } = props

  const userId = JSON.parse(localStorage.getItem('user')!).id
  const cartId = useAppSelector((state) => state.cart.cartId)
  const [AddProductToCart] = useAddProductToCartMutation()

  const handleAddProductToCart = async () => {
    const data: IAddProductRequest = {
      cart_id: cartId!,
      product_id: product.id,
      quantity: 1,
    }

    try {
      const result = await AddProductToCart({ userId, data }).unwrap()
      if (result) {
        toast.success('Add product to cart successfully.')
      }
    } catch {
      toast.error('Something went wrong.')
    }
  }

  return (
    <div className="h-[400px] border border-gray-300 rounded-md hover:cursor-pointer z-10">
      <img src={product.image_url} alt="" className="w-full h-3/5 z-[1] object-cover" />
      <div className="p-4 h-2/5 flex items-center justify-between">
        <div>
          <p className="text-black font-bold">{product.name}</p>
          <p className="text-gray-500 font-medium text-ellipsis">{product.description}</p>
          <p className="text-black">{formatNumber(product.price)} VND</p>
        </div>
        <div
          onClick={handleAddProductToCart}
          className="bg-gray-300 p-3 rounded-full hover:bg-gray-400 flex items-center gap-2"
        >
          +<SlHandbag />
        </div>
      </div>
    </div>
  )
}

export default CardProduct
