import React from "react";

const CountdownTime = ({ days, hours, minutes, seconds }) => {
  return (
    <div className="grid grid-flow-col gap-5 text-center auto-cols-max text-black dark:text-white">
      <div className="flex flex-col">
        <span className="countdown font-mono text-5xl md:text-8xl">
          <span style={{ "--value": days }}></span>
        </span>
        <span className="text-sm md:text-lg">days</span>
      </div>
      <div className="flex flex-col">
        <span className="countdown font-mono text-5xl md:text-8xl">
          <span style={{ "--value": hours }}></span>
        </span>
        <span className="text-sm md:text-lg">hours</span>
      </div>
      <div className="flex flex-col">
        <span className="countdown font-mono text-5xl md:text-8xl">
          <span style={{ "--value": minutes }}></span>
        </span>
        <span className="text-sm md:text-lg">min</span>
      </div>
      <div className="flex flex-col">
        <span className="countdown font-mono text-5xl md:text-8xl">
          <span style={{ "--value": seconds }}></span>
        </span>
        <span className="text-sm md:text-lg">sec</span>
      </div>
    </div>
  );
};

export default CountdownTime;
