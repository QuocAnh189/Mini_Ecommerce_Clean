//hooks
import { useEffect } from 'react'
import { useNavigate } from 'react-router-dom'

//interfaces
import { IUser } from '@interfaces/user'

//redux
import { useAppDispatch } from '@redux/hook'
import { useGetCartQuery } from '@redux/services/cart'
import { setCart } from '@redux/slices/cart'

//icons
import { IoCartOutline } from 'react-icons/io5'

const CartIcon = () => {
  const dispatch = useAppDispatch()
  const navigate = useNavigate()
  const user: IUser = JSON.parse(localStorage.getItem('user') || '{}')
  const { data: cart } = useGetCartQuery(user?.id!)

  useEffect(() => {
    if (cart) {
      dispatch(setCart(cart.id))
    }
  }, [cart])

  const navigateCartScreen = () => {
    navigate('/cart', { state: { lines: cart?.lines } })
  }

  return (
    <div className="flex items-center gap-2">
      <div className="flex items-center gap-2 hover:cursor-pointer" onClick={navigateCartScreen}>
        <div className="relative">
          <IoCartOutline size="32px" />
          {cart && (
            <div className="bg-primary w-4 h-4 rounded-full absolute top-0 right-0">
              <p className="flex items-center justify-center text-sm text-white bg-green-600 rounded-full">
                {cart.lines.length}
              </p>
            </div>
          )}
        </div>
        <div
          className={`flex flex-col items-start hover:underline ${
            location.pathname === '/cart' && 'text-green-600 font-bold'
          }`}
        >
          <p className="text-sm font-bold">Shopping Cart</p>
        </div>
      </div>
    </div>
  )
}

export default CartIcon
