import React from 'react'

export const CardTestimony = ({name, job, desc}) => {
  return (
    <div className='w-full h-[365px] p-[22px] bg-white cursor-pointer drop-shadow-[5px_5px_15px_rgba(39,93,132,0.05)] snap-center'>
        <p className='italic w-full h-[82%] overflow-hidden text-ellipsis '>{desc}</p>
        <p className='font-semibold italic mt-[14px]'>{name}</p>
        <p className='text-[14px] italic mt-[3px]'>{job}</p>
    </div>
  )
}
