//hooks
import { useEffect, useState } from 'react'
import { useNavigate, useLocation } from 'react-router-dom'
import { useDebounce } from '@hooks/useDebounce'

//components
import ProductModalCreate from '@components/ProductModalCreate'
import CartIcon from '@components/CartIcon'

//interfaces
import { IUser } from '@interfaces/user'

//constants
import { ERole } from '@constants/enum'

//assets
import logo from '@assets/images/logo.png'

const AppBar = () => {
  const navigate = useNavigate()
  const location = useLocation()

  const user: IUser = JSON.parse(localStorage.getItem('user') || '{}')

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

      {user.role === ERole.ADMIN && (
        <button
          className="btn btn-info"
          onClick={() => (document?.getElementById('create_product_modal') as HTMLDialogElement).showModal()}
        >
          <p className="text-white">Add Product</p>
        </button>
      )}

      <CartIcon />

      <ProductModalCreate />
    </div>
  )
}

export default AppBar
