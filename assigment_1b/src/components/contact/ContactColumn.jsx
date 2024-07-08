import React from "react";

const ContactColumn = ({ title, type, value, setWhat, placeholder }) => {
  return (
    <div className="mb-5">
      <label htmlFor={title} className="font-bold text-black dark:text-white text-lg mb-3">
        {title}
      </label>
      <br />
      <input
        placeholder={placeholder}
        type={type}
        id={title}
        value={value}
        onChange={(e) => setWhat(e.target.value)}
        className="text-black dark:text-white rounded-lg w-full border-solid border-2 border-gray-600 bg-transparent py-1 px-3 focus:outline-none focus:outline-primary focus:border-2 focus:ring-primary focus:outline-2 placeholder-gray-500 dark:placeholder-gray-400"
      />
    </div>
  );
};

export default ContactColumn;
