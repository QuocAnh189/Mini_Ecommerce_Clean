import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'

//interfaces
import { IListOrderRequest, IOrder, PlaceOrderRequest } from '@interfaces/order'
import { IListData } from '@interfaces/common'

export const apiOrder = createApi({
  reducerPath: 'apiOrder',
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
  tagTypes: ['Order'],

  endpoints: (builder) => ({
    placeOrder: builder.mutation<IOrder, PlaceOrderRequest>({
      query: (data) => ({
        url: '/orders',
        method: 'POST',
        body: data,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Order'],
    }),

    getListMyOrders: builder.query<IListData<IOrder>, IListOrderRequest>({
      query: (params) => ({
        url: `/orders`,
        method: 'GET',
        params: params,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      providesTags: ['Order'],
    }),

    getOrderById: builder.query<IOrder, string>({
      query: (orderId) => ({
        url: `/orders/${orderId}`,
        method: 'GET',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
    }),

    updateStatusOrder: builder.mutation<IOrder, { orderId: string; status: string }>({
      query: ({ orderId, status }) => ({
        url: `/orders/${orderId}/${status}`,
        method: 'PUT',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ['Order'],
    }),
  }),
})

export const { usePlaceOrderMutation, useGetListMyOrdersQuery, useGetOrderByIdQuery, useUpdateStatusOrderMutation } =
  apiOrder
