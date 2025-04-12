import { createApi } from '@reduxjs/toolkit/query/react'

//interceptors
import baseQueryWithReauth from '@redux/interceptor/baseQueryWithReauth'

//interfaces
import { IListUserRequest, IUser } from '@interfaces/user'
import { IListData } from '@interfaces/common'

export const apiUser = createApi({
  reducerPath: 'apiUser',
  baseQuery: baseQueryWithReauth,
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
