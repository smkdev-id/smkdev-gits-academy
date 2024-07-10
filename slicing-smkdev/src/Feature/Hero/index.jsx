import React from 'react'
import { TypeAnimation } from 'react-type-animation';
import skils1 from '../../assets/hero-img.png'

const Hero = () => {
  return (
    <div >
      <div className='flex container mx-auto px-52 py-24 gap-4 justify-between items-center'>
        <div className='flex flex-col gap-4'>
          <p className='text-5xl font-semibold'>Jadilah Talent Digital</p>
          <TypeAnimation
            sequence={[
              'Global',
              5000,
              'Masa Depan Indonesia',
              5000,
            ]}
            wrapper="span"
            speed={20}
            className='text-5xl font-semibold text-[#4C3BCF]'
            // style={{ fontSize: '2em', display: 'inline-block', color: '#3FA2F6' }}
            repeat={Infinity}
            />
          <p>Belajar langsung dengan expert dari industri dengan kurikulum komperhensif <br/> berbasis project-based learning</p>
        </div>
        <div>
          <img src={skils1} alt="skills" className='w-[400px] rounded-3xl'/>
        </div>
      </div>
    </div>
  )
}

export default Hero