import React from "react";
import Spacer from "../components/Spacer";
import MentorCard from "../components/mentor/MentorCard";
import dataMentor from "../data/dataMentor";

const LayoutMentor = () => {
  const padding =
    "pt-[45px] pb-[75px] md:pt-[0px] md:pb-[135px] px-[30px] md:px-[135px]";
  const bgLight = "bg-gradient-to-r from-bgHeroLight-start via-bgHeroLight to-bgHeroLight-end";
  const bgDark = "dark:from-gray-900 dark:via-gray-800 dark:to-black";

  return (
    <>
      <section className={`font-poppins text-black dark:text-white ${padding} ${bgLight} ${bgDark}`}>
        <h1 className="text-4xl text-center font-bold">Expert Mentor</h1>

        <Spacer height={"h-4"} />

        <p className="text-center text-black dark:text-white">
          SMKDEV menghadirkan expert dari industri untuk mendampingi proses
          belajarmu.
        </p>

        <Spacer height={"h-10"} />

        <div className="md:flex md:flex-wrap md:justify-around">
          {dataMentor.map((mentor, index) => (
            <MentorCard
              key={index}
              image={mentor.image}
              name={mentor.name}
              position={mentor.position}
              company={mentor.company}
              description={mentor.description}
            />
          ))}
        </div>
      </section>
    </>
  );
};

export default LayoutMentor;
