"use client";
import { motion } from "framer-motion";

import styles from "../styles";
import { fadeIn, straggerContainer, planetVariants } from "../utils/motion";
import { NewFeatures, TitleText, TypingText } from "../components";
import { Testimoni } from "../constants";

const WhatsNew = () => (
  <section className={`${styles.paddings} relative z-10`}>
    <motion.div
      variants={straggerContainer}
      initial="hidden"
      whileInView="show"
      viewport={{ once: "false", amount: 0.25 }}
      className={`${styles.innerWidth} mx-auto flex lg:flex-row flex-col gap-8`}
    >
      <motion.div
        variants={fadeIn("left", "tween", 0.2, 1)}
        className="flex-[0.75] flex justify-center flex-col "
      >
        <TypingText title="|Apa Kata Mereka?" />
        <TitleText
          title={
            <>
              Mereka telah merasakan{" "}
              <span className="text-[#00bd00]">serunya</span>{" "}
              <span className="text-[#ebb000]">belajar skill digital</span> di
              SMKDEV. Apakah kamu ingin seperti mereka?
            </>
          }
        />
        <div className="mt-[31px] flex flex-warp justify-berween gap-[24px]">
          {Testimoni.map((Testimoni) => (
            <NewFeatures key={Testimoni.name} {...Testimoni} />
          ))}
        </div>
      </motion.div>
    </motion.div>
  </section>
);

export default WhatsNew;
