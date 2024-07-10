import { useEffect } from 'react'
import Navbar from './components/Navbar'
import Hero from './components/Hero';
import Client from './components/Client';
import LearnOrientation from './components/LearnOrientation';
import Talent from './components/Talent';
import Testimonial from './components/Testimonial';
import Mentor from './components/Mentor';
import Article from './components/Article';
import Contact from "./components/Contact";
import Footer from './components/Footer';

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
        <Mentor/>
        <Article/>
        <Contact/>
        <Footer/>
      </div>
    </>
  )
}

export default App
