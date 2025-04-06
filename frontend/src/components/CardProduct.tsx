//components
import toast from 'react-hot-toast'
import Loading from './Loading'
import ProductModalUpdate from './ProductModalUpdate'

//interfaces
import { IProduct } from '@interfaces/product'
import { IAddProductRequest } from '@interfaces/cart'
import { IUser } from '@interfaces/user'

//constants
import { ERole } from '@constants/enum'

//util
import formatNumber from '@utils/formatNumber'

//redux
import { useAddProductToCartMutation } from '@redux/services/cart'
import { useDeleteProductMutation } from '@redux/services/product'
import { useAppSelector } from '@redux/hook'

//icon
import { SlHandbag } from 'react-icons/sl'

interface IProps {
  product: IProduct
}

const CardProduct = (props: IProps) => {
  const { product } = props

  const user: IUser = JSON.parse(localStorage.getItem('user')!)
  const cartId = useAppSelector((state) => state.cart.cartId)

  const [AddProductToCart, { isLoading: loadingAdd }] = useAddProductToCartMutation()
  const [DeleteProduct, { isLoading: loadingDelete }] = useDeleteProductMutation()

  const handleAddProductToCart = async () => {
    const data: IAddProductRequest = {
      cart_id: cartId!,
      product_id: product.id,
      quantity: 1,
    }

    try {
      const result = await AddProductToCart({ userId: user.id, data }).unwrap()
      if (result) {
        toast.success('Add product to cart successfully.')
      }
    } catch {
      toast.error('Something went wrong.')
    }
  }

  const handleDeleteProduct = async () => {
    try {
      const result = await DeleteProduct(product.id).unwrap()
      if (result) {
        toast.success('Delete product successfully.')
      }
    } catch {
      toast.error('Something went wrong.')
    }
  }

  return (
    <div className="h-[400px] border border-gray-300 rounded-md hover:cursor-pointer">
      <img src={product.image_url} alt="" className="w-full h-3/5 z-[1] object-cover" />
      <div className="p-4 h-2/5 flex flex-col items-center justify-between">
        <div className="flex items-center justify-between w-full">
          <div>
            <p className="text-black font-bold">{product.name}</p>
            <p className="text-gray-500 font-medium text-ellipsis">{product.description}</p>
            <p className="text-black">{formatNumber(product.price)} VND</p>
          </div>
          <div
            onClick={handleAddProductToCart}
            className="bg-gray-300 p-3 rounded-full hover:bg-gray-400 flex items-center gap-2"
          >
            {loadingAdd ? <Loading /> : <SlHandbag />}
          </div>
        </div>
        {user.role === ERole.ADMIN && (
          <div className="w-full flex items-center justify-between">
            <button
              onClick={() =>
                (document?.getElementById(`update_product_modal_${product.id}`) as HTMLDialogElement).showModal()
              }
              className="btn btn-info"
            >
              Update
            </button>
            <button onClick={handleDeleteProduct} className="btn btn-error">
              {loadingDelete ? <Loading /> : 'Delete'}
            </button>
          </div>
        )}
      </div>
      {user.role === ERole.ADMIN && (
        <ProductModalUpdate
          product={{
            id: product.id,
            name: product.name,
            description: product.description,
            image: null,
            price: product.price,
          }}
          image_url={product.image_url}
        />
      )}
    </div>
  )
}

export default CardProduct
