import { PayloadAction, createSlice } from '@reduxjs/toolkit'
import { IAuth } from '@interfaces/user'

export const AuthSliceKey = 'auth'

type InitialType = {
  auth: IAuth | null
}

const initialState = {
  auth: null,
} as InitialType

const authSlice = createSlice({
  name: AuthSliceKey,
  initialState,
  reducers: {
    setAuth: (state, action: PayloadAction<IAuth | null>) => {
      state.auth = action.payload
    },
  },
})

export const { setAuth } = authSlice.actions

const authReducer = authSlice.reducer
export default authReducer
