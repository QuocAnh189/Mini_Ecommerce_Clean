//components
import Header from '@components/layouts/Header'
import AppBar from '@components/layouts/AppBar'
import Menu from '@components/layouts/Menu'
import { Outlet } from 'react-router-dom'

const Layout = () => {
  return (
    <div className="">
      <Header />
      <AppBar />
      <Menu />
      <Outlet />
    </div>
  )
}

export default Layout
