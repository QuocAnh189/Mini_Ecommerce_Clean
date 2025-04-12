import { createApi } from '@reduxjs/toolkit/query/react'

//interceptors
import baseQueryWithReauth from '@redux/interceptor/baseQueryWithReauth'

//interfaces
import { IAddProductRequest, ICart, IRemoveProductRequest, IUpdateCartLineRequest } from '@interfaces/cart'

export const apiCart = createApi({
  reducerPath: 'apiCart',
  baseQuery: baseQueryWithReauth,
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
