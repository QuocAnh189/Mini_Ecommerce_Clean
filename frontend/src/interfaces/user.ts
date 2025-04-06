import { ERole } from '@constants/enum'

export interface IUser {
  id: string
  email: string
  name: string
  avatar_url: string
  role: ERole
  created_at: string
  updated_at: string
  deleted_at: string
}

export interface IAuth {
  accessToken: string
  refreshToken: string
  user: IUser
}

export interface SingInRequest {
  email: string
  password: string
}

export interface SignUpRequest {
  email: string
  name: string
  avatar: any
  password: string
  role: ERole
}

export interface IListUserRequest {
  search: string
  page: number
  size: number
  order_by: string
  order_desc: boolean
  take_all: boolean
}
