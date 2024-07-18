import React from "react";

const MengapaAccordion = ({ imgSrc, title, content }) => {
  return (
    <div className="collapse collapse-arrow bg-gray-200 dark:bg-gray-800 md:mb-3">
      <input type="radio" name="my-accordion-2" />
      <div className="collapse-title flex items-center text-xl font-medium text-black dark:text-white">
        <img src={imgSrc} alt="icon" className="w-16 h-16 mr-4" />
        <h1 className="text-2xl">{title}</h1>
      </div>
      <div className="collapse-content text-black dark:text-white">
        <p>{content}</p>
      </div>
    </div>
  );
};

export default MengapaAccordion;
