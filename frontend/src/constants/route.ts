export interface IRoute {
  name: string
  path: string
  icon?: any
}

export const routes: IRoute[] = [
  { name: 'Product', path: '/product' },
  { name: 'Order', path: '/order' },
  { name: 'Profile', path: '/profile' },
]
