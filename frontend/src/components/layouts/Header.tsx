//hook
import { useLocation, useNavigate } from 'react-router-dom'

//component
import toast from 'react-hot-toast'
import Loading from '@components/Loading'

//interface
import { IUser } from '@interfaces/user'

//redux
import { useSignOutMutation } from '@redux/services/auth'
import { useAppDispatch } from '@redux/hook'
import { setAuth } from '@redux/slices/auth'

//icons
import { CiLocationOn } from 'react-icons/ci'

const Header = () => {
  const location = useLocation()
  const navigate = useNavigate()
  const dispatch = useAppDispatch()

  const user: IUser = JSON.parse(localStorage.getItem('user')!)
  const [SignOut, { isLoading }] = useSignOutMutation()

  const handleLogout = async () => {
    try {
      const result = await SignOut()
      if (result) {
        dispatch(setAuth(null))
        localStorage.removeItem('token')
        localStorage.removeItem('user')
        toast.success('Logout successfully.')
        navigate('/')
      }
    } catch (e: any) {
      toast.error('Something went wrong.')
    }
  }

  return (
    <div className="flex bg-[#333333] text-white items-center justify-between px-[150px] py-4">
      <div className="flex items-center gap-2">
        <CiLocationOn />
        <p className="text-sm font-normal">Tran Phuoc Anh Quoc, Software Engineer, UIT</p>
      </div>
      <div className="flex gap-3 text-sm">
        <div className="flex gap-2 text-sm items-center">
          {user ? (
            <>
              <a href="/signin" className="flex items-center gap-2 hover:cursor-pointer hover:underline">
                <div className="avatar">
                  <div className="w-6 rounded-full">
                    <img src={user.avatar_url} />
                  </div>
                </div>
                <p>{user.name}</p>
              </a>
              <span>/</span>
              <button onClick={handleLogout} className="hover:cursor-pointer hover:underline">
                <p>{isLoading ? <Loading /> : 'Logout'}</p>
              </button>
            </>
          ) : (
            <>
              <a href="/signin" className="hover:cursor-pointer hover:underline">
                <p className={`hover:underline ${location.pathname === '/signin' && 'text-green-600 font-bold'}`}>
                  SignIn
                </p>
              </a>
              <span>/</span>
              <a href="/signup" className="hover:cursor-pointer hover:underline">
                <p className={`hover:underline ${location.pathname === '/signup' && 'text-green-600 font-bold'}`}>
                  SignUp
                </p>
              </a>
            </>
          )}
        </div>
      </div>
    </div>
  )
}

export default Header
