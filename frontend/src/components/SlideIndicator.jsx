import React from "react";

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

export default SlideIndicator;
