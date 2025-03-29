//icons
import { CiLocationOn } from 'react-icons/ci'

const Header = () => {
  return (
    <div className="flex bg-black text-white items-center justify-between px-[150px] py-4">
      <div className="flex items-center gap-2">
        <CiLocationOn />
        <p className="text-sm font-normal">Tran Phuoc Anh Quoc, Software Engineer, UIT</p>
      </div>
      <div className="flex gap-3 text-sm">
        <div className="flex gap-2 text-sm items-center">
          {false ? (
            <>
              <a
                href="/signin"
                className="flex items-center gap-2 hover:cursor-pointer hover:underline"
              >
                {/* <Avatar size="xs" src={auth.user.avatar} />
                <p>{auth.user.name}</p> */}
              </a>
              <span>/</span>
              <button className="hover:cursor-pointer hover:underline">
                <p>Logout</p>
              </button>
            </>
          ) : (
            <>
              <a href="/signin" className="hover:cursor-pointer hover:underline">
                SignIn
              </a>
              <span>/</span>
              <a href="/signup" className="hover:cursor-pointer hover:underline">
                SignUp
              </a>
            </>
          )}
        </div>
      </div>
    </div>
  )
}

export default Header
