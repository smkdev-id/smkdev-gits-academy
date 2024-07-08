import React, { useState, useEffect } from "react";
import Spacer from "../components/Spacer";
import TestimonialCard from "../components/testymonial/TestimonialCard";
import testimonials from "../data/dataTesty"; // Import testimonials
import { CSSTransition, TransitionGroup } from "react-transition-group";

const ArrowButton = ({ direction, onClick }) => {
  return (
    <button
      onClick={onClick}
      className={`absolute z-10 ${direction === 'left' ? 'left-4' : 'right-4'} transform -translate-y-1/2 top-1/2`}
      style={{ padding: '10px', backgroundColor: 'rgba(0, 0, 0, 0.5)', borderRadius: '50%' }}
    >
      <svg
        className="h-8 w-8 text-white"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          strokeLinecap="round"
          strokeLinejoin="round"
          strokeWidth="2"
          d={direction === 'left' ? 'M15 19l-7-7 7-7' : 'M9 5l7 7-7 7'}
        />
      </svg>
    </button>
  );
};

const SlideIndicator = ({ currentIndex, totalSlides }) => {
  return (
    <div className="flex justify-center mt-4">
      {Array.from({ length: totalSlides }).map((_, index) => (
        <div
          key={index}
          className={`h-3 w-3 rounded-full mx-1 ${
            index === currentIndex ? "bg-blue-500 dark:bg-gray-500" : "bg-gray-300 dark:bg-gray-700"
          }`}
        ></div>
      ))}
    </div>
  );
};

const LayoutTesty = () => {
  const [currentIndex, setCurrentIndex] = useState(0);
  const [isMediumScreen, setIsMediumScreen] = useState(window.innerWidth >= 768);
  const padding =
    "pt-[45px] pb-[75px] md:pt-[0px] md:pb-[135px] px-[30px] md:px-[135px]";
  const bgLight = "bg-gradient-to-r from-bgHeroLight-start via-bgHeroLight to-bgHeroLight-end";
  const bgDark = "dark:from-gray-900 dark:via-gray-800 dark:to-black";

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
      handleNextClick();
    }, 5000); 

    return () => {
      clearInterval(interval);
    };
  }, [isMediumScreen, currentIndex]); 

  const handlePrevClick = () => {
    setCurrentIndex((prevIndex) =>
      prevIndex === 0 ? (isMediumScreen ? Math.ceil(testimonials.length / 2) - 1 : testimonials.length - 1) : prevIndex - 1
    );
  };

  const handleNextClick = () => {
    setCurrentIndex((prevIndex) =>
      isMediumScreen
        ? (prevIndex + 1) % Math.ceil(testimonials.length / 2)
        : (prevIndex + 1) % testimonials.length
    );
  };

  const totalSlides = isMediumScreen ? Math.ceil(testimonials.length / 2) : testimonials.length;

  return (
    <>
      <section
        className={`${padding} font-poppins text-black dark:text-white ${bgLight} ${bgDark} relative`}
      >
        <h1 className="text-4xl font-bold text-center">Apa Kata Mereka?</h1>

        <Spacer height={"h-10"} />

        <div className="flex justify-center items-center relative slider-wrapper">
          <ArrowButton direction="left" onClick={handlePrevClick} />

          <div className="flex flex-col md:flex-row justify-center items-center w-full max-w-6xl slider-container">
            <TransitionGroup className="flex w-full">
              <CSSTransition
                key={currentIndex}
                timeout={500}
                classNames="fade"
              >
                <div className="slide flex flex-col md:flex-row justify-center items-center w-full">
                  <TestimonialCard testimonial={testimonials[currentIndex * (isMediumScreen ? 2 : 1)]} />
                  {isMediumScreen && currentIndex * 2 + 1 < testimonials.length && (
                    <TestimonialCard testimonial={testimonials[currentIndex * 2 + 1]} />
                  )}
                </div>
              </CSSTransition>
            </TransitionGroup>
          </div>

          <ArrowButton direction="right" onClick={handleNextClick} />
        </div>

        <SlideIndicator currentIndex={currentIndex} totalSlides={totalSlides} />
      </section>
    </>
  );
};

export default LayoutTesty;
