import React, { useEffect, useRef } from "react";
import Typed from "typed.js";
import Spacer from "../components/Spacer";

const LayoutHero = () => {
  const typedElement = useRef(null);

  useEffect(() => {
    const typed = new Typed(typedElement.current, {
      strings: ["Global", "Masa Depan Indonesia", "Digitalent"],
      typeSpeed: 100,
      backSpeed: 30,
      loop: true,
      backDelay: 1500,
    });

    return () => {
      typed.destroy();
    };
  }, []);

  const padding =
    "pt-[45px] pb-[75px] md:pt-[135px] md:pb-[135px] px-[30px] md:px-[135px]";

  return (
    <>
      <section
        className={`min-h-screen ${padding} font-poppins md:flex md:justify-between items-center bg-gradient-to-r from-bgHeroLight-start via-bgHeroLight to-bgHeroLight-end dark:from-gray-900 dark:via-gray-800 dark:to-black`}
      >
        <div className="w-full h-[400px] bg-orange-100 rounded-xl md:w-[45%] overflow-hidden shadow-lg">
          <img
            src="/src/assets/man.png"
            alt="man"
            className="w-full h-full object-cover rounded-xl"
          />
        </div>
        <div className="mt-6 md:mt-0 md:w-[45%] text-center md:text-left">
          <h1 className="font-extrabold text-4xl md:text-5xl leading-tight text-black dark:text-white">
            Jadilah Talenta <br /> Digital <br />
            <span className="text-blue-600">
              <span ref={typedElement}></span>
            </span>
          </h1>

          <Spacer height={"h-4"} />

          <p className="text-lg md:text-xl mt-4 text-gray-700 dark:text-gray-300">
            Belajar langsung dengan expert dari industri dengan kurikulum
            komprehensif berbasis project-based learning
          </p>

          <Spacer height={"h-10 md:h-8"} />

          <button className="hover:cursor-pointer hover:bg-blue-600 transition-all duration-300 font-bold text-xl px-8 py-3 border-2 border-solid rounded-full mt-6 text-black dark:text-white hover:text-white border-blue-600">
            More
          </button>
        </div>
      </section>
    </>
  );
};

export default LayoutHero;
