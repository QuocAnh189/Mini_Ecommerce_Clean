import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'

//interfaces
import { IAddProductRequest, ICart, IRemoveProductRequest, IUpdateCartLineRequest } from '@interfaces/cart'

export const apiCart = createApi({
  reducerPath: 'apiCart',
  baseQuery: fetchBaseQuery({
    baseUrl: import.meta.env.VITE_API_URL,
    prepareHeaders: (headers) => {
      const token = JSON.parse(localStorage.getItem('token')!)?.accessToken

      if (token) {
        headers.set('Authorization', `Bearer ${token}`)
      }

      return headers
    },
  }),
  keepUnusedDataFor: 20,
  tagTypes: ['Cart'],

  endpoints: (builder) => ({
    GetCart: builder.query<ICart, string>({
      query: (userId) => ({
        url: `/carts/${userId}`,
        method: 'GET',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      providesTags: ['Cart'],
    }),

    addProductToCart: builder.mutation<string, { userId: string; data: IAddProductRequest }>({
      query: ({ userId, data }) => ({
        url: `/carts/${userId}`,
        method: 'POST',
        body: data,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Cart'],
    }),

    updateCartLine: builder.mutation<string, { userId: string; data: IUpdateCartLineRequest }>({
      query: ({ userId, data }) => ({
        url: `/carts/cart-line/${userId}`,
        method: 'PUT',
        body: data,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Cart'],
    }),

    removeProductFromCart: builder.mutation<string, { userId: string; data: IRemoveProductRequest }>({
      query: ({ userId, data }) => ({
        url: `/carts/${userId}`,
        method: 'DELETE',
        body: data,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Cart'],
    }),
  }),
})

export const {
  useGetCartQuery,
  useAddProductToCartMutation,
  useUpdateCartLineMutation,
  useRemoveProductFromCartMutation,
} = apiCart
