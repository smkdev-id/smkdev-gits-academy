import React from "react";
import IconInstagram from "../components/icons/IconInstagram";
import IconGit from "../components/icons/IconGit";
import IconTele from "../components/icons/IconTele";
import IconWa from "../components/icons/IconWa";

const LayoutFooter = () => {
  return (
    <footer className="px-10 bg-white dark:bg-black w-full h-full py-10 text-black dark:text-white">
      <div className="lg:flex justify-between">
        <div className="py-3 lg:w-[50%]">
          <img
            src="/src/assets/logosmk.png"
            alt="Logo I-Secret"
            className="lg:mx-0 w-36 h-36 mx-auto my-2 object-contain"
          />
          <p className="lg:text-left text-black dark:text-white font-normal text-lg mx-auto text-center">
            Summarecon Bandung, Jl. Magna Barat Blok MD No.02,
          </p>
          <p className="lg:text-left text-black dark:text-white font-normal text-lg mx-auto text-center">
            Rancabolang, Kec. Gedebage, Kota Bandung, Jawa Barat 40295
          </p>
        </div>
        <div className="py-3 hidden lg:block w-20">
          <p className="font-bold text-black dark:text-white">General</p>
          <ul className="text-black dark:text-white">
            <li className="hover:underline">
              <a href="#">Home</a>
            </li>
            <li className="hover:underline">
              <a href="#">About Us</a>
            </li>
            <li className="hover:underline">
              <a href="#">Events</a>
            </li>
            <li className="hover:underline">
              <a href="#">Projects</a>
            </li>
            <li className="hover:underline">
              <a href="#">Blog</a>
            </li>
          </ul>
        </div>
        <div className="py-5 lg:py-3 lg:mx-2">
          <div className="my-10 px-5 md:my-0 md:px-0 lg:block">
            <p className="text-black dark:text-white text-lg font-bold">
              Join our mailing list!
            </p>
            <p className="text-black dark:text-white text-lg">
              Always get updates about our events and posts
            </p>
            <input
              type="email"
              placeholder="Your Email"
              className="bg-gray-200 dark:bg-black placeholder:text-black dark:placeholder:text-white text-sm w-52 px-2 py-1 my-3 border-slate-500 border-2 shadow-sm rounded-lg text-black dark:text-white"
            />
            <button className="absolute bg-blue-600 dark:bg-blue-800 text-white hover:bg-blue-800 dark:hover:bg-blue-900 font-medium py-[3px] px-6 rounded-md mt-[13px] -ml-2">
              Submit
            </button>
          </div>
          <p className="text-black dark:text-white text-lg font-bold text-center lg:text-left">
            Contact us (Email):
          </p>
          <a href="#">
            <p className="text-black dark:text-white text-lg underline text-center lg:text-left">
              smkdev@gmail.com
            </p>
          </a>
        </div>
        <div className="flex lg:hidden gap-8 justify-center py-5">
          <a href="#">
            <IconInstagram
              fill="currentColor"
              hoverFill="blue"
              className="w-6 h-6 md:w-10 md:h-10"
            />
          </a>
          <a href="#">
            <IconWa
              fill="currentColor"
              hoverFill="blue"
              className="w-6 h-6 md:w-10 md:h-10"
            />
          </a>
          <a href="#">
            <IconGit
              fill="currentColor"
              hoverFill="blue"
              className="w-6 h-6 md:w-10 md:h-10"
            />
          </a>
          <a href="#">
            <IconTele
              fill="currentColor"
              hoverFill="blue"
              className="w-6 h-6 md:w-10 md:h-10"
            />
          </a>
        </div>
        <div className="hidden lg:grid max-w-10 gap-2 py-3 mx-2">
          <a href="#">
            <IconInstagram
              fill="currentColor"
              hoverFill="blue"
              className="w-20 h-8 md:w-10 md:h-10"
            />
          </a>
          <a href="#">
            <IconWa
              fill="currentColor"
              hoverFill="blue"
              className="w-20 h-8 md:w-10 md:h-10"
            />
          </a>
          <a href="#">
            <IconGit
              fill="currentColor"
              hoverFill="blue"
              className="w-20 h-8 md:w-10 md:h-10"
            />
          </a>
          <a href="#">
            <IconTele
              fill="currentColor"
              hoverFill="blue"
              className="w-20 h-8 md:w-10 md:h-10"
            />
          </a>
        </div>
      </div>
      <p className="text-black dark:text-white text-xl text-center pt-32">
        &copy; 2024 SMKDEV – PT Eureka Merdeka Indonesia. All Rights Reserved.
      </p>
    </footer>
  );
};

export default LayoutFooter;
