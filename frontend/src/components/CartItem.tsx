//icon
import { GrSubtract } from 'react-icons/gr'
import { IoAdd } from 'react-icons/io5'

const CartItem = () => {
  return (
    <div className="px-4">
      <div className="flex items-center py-2 relative text-sm">
        <div className="flex flex-1 items-center gap-3">
          <img
            src="https://res.cloudinary.com/dadvtny30/image/upload/v1708061747/organicfood/product/bto2petcvvdqawblw8cj.png"
            alt=""
            className="w-20 h-20 object-cover"
          />
          <p className="font-semibold">Green Capsicum</p>
        </div>
        <p className="w-[200px] font-semibold text-center">$14.00</p>
        <div className="w-[100px] flex p-1  items-center justify-center gap-2 border border-solid border-gray300 rounded-3xl">
          <button className="flex items-center justify-center p-2 rounded-full bg-gray300">
            <GrSubtract />
          </button>
          <p>0</p>
          <button className="flex items-center justify-center p-2 rounded-full bg-gray300">
            <IoAdd />
          </button>
        </div>
        <p className="w-[200px] font-semibold text-center">$70.00</p>
      </div>
      <div className="divider"></div>
    </div>
  )
}

export default CartItem
