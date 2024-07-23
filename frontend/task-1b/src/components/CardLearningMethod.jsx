import React from 'react'

export const CardLearningMethod = ({img, title, desc}) => {
  return (
    <div className='w-full sm:w-[32%] bg-white rounded-[20px] cursor-pointer drop-shadow-[5px_20px_15px_rgba(39,93,132,0.05)]'>
        <img src={`${img}`} alt="" />
        <div className='mx-[32px] mb-[36px]'>
            <p className='text-[20px] text-[#244D73] font-semibold mt-[30px] mb-[20px]'>{title}</p>
            <p className='text-[#828282]'>{desc}</p>
        </div>
    </div>
  )
}
