import { useEffect } from 'react'
import Navbar from './components/Navbar'
import Hero from './components/Hero';
import Client from './components/Client';
import LearnOrientation from './components/LearnOrientation';
import Talent from './components/Talent';
import Testimonial from './components/Testimonial';

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
        <LearnOrientation/>
        <Talent/>
        <Testimonial/>
      </div>
    </>
  )
}

export default App
