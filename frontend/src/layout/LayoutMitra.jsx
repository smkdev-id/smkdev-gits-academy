// LayoutMitra.jsx
import React, { useState, useEffect } from "react";
import Spacer from "../components/Spacer";
import { FaIndustry, FaHandshake, FaBuilding, FaRegHandshake } from 'react-icons/fa';
import SlideIndicator from "../components/SlideIndicator";
import Mitra from "../components/mitra/Mitra";
import dataMitra from "../data/dataMitra"; // Import dataMitra

const LayoutMitra = () => {
  const padding = "pt-[45px] pb-[75px] md:pt-[90px] md:pb-[135px] px-[30px] md:px-[135px]";
  const bgLight = "bg-gradient-to-r from-bgHeroLight-start via-bgHeroLight to-bgHeroLight-end";
  const bgDark = "dark:from-gray-900 dark:via-gray-800 dark:to-black";
  
  const [currentIndex, setCurrentIndex] = useState(0);
  const [isMediumScreen, setIsMediumScreen] = useState(window.innerWidth >= 768);

  useEffect(() => {
    const handleResize = () => {
      setIsMediumScreen(window.innerWidth >= 768);
    };
    window.addEventListener("resize", handleResize);
    return () => {
      window.removeEventListener("resize", handleResize);
    };
  }, []);

  useEffect(() => {
    const interval = setInterval(() => {
      nextSlide();
    }, 5000);

    return () => {
      clearInterval(interval);
    };
  }, [isMediumScreen, currentIndex]);

  const slidesToShow = isMediumScreen ? 3 : 1; // Menentukan jumlah slide berdasarkan ukuran layar
  const totalSlides = Math.ceil(dataMitra.length / slidesToShow);

  const nextSlide = () => {
    setCurrentIndex((prevIndex) => (prevIndex + 1) % totalSlides);
  };

  const prevSlide = () => {
    setCurrentIndex((prevIndex) => (prevIndex - 1 + totalSlides) % totalSlides);
  };

  const startIndex = currentIndex * slidesToShow;
  const displayedImages = dataMitra.slice(startIndex, startIndex + slidesToShow);

  return (
    <section className={`${padding} ${bgLight} ${bgDark} text-black dark:text-white`}>
      <div className="flex flex-col items-center">
        <FaIndustry size={50} className="text-blue-500 dark:text-gray-500 mb-4" />
        <h1 className="text-4xl font-bold text-center">Mitra Industri</h1>
        <div className="w-24 h-1 bg-blue-500 dark:bg-gray-500 mt-4 mb-10"></div>
      </div>

      <div className="w-full flex justify-between items-center px-4">
        <button onClick={prevSlide} className="text-black dark:text-white text-2xl">
          &lt;
        </button>
        <div className="flex gap-4 overflow-hidden justify-center">
          {displayedImages.map((image, index) => (
            <Mitra key={index} image={image.src} optional={image.optional} />
          ))}
        </div>
        <button onClick={nextSlide} className="text-black dark:text-white text-2xl">
          &gt;
        </button>
      </div>

      <Spacer height={"h-10"} />

      <div className="flex justify-center space-x-10">
        <FaHandshake size={40} className="text-blue-500 dark:text-gray-500" />
        <FaBuilding size={40} className="text-blue-500 dark:text-gray-500" />
        <FaRegHandshake size={40} className="text-blue-500 dark:text-gray-500" />
      </div>

      <SlideIndicator currentIndex={currentIndex} totalSlides={totalSlides} /> {/* Menambahkan indikator slide */}
    </section>
  );
};

export default LayoutMitra;
