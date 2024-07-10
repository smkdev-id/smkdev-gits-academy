import React, { FC } from "react";

interface Props {
  children: React.ReactNode;
  className?: string;
}

const MediaQuery: FC<Props> = ({ className, children }) => {
  return (
    <div className={`${className} w-[95%] md:w-[80%] xl:w-[60%]`}>
      {children}
    </div>
  );
};

export default MediaQuery;
