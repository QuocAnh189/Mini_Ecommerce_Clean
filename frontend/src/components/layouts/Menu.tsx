'use client'

import { useLocation } from 'react-router-dom'

//icons
import { FaPhoneVolume } from 'react-icons/fa6'

//constant
import { routes, IRoute } from '@constants/route'

const Menu = () => {
  const location = useLocation()

  return (
    <div className="flex bg-gray-300 text-black items-center justify-between px-[150px] py-3">
      <div className="flex items-center gap-8 text-gray500">
        {routes.map((route: IRoute, index: number) => (
          <a
            key={index}
            href={route.path}
            className={`hover:underline ${
              route.path === location.pathname && 'text-green-600 font-bold'
            }`}
          >
            {route.name}
          </a>
        ))}
      </div>
      <div className="flex items-center gap-2">
        <FaPhoneVolume size="24px" />
        <p className="text-black">(219) 555-2214</p>
      </div>
    </div>
  )
}

export default Menu
