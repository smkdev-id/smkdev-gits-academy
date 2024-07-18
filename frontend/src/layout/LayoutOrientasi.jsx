import React from "react";
import Spacer from "../components/Spacer";
import dataOrientasi from "../data/dataOrientasi"; // Import dataOrientasi

const LayoutOrientasi = () => {
  const padding =
    "pt-[45px] pb-[75px] md:pt-[0px] md:pb-[135px] px-[30px] md:px-[135px]";
  const bgLight = "bg-gradient-to-r from-bgHeroLight-start via-bgHeroLight to-bgHeroLight-end";
  const bgDark = "dark:from-gray-900 dark:via-gray-800 dark:to-black";

  return (
    <>
      <section className={`${padding} ${bgLight} ${bgDark} text-black dark:text-white font-poppins`}>
        <h1 className="font-bold text-4xl text-center text-black dark:text-white">
          Orientasi Belajar SMKDEV
        </h1>

        <Spacer height={"h-10"} />

        <div className="md:flex md:justify-around md:flex-wrap">
          {dataOrientasi.map((item, index) => (
            <Card
              key={index}
              image={item.image}
              title={item.title}
              desc={item.desc}
            />
          ))}
        </div>
      </section>
    </>
  );
};

export default LayoutOrientasi;

const Card = ({ image, title, desc }) => {
  return (
    <>
      <div
        className="my-8 w-full md:w-[45%] py-10 rounded-xl hover:border-2 md:hover:border-4 hover:border-solid hover:border-blue-700 dark:hover:border-blue-300 transition-all duration-900"
        style={{ boxShadow: "0 10px 30px rgba(100, 100, 100, 0.20)" }}
      >
        <div className="w-[30%] mx-auto h-32 ">
          <img src={image} alt="book" className="w-full h-full" />
        </div>

        <div className="h-4"></div>

        <div className="px-12  ">
          <h1 className="text-center text-2xl font-bold text-black dark:text-white">{title}</h1>

          <div className="h-2"></div>

          <h2 className="text-center text-black dark:text-white">{desc}</h2>
        </div>
      </div>
    </>
  );
};
