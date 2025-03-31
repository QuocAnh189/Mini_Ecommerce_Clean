//components
import BannerItem from '@components/BannerItem'

//swiper
import { Swiper, SwiperSlide } from 'swiper/react'
import { Pagination, Navigation } from 'swiper/modules'
import 'swiper/css'
import 'swiper/css/navigation'
import 'swiper/css/pagination'

//image
import image1 from '@assets/images/BG1.png'
import image2 from '@assets/images/BG2.png'
import image3 from '@assets/images/BG3.png'

const Home = () => {
  return (
    <Swiper navigation={true} modules={[Pagination, Navigation]} className="z-[1] h-screen">
      <SwiperSlide className="w-full z-[1]">
        <BannerItem image={image1} left={true} />
      </SwiperSlide>
      <SwiperSlide className="w-full z-[1]">
        <BannerItem image={image2} left={false} />
      </SwiperSlide>
      <SwiperSlide className="w-full z-[1]">
        <BannerItem image={image3} left={true} />
      </SwiperSlide>
    </Swiper>
  )
}

export default Home
