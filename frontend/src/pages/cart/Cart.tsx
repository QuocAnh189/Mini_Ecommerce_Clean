import CartItem from '@components/CartItem'

const Cart = () => {
  return (
    <div className="w-full mt-10">
      <div className="py-8 flex flex-col items-center px-[150px] gap-8">
        <div className="w-full flex gap-4">
          <table className="table">
            <div className="h-[50px]">
              <div className="flex items-center px-4 py-2 text-sm">
                <p className="flex-1 text-gray500">PRODUCT</p>
                <p className="w-[200px] text-gray-500 text-center">PRICE</p>
                <p className="w-[100px] text-gray-500 text-center">QUANTITY</p>
                <p className="w-[200px] text-gray-500 text-center">PRICE</p>
              </div>
              <div className="h-[350px] overflow-hidden">
                <div className="w-full h-full overflow-y-scroll pl-[17px] box-content">
                  <CartItem />
                  <CartItem />
                  <CartItem />
                  <CartItem />
                </div>
              </div>
            </div>
          </table>
          <div className="w-1/3 h-[400px] flex flex-col justify-between border-[2px] border-solid border-gray-300 space-y-2 rounded-lg p-4">
            <div>
              <h2 className="text-2xl py-2">Cart Total</h2>
              <div className="flex items-center justify-between py-2 text-sm">
                <p className="text-gary500">Subtotal:</p>
                <p className="text-black">$84.00</p>
              </div>
              <div className="divider"></div>
              <div className="flex items-center justify-between py-2 text-sm">
                <p className="text-gary500">Shipping:</p>
                <p className="text-black">Free</p>
              </div>
              <div className="divider"></div>
              <div className="flex items-center justify-between py-2">
                <p className="text-gary500">Total:</p>
                <p className="text-black font-medium">$84.00</p>
              </div>
            </div>

            <a
              href="/cart/checkout"
              className="flex justify-center text-white bg-green-600 hover:bg-green-500 rounded-3xl font-semibold py-3 hover:bg-subprimary"
            >
              Order
            </a>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Cart
