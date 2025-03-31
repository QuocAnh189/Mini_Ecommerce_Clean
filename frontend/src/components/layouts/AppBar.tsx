//hooks
import { useNavigate, useLocation } from 'react-router-dom'
import { useEffect, useState } from 'react'
import { useDebounce } from '@hooks/useDebounce'

//icons
import { IoCartOutline } from 'react-icons/io5'

//image
import logo from '@assets/images/logo.png'

//components
import ProductModalCreate from '@components/ProductModalCreate'

const AppBar = () => {
  const navigate = useNavigate()
  const location = useLocation()

  const [search, setSearch] = useState('')
  const debouncedSearchTerm = useDebounce(search, 500)

  useEffect(() => {
    if (debouncedSearchTerm !== '' || location.pathname === '/product') {
      navigate('/product', { state: { search: debouncedSearchTerm } })
    }
  }, [debouncedSearchTerm])

  return (
    <div className="flex bg-white text-black items-center justify-between px-[150px] py-3">
      <a href="/" className="flex items-center gap-2 text-xl font-bold">
        <img src={logo} height={0} width={0} alt="" className="w-8 h-8" />
        Ecommerce Clean
      </a>

      <div className="flex items-center w-[600px] rounded-md border border-gray-400 overflow-hidden">
        <div className="p-3 text-gray-600">🔍</div>
        <input
          type="text"
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          placeholder="What do you want to find today?..."
          className="flex-1 border-none p-2 outline-none"
        />
      </div>

      <button
        className="btn btn-info"
        onClick={() => (document?.getElementById('create_product_modal') as HTMLDialogElement).showModal()}
      >
        <p className="text-white">Add Product</p>
      </button>

      <div className="flex items-center gap-2">
        <button className="flex items-center gap-2">
          <div className="relative">
            <IoCartOutline size="32px" />
            <div className="bg-primary w-4 h-4 rounded-full absolute top-0 right-0">
              <p className="text-sm text-white bg-green-600 rounded-full">2</p>
            </div>
          </div>
          <a
            href="/cart"
            className={`flex flex-col items-start hover:underline ${
              location.pathname === '/cart' && 'text-green-600 font-bold'
            }`}
          >
            <p className="text-sm font-bold">Shopping Cart</p>
          </a>
        </button>
      </div>

      <ProductModalCreate />
    </div>
  )
}

export default AppBar
