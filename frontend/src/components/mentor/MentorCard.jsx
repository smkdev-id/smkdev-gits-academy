import React, { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import Spacer from "../Spacer";
import IconInstagram from "../icons/IconInstagram";
import IconWa from "../icons/IconWa";
import IconGit from "../icons/IconGit";
import IconTele from "../icons/IconTele";

const MentorCard = ({ name, position, company, description, image }) => {
  const [isMore, setIsMore] = useState(false);

  const handleToggle = () => {
    setIsMore(!isMore);
  };

  return (
    <div
      className="w-full py-10 px-10 rounded-3xl my-10 md:w-[40%] md:my-5 bg-white dark:bg-gray-800 text-black dark:text-white shadow-lg dark:shadow-blue-900"
    >
      <div className="w-full flex justify-center rounded-full overflow-hidden">
        <div className=" w-36 h-36 md:w-52 md:h-52">
          <img
            src={image}
            alt="mentor"
            className=" w-full h-full rounded-full"
          />
        </div>
      </div>

      <Spacer height={"h-4"} />

      <div className="flex justify-center space-x-2 md:space-x-5">
        <a href="/" className="transition-all">
          <IconInstagram
            fill="currentColor"
            hoverFill="blue"
            className="w-6 h-6 md:w-10 md:h-10"
          />
        </a>
        <a href="/" className="transition-all">
          <IconWa
            fill="currentColor"
            hoverFill="blue"
            className="w-6 h-6 md:w-10 md:h-10"
          />
        </a>
        <a href="/" className="transition-all">
          <IconGit
            fill="currentColor"
            hoverFill="blue"
            className="w-6 h-6 md:w-10 md:h-10"
          />
        </a>
        <a href="/" className="transition-all">
          <IconTele
            fill="currentColor"
            hoverFill="blue"
            className="w-6 h-6 md:w-10 md:h-10"
          />
        </a>
      </div>

      <Spacer height={"h-4"} />

      <h1 className="font-bold text-xl text-center font-playwrite">{name}</h1>

      <Spacer height={"h-2"} />

      <div className="h-1 bg-black dark:bg-white w-full md:w-[50%] md:mx-auto"></div>

      <Spacer height={"h-2"} />

      <h1 className="font-bold text-xl text-center">{position}</h1>

      <Spacer height={"h-1"} />

      <h1 className="text-lg font-bold text-gray-700 dark:text-gray-300 text-center">
        {company}
      </h1>

      <Spacer height={"h-2"} />

      <AnimatePresence>
        {isMore && (
          <motion.div
            initial={{ opacity: 0, height: 0 }}
            animate={{ opacity: 1, height: "auto" }}
            exit={{ opacity: 0, height: 0 }}
            transition={{ duration: 0.3 }}
            className="mt-4 text-center text-black dark:text-gray-400"
          >
            <p>{description}</p>
            <Spacer height={"h-4"} />
          </motion.div>
        )}
      </AnimatePresence>

      <button
        onClick={handleToggle}
        className="font-bold text-lg text-center text-gray-500 dark:text-gray-400 focus:outline-none hover:text-blue-500 dark:hover:text-blue-300 hover:underline"
      >
        {isMore ? "Less" : "More"}
      </button>
    </div>
  );
};

export default MentorCard;
