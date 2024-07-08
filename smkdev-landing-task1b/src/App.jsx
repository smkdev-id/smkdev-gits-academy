import { useEffect } from 'react'
import Navbar from './components/Navbar'
import Hero from './components/Hero';
import Client from './components/Client';

function App() {

  useEffect(() => {
    document.title = 'SMKDEV • Creating High-Caliber Digital Talent';
  }, []);

  return (
    <>
      <Navbar/>
      <div className='pt-16'>
        <Hero/>
        <Client/>
      </div>
    </>
  )
}

export default App
