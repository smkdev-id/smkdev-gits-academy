import React from "react";
import ComponentTalenta from "../components/talenta/ComponentTalenta";
import Spacer from "../components/Spacer";

const LayoutTalenta = () => {
  const padding =
    "pt-[45px] pb-[75px] md:pt-[0px] md:pb-[135px] px-[30px] md:px-[135px]";
  const bgLight = "bg-gradient-to-r from-bgHeroLight-start via-bgHeroLight to-bgHeroLight-end";
  const bgDark = "dark:from-gray-900 dark:via-gray-800 dark:to-black";

  return (
    <>
      <section className={`${bgLight} ${bgDark} ${padding} font-poppins text-black dark:text-white`}>
        <h1 className="text-center font-bold text-4xl text-black dark:text-white">
          Talenta Digital SMKDEV
        </h1>
        <Spacer height={"h-10"} />

        <ComponentTalenta />
      </section>
    </>
  );
};

export default LayoutTalenta;
