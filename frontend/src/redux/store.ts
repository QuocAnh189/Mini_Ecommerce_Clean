import { configureStore } from '@reduxjs/toolkit'

//service
import { apiAuth } from './services/auth'
import { apiUser } from './services/user'
import { apiProduct } from './services/product'
import { apiCart } from './services/cart'
import { apiOrder } from './services/order'

//slice
import authReducer, { AuthSliceKey } from './slices/auth'

const store = configureStore({
  reducer: {
    [apiAuth.reducerPath]: apiAuth.reducer,
    [apiUser.reducerPath]: apiUser.reducer,
    [apiProduct.reducerPath]: apiProduct.reducer,
    [apiCart.reducerPath]: apiCart.reducer,
    [apiOrder.reducerPath]: apiOrder.reducer,

    [AuthSliceKey]: authReducer,
  },

  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false,
      immutableCheck: false,
    }).concat([apiAuth.middleware, apiUser.middleware, apiProduct.middleware, apiCart.middleware, apiOrder.middleware]),
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch

export default store
