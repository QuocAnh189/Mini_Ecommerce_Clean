//assets
import logo from '@assets/images/logo.png'
import applePay from '@assets/images/applePay.png'
import visaPay from '@assets/images/visa.png'
import discoverPay from '@assets/images/discover.png'
import masterCart from '@assets/images/masterCart.png'
import secureCart from '@assets/images/secureCart.png'

const Footer = () => {
  return (
    <div>
      <div className="flex bg-[#333333] items-center justify-between px-[150px] py-10">
        <div className="flex flex-col items-start gap-3">
          <div className="flex items-center gap-2">
            <img src={logo} height={0} width={0} alt="" className="w-8 h-8" />
            <p className="text-white font-bold text-2xl">Ecommerce Clean</p>
          </div>
          <p className="text-gray-500 text-sm max-w-80">
            Morbi cursus porttitor enim lobortis molestie. Duis gravida turpis dui, eget bibendum magna congue nec.
          </p>
          <div className="flex items-center gap-3">
            <button>
              <p className="text-white text-sm font-semibold">(219) 555-014</p>
              <div className="w-full bg-primary h-[2px]" />
            </button>
            <span className="text-gray500">or</span>
            <button>
              <p className="text-white  text-sm font-semibold">21521360@gm.uit.edu.vn</p>
              <div className="w-full bg-primary h-[2px]" />
            </button>
          </div>
        </div>
        <div className="flex items-center gap-20">
          <div className="flex text-gray-500 flex-col items-start text-gray500 gap-1">
            <h1 className="pb-2 text-white font-bold">My Account</h1>
            <p className="hover:underline hover:cursor-pointer">My Account</p>
            <p className="hover:underline hover:cursor-pointer">Order History</p>
            <p className="hover:underline hover:cursor-pointer">Shopping Cart</p>
            <p className="hover:underline hover:cursor-pointer">Wishlist</p>
          </div>
          <div className="flex text-gray-500 flex-col items-start text-gray500 gap-1">
            <h1 className="text-white font-bold pb-2">Helps</h1>
            <p className="hover:underline hover:cursor-pointer">Contact</p>
            <p className="hover:underline hover:cursor-pointer">Faqs</p>
            <p className="hover:underline hover:cursor-pointer">Terms & Condition</p>
            <p className="hover:underline hover:cursor-pointer">Privacy Policy</p>
          </div>
          <div className="flex text-gray-500 flex-col items-start text-gray500 gap-1">
            <h1 className="text-white font-bold pb-2">Proxy</h1>
            <p className="hover:underline hover:cursor-pointer">About</p>
            <p className="hover:underline hover:cursor-pointer">Shop</p>
            <p className="hover:underline hover:cursor-pointer">Product</p>
            <p className="hover:underline hover:cursor-pointer">Track Order</p>
          </div>
          <div className="flex text-gray-500 flex-col items-start text-gray500 gap-1">
            <h1 className="text-white font-bold pb-2">Categories</h1>
            <p className="hover:underline hover:cursor-pointer">Fruit & Vegetables</p>
            <p className="hover:underline hover:cursor-pointer">Meat & Fish</p>
            <p className="hover:underline hover:cursor-pointer">Break & Bakery</p>
            <p className="hover:underline hover:cursor-pointer">Beauty & Health</p>
          </div>
        </div>
      </div>

      <div className="flex bg-[#333333] items-center justify-between px-[150px] py-3 gap-3">
        <p className="text-gray-500 text-sm">Clean eCommerce © 2025. All Rights Reserved</p>
        <div className="flex items-center gap-2">
          <img src={applePay} className="w-14 h-8 object-contain" alt="" />
          <img src={visaPay} className="w-14 h-8 object-contain" alt="" />
          <img src={discoverPay} className="w-14 h-8 object-contain" alt="" />
          <img src={masterCart} className="w-14 h-8 object-contain" alt="" />
          <img src={secureCart} className="w-14 h-8 object-contain" alt="" />
        </div>
      </div>
    </div>
  )
}

export default Footer
