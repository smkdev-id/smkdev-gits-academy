import React, { FC } from "react";

interface Props {
  title: string;
  desc: string;
  color: string;
  svg: React.ReactNode;
}

const WhyUsCard: FC<Props> = ({ color, desc, svg, title }) => {
  return (
    <div className="flow-x-visible relative flex h-[80%] w-full items-center overflow-visible rounded-2xl bg-white shadow-xl">
      <div
        style={{ backgroundColor: color }}
        className="absolute left-0 top-1/2 h-fit w-fit -translate-x-1/2 -translate-y-1/2 transform overflow-visible rounded-xl p-5"
      >
        {svg}
      </div>
      <div className="flex flex-col gap-y-3 p-8 pl-20">
        <h1 className="text-2xl font-semibold capitalize text-primary">
          {title}
        </h1>
        <p className="text-base">{desc}</p>
      </div>
    </div>
  );
};

export default WhyUsCard;
