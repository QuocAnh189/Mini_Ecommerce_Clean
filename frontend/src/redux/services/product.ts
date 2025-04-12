import { createApi } from '@reduxjs/toolkit/query/react'

//interceptors
import baseQueryWithReauth from '@redux/interceptor/baseQueryWithReauth'

//interfaces
import { IListProductRequest, IProduct } from '@interfaces/product'
import { IListData } from '@interfaces/common'

export const apiProduct = createApi({
  reducerPath: 'apiProduct',
  baseQuery: baseQueryWithReauth,
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
