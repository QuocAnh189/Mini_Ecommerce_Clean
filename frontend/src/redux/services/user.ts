import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'

//interfaces
import { IListUserRequest, IUser } from '@interfaces/user'
import { IListData } from '@interfaces/common'

export const apiUser = createApi({
  reducerPath: 'apiUser',
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

  endpoints: (builder) => ({
    getListUsers: builder.query<IListData<IUser>, IListUserRequest>({
      query: (params) => ({
        url: '/users',
        method: 'GET',
        params: params,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
    }),

    getUser: builder.query<IUser, string>({
      query: (userId) => ({
        url: `/users/${userId}`,
        method: 'GET',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
    }),

    deleteUser: builder.mutation<string, string>({
      query: (userId) => ({
        url: `/users/${userId}`,
        method: 'DELETE',
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
    }),
  }),
})

export const { useGetListUsersQuery, useGetUserQuery, useDeleteUserMutation } = apiUser
