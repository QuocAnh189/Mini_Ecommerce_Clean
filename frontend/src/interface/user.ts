export interface IUser {
  id: string
  email: string
  name: string
  avatar_url: string
  created_at: string
  updated_at: string
}

export interface SingInRequest {
  email: string
  password: string
}

export interface SingInResponse {
  accessToken: string
  refreshToken: string
  user: IUser
}

export interface SingUpResponse {
  accessToken: string
  refreshToken: string
  user: IUser
}

export interface SignUpRequest {
  email: string
  name: string
  avatar: any
  password: string
}
