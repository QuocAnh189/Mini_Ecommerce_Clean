//icon
import { SlHandbag } from 'react-icons/sl'

const CardProduct = () => {
  return (
    <div className="border border-gray-300 rounded-md hover:cursor-pointer z-10">
      <img
        src="https://res.cloudinary.com/dadvtny30/image/upload/v1708061748/organicfood/product/cqodgw5bfg90smswkate.png"
        alt=""
        className="w-full z-[1] object-cover"
      />
      <div className="p-4 flex items-center justify-between">
        <div>
          <p className="text-gray500 font-medium">Green Apple</p>
          <p className="text-black">100000 VND</p>
        </div>
        <div className="bg-gray-300 p-3 rounded-full hover:bg-gray-400 flex items-center gap-2">
          +<SlHandbag />
        </div>
      </div>
    </div>
  )
}

export default CardProduct
