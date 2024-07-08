import HeroImage from '../assets/hero-image.jpg'

const Hero = () => {
    return (
        <>
            <div className='flex flex-col md:flex-row justify-around mt-6 items-center w-screen gap-4 p-10'>
                <div>
                    <h1 className='md:text-5xl text-3xl font-bold leading-tight'>Siapkan Karir Digital Unggulan Anda</h1>
                    <h1 className='font-medium text-xs md:text-[17px] mt-4 leading-tight'>Ikuti pelatihan langsung dengan profesional berpengalaman dan kembangkan keterampilan dengan kurikulum berbasis proyek yang mendalam</h1>
                    <img className='h-[290px] mt-4 rounded-3xl md:hidden' src={HeroImage}/>
                    <a className='text-[16px] font-medium px-4 py-3 bg-black md:inline-block text-white mt-3 text-center rounded-xl block' href='#'>Get Started</a>
                </div>
                <img className='h-[420px] rounded-3xl hidden md:block' src={HeroImage}/>
            </div>
        </>
    )
}

export default Hero