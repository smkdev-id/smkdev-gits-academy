import React from 'react'
import { Card} from 'antd';
import talent from '../../assets/talent.png'
import { cardsWhy } from '../../constant/constant';


const WhySmkDev = () => {
  return (
    <div>
        <div className='flex container mx-auto justify-evenly py-24 px-36'>
            {/* card */}
            <div className='flex flex-col gap-4'>
                {cardsWhy.map((card) => (
                    <div key={card.id} className="">
                    {/* card item */}
                    <Card 
                        hoverable
                        className='rounded-3xl shadow-xl w-[600px]'>
                            <div className='flex gap-6'>
                                <div className='p-4 bg-gray-100 hover:bg-gray-300  rounded-lg flex items-center'>
                                    <img src={card.src} alt="icon" className='w-[70px]'/>
                                </div>
                                <div className=''>
                                    <p className='text-3xl font-bold'>{card.title}</p>
                                    <p className='text-lg font-semibold'>{card.desc}</p>
                                </div> 
                            </div>
                    </Card>
                </div>
                ))}
            </div>
            {/* desc */}
            <div className='flex flex-col gap-4 justify-center'>
                <p className='text-3xl font-bold'>MENGAPA HARUS MEMILIH BELAJAR <br />DI SMKDEV</p>
                <p className='text-lg'>Pengalaman belajarmu jauh lebih menyenangkan dengan perpaduan <br /> belajar teori dan capstone project yang dapat menjadi portofolio <br />pribadimu. Tidak lupa, Kamu juga dapat belajar langsung dari expert <br />mentor dari dunia industri.</p>
            </div>
        </div>
        {/* Talent Digital SMKDEV */}
        <div className='flex flex-col justify-center py-6'>
            <div className='flex justify-center'>
                <h1 className='text-4xl font-semibold'>Talent Digital SMKDEV</h1> 
            </div>
            <div className='flex justify-center'>
                <img src={talent} alt="talent"/>  
            </div>
        </div>
    </div>
  )
}

export default WhySmkDev