export interface IRoute {
  name: string
  path: string
  icon?: any
  isAdmin?: boolean
}

export const routes: IRoute[] = [
  { name: 'Home', path: '/' },
  { name: 'Product', path: '/product' },
  { name: 'Order', path: '/order' },
  { name: 'Profile', path: '/profile' },
  { name: 'User', path: '/users', isAdmin: true },
]
