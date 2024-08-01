"use client";
import { useState } from "react";
import { motion } from "framer-motion";

import styles from "../styles";
import { fadeIn, straggerContainer, planetVariants } from "../utils/motion";
import { StartSteps, TitleText, TypingText } from "../components";
import { Talenta } from "../constants";

const GetStarted = () => (
  <section className={`${styles.paddings} relative z-10`}>
    <motion.div
      variants={straggerContainer}
      initial="hidden"
      whileInView="show"
      viewport={{ once: "false", amount: 0.25 }}
      className={`${styles.innerWidth} mx-auto flex lg:flex-row flex-col gap-8`}
    >
      <motion.div
        variants={planetVariants("left")}
        className={`flex-1 ${styles.flexCenter}`}
      >
        <img
          src="/Professional.png"
          alt="Professional"
          className="w-[90%] h-[90%] object-contain"
        />
      </motion.div>
      <motion.div
        variants={fadeIn("left", "tween", 0.2, 1)}
        className="flex-[0.75] flex justify-center flex-col "
      >
        <TypingText title="|Talenta Digital SMKDEV" />
        <TitleText
          title={
            <>
              Langkah Menjad
              <span className="text-[#ebb000]">Talenta Digital</span>
              <span className="text-[#00bd00]">Bersama</span>
              SMKDEV
            </>
          }
        />
        <div className="mt-[31px] flex flex-col max-w-[370px] gap-[24px]">
          {Talenta.map((features, index) => (
            <StartSteps key={features} number={index + 1} text={features} />
          ))}
        </div>
      </motion.div>
    </motion.div>
  </section>
);

export default GetStarted;
