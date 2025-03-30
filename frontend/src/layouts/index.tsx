//components
import Header from '@components/layouts/Header'
import AppBar from '@components/layouts/AppBar'
import Menu from '@components/layouts/Menu'
import Footer from '@components/layouts/Footer'
import { Outlet } from 'react-router-dom'

//toast
import { Toaster } from 'react-hot-toast'

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
