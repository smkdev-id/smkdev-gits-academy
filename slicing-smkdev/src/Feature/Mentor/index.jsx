import React from 'react'
import { Card } from 'antd';
import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/css';
import 'swiper/css/autoplay';
import { mentors } from '../../constant/constant';

const Mentor = () => {
  return (
    <div>
      <div >
          {/* title - desc */}
          <div className='flex flex-col items-center py-4'>
            <h1 className='text-3xl py-2 font-semibold'>Expert Mentor</h1>
            <p className='text-lg pb-1 '>SMKDEV menghadirkan expert dari industri untuk mendampingi proses belajarmu.</p>
          </div>
          {/* mentor */}
          <div className='container mx-auto px-32'>
              <Swiper
                  spaceBetween={50}
                  slidesPerView={4}
                  autoplay={true}
                  loop={true}
                  loopedSlides={mentors.length}
                  onSlideChange={() => console.log('slide change')}
                  onSwiper={(swiper) => console.log(swiper)}
                  className=''
                  >
                  {mentors.map((mentor) => (
                    <SwiperSlide key={mentor.id} className='flex justify-center'>
                      <Card
                      bordered={false}
                      className='py-4 w-[250px] h-[400px]'>
                        <div className='flex justify-center'> 
                          <img src={mentor.src} alt="mentor" className='w-[190px] rounded-full'/>
                        </div>
                        <div className='text-center py-4'>
                          <p className='text-lg font-semibold border-b-2 py-1'>{mentor.name}</p>
                          <p className='text-lg font-normal py-2'>{mentor.role}</p>
                          <p className='text-sm font-medium'>{mentor.Workplace}</p>
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

export default Mentor