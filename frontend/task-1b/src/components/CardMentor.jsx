import React from 'react'

export const CardMentor = ({name, job, company, img}) => {
  return (
    <div className='w-full px-[20px] flex flex-col items-center cursor-pointer'>
      <img src={`${img}`} alt="" className='w-full rounded-full bg-cover'/>
      <p className='text-[20px] font-semibold mt-[17px] mb-[10px] text-center'>{name}</p>
      <div className='w-[106px] h-[2px] bg-black'></div>
      <p className='mt-[10px] text-center'>{job}</p>
      <p className='mt-[10px] text-[14px] text-center'>{company}</p>
    </div>
  )
}
