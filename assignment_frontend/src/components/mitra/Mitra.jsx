// Mitra.jsx
import React from "react";

const Mitra = ({ image, optional }) => {
  return (
    <div className={`w-32 h-32 md:w-40 md:h-40 my-4 mx-6 flex items-center justify-center transform transition-transform duration-300 hover:scale-110 `}>
      <img
        src={image}
        alt="mitragits"
        className={`w-full h-full object-contain hover:shadow-lg ${optional ? "logo" : ""}`}
      />
    </div>
  );
};

export default Mitra;
