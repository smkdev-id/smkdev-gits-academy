import React from 'react'
import { Card } from 'antd';
import { cardsOrientasi } from '../../constant/constant';

const Orientasi = () => {
  return (
    <div>
        <div className='text-center p-8'>
            <h1 className='text-3xl font-semibold py-4'>Orientasi Belajar SMKDEV</h1>
            <p className='text-xl font-semibold'>Dapatkan pengalaman belajar berorientasi pengalaman kerja yang dapat mengantarkan <br /> Anda menjadi talenta yang dibutuhkan oleh industri digital terkini.</p>
        </div>
        <div className='flex container mx-auto justify-evenly py-4'>
            {cardsOrientasi.map((card) => (
                <div key={card.id} className=''>
                    <Card
                    hoverable
                    bordered={false}
                    className='text-center w-[400px] h-[300px] shadow-lg shadow-indigo-500/50 rounded-xl'
                    >
                    <div className='flex justify-center'>
                        <img src={card.src} alt="icon" className='w-[100px] h-[100px]  bg-white rounded-xl'/>
                    </div>  
                    <div className='text-center'>
                        <p className='text-xl font-semibold py-4'>{card.title}</p>
                        <p className='text-lg'>{card.desc}</p>
                    </div>
                    </Card>
                </div>
            ))}
        </div>
    </div>
  )
}

export default Orientasi