import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'

//interfaces
import { IListProductRequest, IProduct } from '@interfaces/product'
import { IListData } from '@interfaces/common'

export const apiProduct = createApi({
  reducerPath: 'apiProduct',
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
  tagTypes: ['Product'],

  endpoints: (builder) => ({
    getListProducts: builder.query<IListData<IProduct>, IListProductRequest>({
      query: (params) => ({
        url: '/products',
        method: 'GET',
        params: params,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      providesTags: ['Product'],
    }),

    getProduct: builder.query<IProduct, string>({
      query: (productId) => ({
        url: `/products/${productId}`,
        method: 'GET',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      providesTags: ['Product'],
    }),

    createProduct: builder.mutation<string, FormData>({
      query: (data) => ({
        url: `/products`,
        method: 'POST',
        body: data,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Product'],
    }),

    updateProduct: builder.mutation<string, FormData>({
      query: (data) => ({
        url: `/products/${data.get('id')}`,
        method: 'PUT',
        body: data,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Product'],
    }),

    deleteProduct: builder.mutation<string, string>({
      query: (productId) => ({
        url: `/products/${productId}`,
        method: 'DELETE',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Product'],
    }),
  }),
})

export const {
  useGetListProductsQuery,
  useGetProductQuery,
  useCreateProductMutation,
  useUpdateProductMutation,
  useDeleteProductMutation,
} = apiProduct
