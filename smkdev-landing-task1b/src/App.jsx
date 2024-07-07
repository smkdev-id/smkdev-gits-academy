import { useEffect, useState } from 'react'
import Navbar from './components/Navbar'
import Hero from './components/Hero';

function App() {

  useEffect(() => {
    document.title = 'SMKDEV • Creating High-Caliber Digital Talent';
  }, []);

  return (
    <>
      <Navbar/>
    </>
  )
}

export default App
