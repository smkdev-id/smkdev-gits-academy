import React from 'react';
import './App.css';
import Header from './components/Header';

import Footer from './components/Footer';
import HeroSection from './components/HeroSection';
import Features from './components/Features';
import Testimonials from './components/Testimonials';
import CallToAction from './components/CallToAction';
import Articles from './components/Article';

function App() {
  return (
    <div className="App">
      <Header />
      
      <HeroSection />
      <Features />
      <Testimonials />
      <Articles />
      <CallToAction />
      <Footer />
    </div>
  );
}

export default App;

