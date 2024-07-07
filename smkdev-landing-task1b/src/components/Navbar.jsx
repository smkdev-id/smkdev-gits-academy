import { useState } from 'react'
import LogoSmkDev from '../assets/logo-smkdev.png'

const Navbar = () => {
    const [isMenuOpen, setIsMenuOpen] = useState(false)

    const toggleMenu = () => {
        setIsMenuOpen(!isMenuOpen)
    }

    return (
        <>
            <nav className="bg-white p-5 fixed w-full text-black z-50">
                <div className='flex justify-between items-center'>
                    <div>
                        <img src={LogoSmkDev} className='h-5' alt="SMK Dev Logo" />
                    </div>
                    <button className='md:hidden text-2xl' onClick={toggleMenu}>
                        {isMenuOpen ? 'X' : '≡'}
                    </button>
                    <div className='hidden md:block'>
                        <ul className='flex gap-8 font-medium'>
                            <li><a className='text-xl' href="#">Learn</a></li>
                            <li><a className='text-xl' href="#">Community</a></li>
                            <li><a className='text-xl' href="#">Blog</a></li>
                        </ul>
                    </div>
                    <div className='hidden md:block'>
                        <a className='px-5 py-3 bg-slate-300 rounded-md font-medium mr-2' href="#">Sign Up</a>
                        <a className='px-5 py-3 bg-slate-300 rounded-md font-medium' href="#">Sign In</a>
                    </div>
                </div>
            </nav>
            {isMenuOpen && (
                <div className='fixed inset-0 bg-white z-40 pt-20 px-5 md:hidden'>
                    <ul className='space-y-4'>
                        <li><a href="#" className="text-xl flex justify-between items-center">Learn <span>›</span></a></li>
                        <li><a href="#" className="text-xl flex justify-between items-center">Community <span>›</span></a></li>
                        <li><a href="#" className="text-xl flex justify-between items-center">Blog <span>›</span></a></li>
                    </ul>
                    <a href='#' className="mt-8 w-full block text-center bg-slate-300 py-3 rounded-md font-medium">
                        Sign In
                    </a>
                    <a href='#' className="mt-4 w-full block text-center bg-slate-300 py-3 rounded-md font-medium">
                        Sign Up
                    </a>
                </div>
            )}
        </>
    )
}

export default Navbar