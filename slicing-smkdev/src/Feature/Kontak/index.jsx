import React from 'react'
import smkdev_logo from '../../assets/smkdev.png'
import yt from '../../assets/icon_social/youtube.png'
import linkedin from '../../assets/icon_social/linkedin.png'
import ig from '../../assets/icon_social/instagram.png'
import tt from '../../assets/icon_social/tik-tok.png'
import tele from '../../assets/icon_social/telegram.png'
import wa from '../../assets/icon_social/whatsapp.png'
import TextArea from 'antd/es/input/TextArea'

const Kontak = () => {
  return (
    <div>
        <div className='bg-[#f2f2f8]'>
            {/* form */}
            <div className='flex justify-center py-12 px-20'>
                <div className='flex flex-col text-left gap-4'>
                    <h1 className='text-3xl font-semibold'>Kontak SMKDEV</h1>
                    <p className='text-lg'>Berapa investasi untuk menggunakan SMKDEV? Anda bisa mulai dengan <br /><span className='font-bold'>GRATIS.</span></p>
                    <p className='text-lg'>Silahkan isi form disebelah untuk informasi lebih lanjut</p>
                </div>
                <div className='mx-12'>
                    <form action="" className='flex flex-col gap-4'>
                        <div>
                            <label for="name">Nama Anda<span className='text-red-500'>&#42;</span></label>
                            <input type="text" name="name" id="name" className='mt-1 block w-[500px] px-3 py-2 bg-[#cccccc] border border-slate-300 rounded-md text-sm shadow-sm placeholder-slate-400 focus:outline-none focus:border-sky-500 focus:ring-1 focus:ring-sky-500' required/>
                        </div>
                        <div>
                            <label for="email">Nama Email<span className='text-red-500'>&#42;</span></label>
                            <input type="email" name="email" id="email" className='mt-1 block w-[500px] px-3 py-2 bg-[#cccccc] border border-slate-300 rounded-md text-sm shadow-sm placeholder-slate-400 focus:outline-none focus:border-sky-500 focus:ring-1 focus:ring-sky-500' required/>
                        </div>
                        <div>
                            <label for="number">Nomor WhatsApp<span className='text-red-500'>&#42;</span></label>
                            <input type="text" name="number" id="number" className='mt-1 block w-[500px] px-3 py-2 bg-[#cccccc] border border-slate-300 rounded-md text-sm shadow-sm placeholder-slate-400 focus:outline-none focus:border-sky-500 focus:ring-1 focus:ring-sky-500' required placeholder='0822-9922-8847'/>
                        </div>
                        <div>
                            <label for="instansi">Nama Satuan Pendidikan/Instansi <span className='text-red-500'>&#42;</span></label>
                            <input type="text" name="instansi" id="instansi" className='mt-1 block w-[500px] px-3 py-2 bg-[#cccccc] border border-slate-300 rounded-md text-sm shadow-sm placeholder-slate-400 focus:outline-none focus:border-sky-500 focus:ring-1 focus:ring-sky-500' required/>
                        </div>
                        <div className='flex flex-col'>
                            <label for="pesan">Pesan<span className='text-red-500'>&#42;</span></label>
                            <TextArea rows={4} cols={50} name="pesan" id="pesan" className='mt-1 block w-[500px] h-[400px] px-3 py-2 bg-[#cccccc] hover:bg-[#cccccc]  border border-slate-300 rounded-md text-sm shadow-sm placeholder-slate-400 focus:outline-none focus:border-sky-500 focus:ring-1 focus:ring-sky-500 focus:text-left' required/>
                        </div>
                        <div className='flex gap-2 bg-slate-300 p-2 rounded-lg w-[180px]'>
                            <input type="checkbox" name="check" id="check"/>
                            <label for="check" className='text-sm'>I Am Human</label>
                        </div>
                        <div className=''>
                            <button className='bg-sky-500 text-white px-4 py-2 rounded-md'>Send</button>
                        </div>
                    </form>
                </div>
            </div>
            {/* follow us */}
            <div className='flex container mx-auto flex-col text-left gap-4 px-44 py-4'>
                <img src={smkdev_logo} alt="smkdev" className='w-[120px]'/>
                <p className='text-lg font-semibold'>Creating High-Caliber Digital Talent</p>
                <p className='text-lg'>Summarecon Bandung, Jl. Magna Barat Blok MD No.02, Rancabolang, Kec. Gedebage, Kota Bandung, Jawa Barat 40295</p>
                <div className='flex flex-col'>
                    <p className='text-lg font-semibold'>follow us</p>
                    <div className='flex gap-4'>
                        <a href="https://www.youtube.com/@smkdev" target="_blank">
                            <img src={yt} alt="youtube" className='w-[30px] h-[30px] cursor-pointer'/>
                        </a>
                        <a href="https://www.linkedin.com/company/smkdev/" target="_blank">
                            <img src={linkedin} alt="linkedin" className='w-[30px] h-[30px] cursor-pointer'/>
                        </a>
                        <a href="https://www.instagram.com/smkdev.offic" target="_blank">
                            <img src={ig} alt="instagram" className='w-[30px] h-[30px] cursor-pointer'/>
                        </a> 
                        <a href="https://www.tiktok.com/@smkdev.official" target="_blank">
                            <img src={tt} alt="tiktok" className='w-[30px] h-[30px] cursor-pointer'/>
                        </a>
                        <a href="https://t.me/+wPXVCS29UKgwNmJl" target="_blank">
                            <img src={tele} alt="telegram" className='w-[30px] h-[30px] cursor-pointer'/>
                        </a>
                        <a href="https://chat.whatsapp.com/D5gVS4FbxSc2ij1zbZVxAV" target="_blank">    
                            <img src={wa} alt="whatsapp" className='w-[30px] h-[30px] cursor-pointer'/>
                        </a>
                    </div>
                </div>
            </div>
        </div>
    </div>
  )
}

export default Kontak