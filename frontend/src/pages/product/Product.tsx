//components
import CardProduct from '@components/CardProduct'

const Product = () => {
  return (
    <div className="min-h-screen">
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 justify-start px-10 py-5">
        <CardProduct />
        <CardProduct />
        <CardProduct />
        <CardProduct />
      </div>
    </div>
  )
}

export default Product
