import React from 'react';
import { useInView } from 'react-intersection-observer';
import 'tailwindcss/tailwind.css';
import logo from '../assets/bootcamp.png';  // Ensure this path is correct

const HeroSection = () => {
  const { ref, inView } = useInView({
    triggerOnce: true,
    threshold: 0.1,
  });

  return (
    <section
      ref={ref}
      className="bg-white py-20 flex flex-col items-center md:flex-row md:justify-between container mx-auto"
    >
      <div
        className={`transition-opacity duration-1000 ${
          inView ? 'opacity-100 animate-fadeIn' : 'opacity-0'
        }`}
      >
        <h1 className="text-4xl font-bold mb-4 text-customBlue font-customFont animate-fadeIn">
          Become The Future Digital Talent
        </h1>
        <p className="text-lg animate-fadeIn">
          Learn directly from industry experts with a comprehensive curriculum based on project-based learning.
        </p>
      </div>
      <div
        className={`transition-opacity duration-1000 delay-500 ${
          inView ? 'opacity-100 animate-fadeIn' : 'opacity-0'
        } md:w-1/2 flex justify-center`}
      >
        <img
          src={logo}
          alt="Learning"
          className="rounded-lg mt-10 md:mt-0"
          style={{ width: '350px', height: 'auto' }}
        />
      </div>
    </section>
  );
};

export default HeroSection;


