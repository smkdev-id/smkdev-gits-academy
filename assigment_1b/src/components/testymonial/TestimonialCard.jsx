import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faQuoteLeft, faQuoteRight } from "@fortawesome/free-solid-svg-icons";
import Spacer from "../Spacer";

const TestimonialCard = ({ testimonial }) => (
  <div className="w-full max-w-md bg-gradient-to-r from-gray-200 via-gray-100 to-gray-300 dark:from-gray-800 dark:via-gray-700 dark:to-gray-900 rounded-xl p-10 text-center m-2 shadow-lg transition-transform transform hover:scale-105">
    <div className="backdrop-filter backdrop-blur-lg bg-opacity-50 bg-white dark:bg-black dark:bg-opacity-20 p-6 rounded-xl">
      <div className="flex justify-center mb-4">
        <FontAwesomeIcon icon={faQuoteLeft} className="text-gray-500 dark:text-gray-400 text-3xl" />
      </div>
      <p className="italic text-gray-800 dark:text-gray-200">{testimonial.description}</p>
      <div className="flex justify-center mt-4">
        <FontAwesomeIcon icon={faQuoteRight} className="text-gray-500 dark:text-gray-400 text-3xl" />
      </div>

      <Spacer height={"h-10"} />
      
      <h1 className="text-xl font-bold text-gray-900 dark:text-white text-start">{testimonial.name}</h1>
      <h1 className="text-gray-600 dark:text-gray-400 text-start">{testimonial.activity}</h1>
    </div>
  </div>
);

export default TestimonialCard;
