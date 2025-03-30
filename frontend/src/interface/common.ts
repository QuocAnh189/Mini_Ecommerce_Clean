export interface IListData<T> {
  items: T[]
  metadata: IMetadata
}

export interface IMetadata {
  page: number
  size: number
  take_all: boolean
  skip: number
  total_count: number
  total_pages: number
  has_previous: boolean
  has_next: boolean
}
