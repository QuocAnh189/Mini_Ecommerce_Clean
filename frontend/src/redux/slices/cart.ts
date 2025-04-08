//redux
import { PayloadAction, createSlice } from '@reduxjs/toolkit'

export const CartSliceKey = 'cart'

type InitialType = {
  cartId: string | null
}

const initialState = {
  cartId: null,
} as InitialType

const cartSlice = createSlice({
  name: CartSliceKey,
  initialState,
  reducers: {
    setCart: (state, action: PayloadAction<string | null>) => {
      state.cartId = action.payload
    },
  },
})

export const { setCart } = cartSlice.actions

const CartReducer = cartSlice.reducer
export default CartReducer
