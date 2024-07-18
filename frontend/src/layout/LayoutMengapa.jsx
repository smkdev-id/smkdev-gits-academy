import React from "react";
import Spacer from "../components/Spacer";
import MengapaAccordion from "../components/mengapa/MengapaAccordion";
import dataMengapa from "../data/dataMengapa"; // Import dataMengapa

const LayoutMengapa = () => {
  const padding =
    "pt-[45px] pb-[75px] md:pt-[0px] md:pb-[135px] px-[30px] md:px-[135px]";
  const bgLight = "bg-gradient-to-r from-bgHeroLight-start via-bgHeroLight to-bgHeroLight-end";
  const bgDark = "dark:from-gray-900 dark:via-gray-800 dark:to-black";

  return (
    <>
      <section className={`${padding} ${bgLight} ${bgDark} text-black dark:text-white font-poppins`}>
        <div className="md:w-[60%] md:mx-auto">
          <h1 className="text-4xl font-bold text-center text-black dark:text-white">
            MENGAPA HARUS MEMILIH BELAJAR DI SMKDEV
          </h1>

          <Spacer height={"h-4"} />

          <p className="text-justify md:text-center text-black dark:text-white">
            Pengalaman belajarmu jauh lebih menyenangkan dengan perpaduan
            belajar teori dan capstone project yang dapat menjadi portofolio
            pribadimu. Tidak lupa, Kamu juga dapat belajar langsung dari expert
            mentor dari dunia industri.
          </p>
        </div>

        <Spacer height={"h-10"} />

        {dataMengapa.map((item, index) => (
          <MengapaAccordion
            key={index}
            imgSrc={item.imgSrc}
            title={item.title}
            content={item.content}
          />
        ))}

        <Spacer height={"h-10"} />
      </section>
    </>
  );
};

export default LayoutMengapa;
