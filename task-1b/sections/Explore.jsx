"use client";
import { useState } from "react";
import { motion } from "framer-motion";

import styles from "../styles";
import { fadeIn, straggerContainer } from "../utils/motion";
import { ExploreCard, TitleText, TypingText } from "../components";
import { expertMentor } from "../constants";

const Explore = () => {
  const [active, setActive] = useState("mentor-2");
  return (
    <section className={`${styles.paddings}`} id="mentor">
      <motion.div
        variants={straggerContainer}
        initial="hidden"
        whileInView="show"
        viewport={{ once: false, amount: 0.25 }}
        className={`${styles.innerWidth} max-auto flex flex-col`}
      >
        <TypingText title="| Expert Mentor" textStyles="text-center" />
        <TitleText
          title={
            <>
              SMKDEV menghadirkan expert <br className="md:block hidden" />
              dari industri untuk mendampingi proses belajarmu
            </>
          }
          textStyles="text-center"
        />
        <div className="mt-[50px] flex lg:flex-row flex-col min-h-[70vh] gap-5 ">
          {expertMentor.map((expert, index) => (
            <ExploreCard
              key={expert.id}
              {...expert}
              index={index}
              active={active}
              handleClick={setActive}
            />
          ))}
        </div>
      </motion.div>
    </section>
  );
};

export default Explore;
