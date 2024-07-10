import React from 'react'
import { Carousel } from 'antd';
import { logos } from '../../constant/constant';
import { shuffledLogos } from '../../constant/constant';
import { shuffledLogos1 } from '../../constant/constant';

const Mitra = () => {
return (
    <div>
      <div className='flex flex-col gap-0 bg-[#EEEDEB] pt-10'>
        <div className='text-center'>
          <h1 className='text-xl font-semibold text-[#4C3BCF]'>Dipercaya Oleh Mitra Industri: </h1>
        </div>
        <Carousel autoplay autoplaySpeed={3000} autoplayInterval={3000} fade dots={false}>
          <div className="h-40 items-center text-black leading-40 text-center">
              <div className='flex gap-4 justify-center items-center' >
                {
                  logos.map((logo) => (  
                    <img key={logo.id} src={logo.src} alt="logo Mitra" className='w-[150px]'/>
                  ))
                }
              </div>
          </div>
          <div className="h-40 items-center text-black leading-40 text-center">
              <div className='flex gap-4 justify-center items-center' >
                {
                  shuffledLogos1.map((logo) => (
                    <img key={logo.id} src={logo.src} alt="logo Mitra" className='w-[150px]' />
                  ))
                }
              </div>
          </div>
          <div className="h-40 items-center text-black leading-40 text-center">
              <div className='flex gap-4 justify-center items-center' >
                {
                  shuffledLogos.map((logo) => (
                    <img key={logo.id} src={logo.src} alt="logo Mitra" className='w-[150px]' />
                  ))
                }
              </div>
          </div>
        </Carousel>
        </div>
    </div>
  )
}

export default Mitra