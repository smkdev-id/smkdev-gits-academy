import { useEffect } from 'react'
import Navbar from './components/Navbar'
import Hero from './components/Hero';

function App() {

  useEffect(() => {
    document.title = 'SMKDEV • Creating High-Caliber Digital Talent';
  }, []);

  return (
    <>
      <Navbar/>
      <div className='pt-16'>
        <Hero/>
      </div>
    </>
  )
}

export default App
