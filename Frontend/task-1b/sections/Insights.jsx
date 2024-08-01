"use client";
import { motion } from "framer-motion";

import styles from "../styles";
import { fadeIn, straggerContainer, planetVariants } from "../utils/motion";
import { InsightCard, TitleText, TypingText } from "../components";
import { Artikel } from "../constants";

const Insights = () => (
  <section className={`${styles.paddings} relative z-10`}>
    <motion.div
      variants={straggerContainer}
      initial="hidden"
      whileInView="show"
      viewport={{ once: "false", amount: 0.25 }}
      className={`${styles.innerWidth} mx-auto flex flex-col`}
    >
      <TypingText title="artikel" textStyles="text-center" />
      <TitleText title="Artikel Pilihan SMKDEV" textStyles="text-center" />
      <div className="mt-[50px] flex flex-col gap-[30px]">
        {Artikel.map((artikel, index) => (
          <InsightCard
            key={`artikel-${index}`}
            {...artikel}
            index={index + 1}
          />
        ))}
      </div>
    </motion.div>
  </section>
);

export default Insights;
