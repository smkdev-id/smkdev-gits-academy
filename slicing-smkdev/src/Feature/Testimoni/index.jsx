import React from 'react'
import { Card } from 'antd';
import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/css';
import { testiCards } from '../../constant/constant';


const Testimoni = () => {
  return (
    <div>
        <div className='container mx-auto justify-center py-12'>
            {/* title - desc */}
            <div className='flex flex-col items-center'>
                <h1 className='text-4xl font-semibold'>Apa Kata Mereka</h1>
                <p className='text-xl py-4'>Mereka telah merasakan serunya belajar skill digital di SMKDEV. Apakah kamu ingin seperti mereka?</p>
            </div>
            {/* card */}
            <div className='container mx-auto px-40 py-4'>
                <Swiper
                    spaceBetween={50}
                    slidesPerView={3}
                    loop={true}
                    loopedSlides={testiCards.length}
                    onSlideChange={() => console.log('slide change')}
                    onSwiper={(swiper) => console.log(swiper)}
                    >
                        {testiCards.map((testi) => (
                            <SwiperSlide key={testi.id} className='flex justify-center'>
                                <Card
                                    bordered={true}
                                    className='text-center w-[300px] h-max shadow-lg rounded-xl '
                                    key={testi.id}
                                    >
                                <div className='text-left'>
                                    <p className='text-xl italic py-4'>{testi.desc}</p>
                                    <p className='text-2xl italic font-semibold'>{testi.name}</p>
                                    <p className='text-lg italic font-selgem'>{testi.major}</p>
                                </div>
                                </Card>
                            </SwiperSlide>
                        ))} 
                </Swiper>
            </div>
        </div>
    </div>
  )
}

export default Testimoni