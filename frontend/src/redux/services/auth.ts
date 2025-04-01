import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'

//interfaces
import { IAuth, SingInRequest } from '@interfaces/user'

export const apiAuth = createApi({
  reducerPath: 'apiAuth',
  baseQuery: fetchBaseQuery({
    baseUrl: import.meta.env.VITE_API_URL,
  }),
  keepUnusedDataFor: 20,

  endpoints: (builder) => ({
    signUp: builder.mutation<IAuth, FormData>({
      query: (data) => ({
        url: '/auth/signup',
        method: 'POST',
        body: data,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
    }),

    signIn: builder.mutation<IAuth, SingInRequest>({
      query: (data) => ({
        url: '/auth/signin',
        method: 'POST',
        body: data,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
    }),

    signOut: builder.mutation<IAuth, void>({
      query: () => ({
        url: '/auth/signout',
        method: 'POST',
        headers: {
          Authorization: `Bearer ${JSON.parse(localStorage.getItem('token')!).accessToken}`,
        },
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
    }),
  }),
})

export const { useSignInMutation, useSignUpMutation, useSignOutMutation } = apiAuth
