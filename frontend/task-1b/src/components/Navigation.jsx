import React from 'react'

export const Navigation = () => {
  return (
    <nav className='w-screen sm:w-full h-[60px] px-[20px] sm:px-[30px] lg:px-[70px] xl:px-[140px] flex gap-6 border-b-2 bg-white border-gray-200 fixed z-50'>
        <div className='flex justify-between w-[88%] py-[17px]'>
          <div className='h-full flex items-center sm:hidden'>
            <svg width="20" height="15" viewBox="0 0 20 15" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M1 1.83496H19M1 7.83496H19M1 13.835H19" stroke="black" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
          </div>
            <img src="Logo-smkdev.png" alt="" className='mr-[-30px] sm:mr-0'/>
            <ul className='gap-8 hidden sm:flex'>
                <li className='text-[14px] sm:text-[16px] lg:text-[18px] text-[#1C3965] font-semibold cursor-pointer'>Learn</li>
                <li className='text-[14px] sm:text-[16px] lg:text-[18px] text-[#1C3965] font-semibold cursor-pointer'>Community</li>
                <li className='text-[14px] sm:text-[16px] lg:text-[18px] text-[#1C3965] font-semibold cursor-pointer'>Blog</li>
            </ul>
        </div>
        <button className='bg-[#1C3965] sm:w-[120px] md:w-[125px] lg:w-[140px] h-[38px] my-auto rounded-[10px] text-[14px] sm:text-[16px] lg:text-[18px] text-white font-semibold hidden sm:inline-block'>Dashboard</button>
    </nav>
  )
}
