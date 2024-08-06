import React from "react";
import dataTalenta from "../../data/dataTalenta";
import SVGTalenta from "../SVG/SVGTalenta";

const ComponentTalenta = () => {
  return (
    <>
      <ul className="timeline timeline-vertical">
        {dataTalenta.map((text, index) => (
          <li key={index}>
            {index !== 0 && (
              <hr className="bg-primary dark:bg-primary-dark h-10 my-5" />
            )}
            <div
              className={`timeline-${
                index % 2 === 0 ? "start" : "end"
              } timeline-box bg-blue-900 dark:bg-blue-700 text-white dark:text-gray-300 p-2 rounded-md`}
            >
              {text}
            </div>
            <div className="timeline-middle">
              {index < 5 && <SVGTalenta />}
            </div>
            {index !== dataTalenta.length - 1 && (
              <hr className="bg-primary dark:bg-primary-dark h-10 my-5" />
            )}
          </li>
        ))}
      </ul>
    </>
  );
};

export default ComponentTalenta;
