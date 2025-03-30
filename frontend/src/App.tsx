import { Suspense, lazy } from 'react'
import { BrowserRouter, Routes, Route } from 'react-router-dom'

//components
import Layout from './layouts'

//page
const SignInPage = lazy(() => import('@pages/auth/SignIn'))
const SignUpPage = lazy(() => import('@pages/auth/SignUp'))
const ProductPage = lazy(() => import('@pages/product/Product'))
const CartPage = lazy(() => import('@pages/cart/Cart'))
const OrderPage = lazy(() => import('@pages/order/Order'))
const OrderLinePage = lazy(() => import('@pages/order/OrderLine'))
const ProfilePage = lazy(() => import('@pages/profile/Profile'))

//component
import Loader from '@components/Loader'

function App() {
  return (
    <BrowserRouter>
      <Suspense fallback={<Loader />}>
        <Routes>
          <Route path="/" element={<Layout />}>
            <Route path="/signin" element={<SignInPage />} />
            <Route path="/signup" element={<SignUpPage />} />
            <Route path="/product" element={<ProductPage />} />
            <Route path="/cart" element={<CartPage />} />
            <Route path="/order" element={<OrderPage />} />
            <Route path="/order/:id" element={<OrderLinePage />} />
            <Route path="/profile" element={<ProfilePage />} />
          </Route>
        </Routes>
      </Suspense>
    </BrowserRouter>
  )
}

export default App
