import React from 'react'
// Import Swiper React components
// import { Swiper, SwiperSlide } from 'swiper/react';

// Import Swiper styles
// import 'swiper/css';
import { Carousel } from 'antd';
import gambar1 from '../../assets/digits.png'
import gambar2 from '../../assets/arkana.png'
import gambar3 from '../../assets/eudeka.png'
import gambar4 from '../../assets/gitsid.png'
import gambar5 from '../../assets/hitopia.png'
import gambar6 from '../../assets/mantab.png'

const Mitra = () => {
  const logos = [{
    id: 1,
    src: gambar1
  },
  {
    id: 2,
    src: gambar2
  },
  {
    id: 3,
    src: gambar3
  },
  {
    id: 4,
    src: gambar4
  },
  {
    id: 5,
    src: gambar5
  },
  {
    id: 6,
    src: gambar6
  }
]
const shuffledLogos = [...logos].sort(() => Math.random() - 0.5);
return (
    <div>
      <div className='text-center bg-[#EEEDEB] py-8' >
        <h1 className='text-xl font-semibold text-[#4C3BCF]'>Dipercaya Oleh Mitra Industri: </h1>
      </div>
      <Carousel autoplay autoplaySpeed={5000} autoplayInterval={10000}>
        <div className="h-40 items-center text-black leading-40 text-center bg-[#EEEDEB]">
            <div className='flex gap-4 justify-center items-center' >
              {
                logos.map((logo) => (
                  <img key={logo.id} src={logo.src} alt="logo Mitra" className='w-[100px]'/>
                ))
              }
            </div>
        </div>
        <div className="h-40 items-center text-black leading-40 text-center bg-[#EEEDEB]">
            <div className='flex gap-4 justify-center items-center' >
              {
                shuffledLogos.map((logo) => (
                  <img key={logo.id} src={logo.src} alt="logo Mitra" className='w-[100px]' />
                ))
              }
            </div>
        </div>
      </Carousel>
    </div>
  )
}

export default Mitra