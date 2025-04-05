//utils
import formatNumber from '@utils/formatNumber'

//icon
import { GrSubtract } from 'react-icons/gr'
import { IoAdd } from 'react-icons/io5'
import { BiTrash } from 'react-icons/bi'

//interfaces
import { ICartLine } from '@interfaces/cart'

interface IProps {
  cartLine: ICartLine
  onRemove: (productId: string) => void
  onUpdate: (cartLineID: string, productId: string, quantity: number) => void
  onChecked: (cartLine: ICartLine) => void
  isSelected: boolean
}

const CartItem = (props: IProps) => {
  const { cartLine, onRemove, onUpdate, onChecked, isSelected } = props

  return (
    <div className="px-4">
      <input onChange={() => onChecked(cartLine)} type="checkbox" checked={isSelected} className="checkbox" />
      <div className="flex items-center py-2 relative text-sm">
        <div className="flex flex-1 items-center gap-3">
          <img src={cartLine.product.image_url} alt="" className="w-20 h-20 object-cover rounded-2xl" />
          <p className="font-semibold">{cartLine.product.name}</p>
        </div>
        <p className="w-[200px] font-semibold text-center">{formatNumber(cartLine.product.price)} VND</p>
        <div className="w-[100px] flex p-1  items-center justify-center gap-2 border border-solid border-gray300 rounded-3xl">
          <button
            disabled={cartLine.quantity === 1}
            className={`flex items-center justify-center p-2 rounded-full bg-gray-300 ${
              cartLine.quantity > 1 ? 'hover:cursor-pointer hover:bg-gray-400 ' : ''
            }`}
            onClick={() => onUpdate(cartLine.id, cartLine.product.id, cartLine.quantity - 1)}
          >
            <GrSubtract />
          </button>
          <p>{cartLine.quantity}</p>
          <button
            onClick={() => onUpdate(cartLine.id, cartLine.product.id, cartLine.quantity + 1)}
            className="flex items-center justify-center p-2 rounded-full bg-gray-300 hover:bg-gray-400 hover:cursor-pointer"
          >
            <IoAdd />
          </button>
        </div>
        <p className="w-[200px] font-semibold text-center">
          {formatNumber(cartLine.product.price * cartLine.quantity)} VND
        </p>
        <button
          onClick={() => onRemove(cartLine.product.id)}
          className="absolute right-0 p-2 rounded-full bg-gray-300 hover:bg-gray-400 hover:cursor-pointer"
        >
          <BiTrash className="text-xl" />
        </button>
      </div>
      <div className="divider"></div>
    </div>
  )
}

export default CartItem
