import React from 'react'

export const Footer = () => {
    const SocialMedia = ["youtube", "linkedin", "instagram", "tiktok", "telegram", "whatsapp",]
    return (
        <div>
            <div className='w-full h-full py-auto bg-[#F3F8FB] py-[40px] px-[20px] sm:px-[30px] lg:px-[70px] xl:px-[140px]'>
                <img src="Logo-smkdev.png" alt="" width={128} height={32}/>
                <p className="text-[#616364] font-semibold mt-[10px]">Creating High-Caliber Digital Talent</p>
                <p className='text-[#616364] my-[20px]'>Summarecon Bandung, Jl. Magna Barat Blok MD No.02, Rancabolang, Kec. Gedebage, Kota Bandung, Jawa Barat 40295</p>
                <p className='text-[#244D73] font-bold mb-[10px]'>Follow Us</p>
                <div className='w-[275px] flex justify-between'>
                    {SocialMedia.map(sosmed => {
                        return (
                            <img src={`${sosmed}.svg`} alt="" className="cursor-pointer"/>
                        )
                    })}
                </div>
            </div>
            <div className='w-full bg-white p-[20px]'>
                <p className='text-[#1C3965] text-center'>© 2024 SMKDEV – PT Eureka Merdeka Indonesia. All Rights Reserved.</p>
            </div>
        </div>
    )
}   
