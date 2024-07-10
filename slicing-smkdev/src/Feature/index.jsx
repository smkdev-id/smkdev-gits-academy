import React from 'react'
import Navbar from './Navbar'
import Hero from './Hero'
import Mitra from './Mitra'
import Orientasi from './Orientasi_Belajar'
import WhySmkDev from './WhySmkDev'
import Testimoni from './Testimoni'
import Mentor from './Mentor'
import Artikel from './Artikel'
import Kontak from './Kontak'
import Footer from './Footer'

const Main = () => {
  return (
    <div>
    <Navbar/>
      <Hero/>
      <Mitra/>
      <Orientasi/>
      <WhySmkDev/>
      <Testimoni/>
      <Mentor/>
      <Artikel/>
      <Kontak/>
      <Footer/>
      </div>
  )
}

export default Main