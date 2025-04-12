import { fetchBaseQuery } from '@reduxjs/toolkit/query/react'
import type { BaseQueryFn, FetchArgs, FetchBaseQueryError } from '@reduxjs/toolkit/query'
import { Mutex } from 'async-mutex'

const baseQuery = fetchBaseQuery({
  baseUrl: import.meta.env.VITE_API_URL,
  prepareHeaders: (headers) => {
    const token = JSON.parse(localStorage.getItem('token')!)?.accessToken
    if (token) {
      headers.set('Authorization', `Bearer ${token}`)
    }
    return headers
  },
})

const rawBaseQuery = fetchBaseQuery({
  baseUrl: import.meta.env.VITE_API_URL,
})

const mutex = new Mutex()

const baseQueryWithReauth: BaseQueryFn<string | FetchArgs, unknown, FetchBaseQueryError> = async (
  args,
  api,
  extraOptions,
) => {
  await mutex.waitForUnlock()

  let result = await baseQuery(args, api, extraOptions)

  if (result.error && result.error.status === 401) {
    if (!mutex.isLocked()) {
      const release = await mutex.acquire()

      try {
        const refreshToken = JSON.parse(localStorage.getItem('token')!)?.refreshToken

        const refreshResult: any = await rawBaseQuery(
          {
            url: '/auth/refresh-token',
            method: 'POST',
            headers: {
              Authorization: `Bearer ${refreshToken}`,
            },
          },
          api,
          extraOptions,
        )

        if (refreshResult.data) {
          const newToken = { accessToken: refreshResult.data.data.accessToken, refreshToken }

          localStorage.setItem('token', JSON.stringify(newToken))

          // Retry lại request ban đầu với token mới
          result = await baseQuery(args, api, extraOptions)
        } else {
          api.dispatch({ type: 'auth/signout' })
        }
      } finally {
        release()
      }
    } else {
      await mutex.waitForUnlock()
      result = await baseQuery(args, api, extraOptions)
    }
  }

  return result
}

export default baseQueryWithReauth
