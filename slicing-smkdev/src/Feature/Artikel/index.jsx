import React from 'react'
import { Card } from 'antd';
import { artikels } from '../../constant/constant';

const Artikel = () => {
  return (
    <div>
        <div className='flex flex-col items-center gap-4 py-12'>
            <div>
                <h1 className='text-3xl font-semibold'>Artikel Pilihan SMKDEV</h1>
            </div>
            <div className='flex container mx-auto justify-center py-4 gap-4' >
                {artikels.map((artikel) => (
                    <div key={artikel.id} >
                        <Card
                        bordered={true}
                        hoverable
                        className='text-center w-[300px] shadow-lg rounded-xl'
                        >
                        <div className='flex justify-center'>
                            <a href={artikel.link}>
                                <img src={artikel.src} alt="icon" className='w-[200px] h-[200px]  bg-white rounded-xl'/>
                            </a>
                        </div>  
                        <div className='text-left'>
                            <p className='text-xl font-semibold py-4'>{artikel.title}</p>
                            <p className='text-lg'>{artikel.desc}...</p>
                            <a href={artikel.link}>
                                <p className='text-lg py-4 italic'>Baca Lebih Lanjut</p>
                            </a>
                        </div>
                        </Card>
                    </div>
                ))}
            </div>
            <div>
                <a href="https://www.smk.dev/blog/" target='_blank'>
                    <button href="https://www.smk.dev/blog/" className='bg-[#271b88] text-white py-2 px-4 rounded-full' >Liat Semua Artikel</button>
                </a>
            </div>
        </div>
    </div>
  )
}

export default Artikel