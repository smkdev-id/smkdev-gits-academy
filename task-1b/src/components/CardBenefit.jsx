import React from 'react'

export const CardBenefit = ({img, title, desc}) => {
  return (
    <div className='w-full h-[145px] p-[16px] bg-white flex items-center rounded-[16px] drop-shadow-lg cursor-pointer'>
        <img src={`${img}`} alt="" />
        <div className='ml-[16px]'>
            <p className='text-[20px] sm:text-[24px] text-[#1C3965] font-semibold'>{title}</p>
            <p className='mt-[5px] sm:mt-[10px]'>{desc}</p>
        </div>
    </div>
  )
}
