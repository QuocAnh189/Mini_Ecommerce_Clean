interface IProps {
  image: any
  left: boolean
}

//icon
import { FaArrowRight } from 'react-icons/fa'

const BannerItem = (props: IProps) => {
  const { image, left } = props

  return (
    <div className="w-full h-full flex items-center justify-center px-6 gap-6">
      {left && <img src={image} alt="" className="w-2/5 object-cover" />}
      <div className="flex flex-col items-start gap-4">
        <p className="text-primary text-md">WELCOME TO SHOPERY</p>
        <h1 className="text-black text-6xl font-bold">Fresh & Healthy</h1>
        <h1 className="text-black text-6xl font-bold">Organic Food</h1>
        <p className="text-xl font-semibold">
          Sale up to <span className="">30% OFF</span>
        </p>
        <p className="text-sm text-gray500">
          Free shipping on all your order we delivery, you enjoin
        </p>
        <button className="flex items-center gap-3 bg-green-600 rounded-3xl px-6 py-2">
          <p className="text-white font-medium">Show now</p>
          <FaArrowRight color="white" />
        </button>
      </div>
      {!left && <img src={image} height={0} width={0} alt="" className="w-2/5 object-cover" />}
    </div>
  )
}

export default BannerItem
