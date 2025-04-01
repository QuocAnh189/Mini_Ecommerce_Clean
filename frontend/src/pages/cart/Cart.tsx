//hooks
import { useEffect, useState } from 'react'
import { useLocation } from 'react-router-dom'

//components
import ProtectedLayout from '@layouts/protected'
import CartItem from '@components/CartItem'
import Loading from '@components/Loading'
import toast from 'react-hot-toast'

//redux
import { useAppSelector } from '@redux/hook'
import { useRemoveProductFromCartMutation, useUpdateCartLineMutation } from '@redux/services/cart'
import { usePlaceOrderMutation } from '@redux/services/order'

//interfaces
import { ICartLine, IRemoveProductRequest, IUpdateCartLineRequest } from '@interfaces/cart'
import { IPlaceOrderLineRequest } from '@interfaces/order'

//utils
import formatNumber from '@utils/formatNumber'

const Cart = () => {
  const location = useLocation()

  const [cartLines, setCartLines] = useState<ICartLine[]>(location.state?.lines || [])
  const [cartLinesChecked, setCartLinesChecked] = useState<ICartLine[]>([])
  const [totalPrice, setTotalPrice] = useState<number>(0)

  const userId = JSON.parse(localStorage.getItem('user')!)?.id
  const cartId = useAppSelector((state) => state.cart.cartId)

  const [RemoveProduct] = useRemoveProductFromCartMutation()
  const [UpdateCartLine] = useUpdateCartLineMutation()
  const [PlaceOrder, { isLoading: loadingOrder }] = usePlaceOrderMutation()

  const handleRemoveProductFromCart = async (productId: string) => {
    const data: IRemoveProductRequest = {
      cart_id: cartId!,
      product_id: productId,
    }

    try {
      const result = await RemoveProduct({ userId, data }).unwrap()
      if (result) {
        setCartLines((prev: ICartLine[]) => prev.filter((line: ICartLine) => line.product.id !== productId))
        toast.success('Remove product from cart successfully.')
      }
    } catch {
      toast.error('Something went wrong.')
    }
  }

  const handleUpdateCartLine = async (cartLineID: string, productId: string, quantity: number) => {
    const data: IUpdateCartLineRequest = {
      id: cartLineID,
      cart_id: cartId!,
      product_id: productId,
      quantity,
    }

    try {
      const result = await UpdateCartLine({ userId, data }).unwrap()
      if (result) {
        setCartLines((prev: ICartLine[]) =>
          prev.map((line: ICartLine) => {
            if (line.product.id === productId) {
              return { ...line, quantity, price: line.product.price * quantity }
            }
            return line
          }),
        )
        toast.success('Update successfully.')
      }
    } catch {
      toast.error('Something went wrong.')
    }
  }

  const handleSelectedLine = (cartLine: ICartLine) => {
    const isChecked = cartLinesChecked.some((line) => line.id === cartLine.id)
    if (isChecked) {
      setCartLinesChecked((prev) => prev.filter((line) => line.id !== cartLine.id))
    } else {
      setCartLinesChecked((prev) => [...prev, cartLine])
    }
  }

  const handlePlaceOrder = async () => {
    const lines: IPlaceOrderLineRequest[] = cartLinesChecked.map((line) => ({
      product_id: line.product.id,
      quantity: line.quantity,
    }))

    try {
      const result = await PlaceOrder({ user_id: userId, lines }).unwrap()
      if (result) {
        toast.success('Order successfully.')
        setCartLinesChecked([])
      }
    } catch {
      toast.error('Something went wrong.')
    }
  }

  useEffect(() => {
    const total = cartLinesChecked.reduce((acc: number, line: ICartLine) => {
      return acc + line.price
    }, 0)
    setTotalPrice(total)
  }, [cartLinesChecked.length])

  return (
    <ProtectedLayout>
      <div className="w-full h-screen mt-10">
        <div className="py-8 flex flex-col items-center px-[150px] gap-8">
          <div className="w-full flex gap-4">
            <div className="table">
              <div className="h-[50px]">
                <div className="flex items-center px-4 py-2 text-sm">
                  <p className="flex-1 text-gray500">PRODUCT</p>
                  <p className="w-[200px] text-gray-500 text-center">BASE PRICE</p>
                  <p className="w-[100px] text-gray-500 text-center">QUANTITY</p>
                  <p className="w-[200px] text-gray-500 text-center">PRICE</p>
                </div>
                <div className="h-[500px] overflow-hidden">
                  <div className="w-full h-full overflow-y-scroll pl-[17px] box-content">
                    {cartLines.length !== 0 &&
                      cartLines.map((line: ICartLine, index: number) => (
                        <CartItem
                          key={`cart-line-${index}`}
                          cartLine={line}
                          onRemove={handleRemoveProductFromCart}
                          onUpdate={handleUpdateCartLine}
                          onChecked={handleSelectedLine}
                          isSelected={cartLinesChecked.some((l) => l.id === line.id)}
                        />
                      ))}
                  </div>
                </div>
              </div>
            </div>
            <div className="w-1/3 h-[400px] flex flex-col justify-between border-[2px] border-solid border-gray-300 space-y-2 rounded-lg p-4">
              <div>
                <h2 className="text-2xl py-2">Cart Total</h2>
                <div className="flex items-center justify-between py-2 text-sm">
                  <p className="text-gary500">Subtotal:</p>
                  <p className="text-black">{formatNumber(totalPrice)} VND</p>
                </div>
                <div className="divider"></div>
                <div className="flex items-center justify-between py-2 text-sm">
                  <p className="text-gary500">Shipping:</p>
                  <p className="text-black">Free</p>
                </div>
                <div className="divider"></div>
                <div className="flex items-center justify-between py-2">
                  <p className="text-gary500">Total:</p>
                  <p className="text-black font-medium">{formatNumber(totalPrice)} VND</p>
                </div>
              </div>

              <button
                disabled={cartLinesChecked.length === 0 || loadingOrder}
                onClick={handlePlaceOrder}
                className="flex justify-center text-white bg-green-600 hover:bg-green-500 rounded-3xl font-semibold py-3 hover:bg-subprimary hover:cursor-pointer transition-all"
              >
                {loadingOrder ? <Loading /> : 'Place Order'}
              </button>
            </div>
          </div>
        </div>
      </div>
    </ProtectedLayout>
  )
}

export default Cart
