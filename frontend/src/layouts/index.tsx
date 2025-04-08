//components
import { Outlet } from 'react-router-dom'
import { Toaster } from 'react-hot-toast'
import Header from '@components/layouts/Header'
import AppBar from '@components/layouts/AppBar'
import Menu from '@components/layouts/Menu'
import Footer from '@components/layouts/Footer'

const Layout = () => {
  return (
    <div className="">
      <Header />
      <AppBar />
      <Menu />
      <Outlet />
      <Footer />
      <Toaster />
    </div>
  )
}

export default Layout
