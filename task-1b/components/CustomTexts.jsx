"use client";
import { motion } from "framer-motion";

import styles from "../styles";

import { textContainer, textVariant2 } from "../utils/motion";

export const TypingText = ({ title, textStyles }) => (
  <motion.p
    variants={textContainer}
    className={`font-normal text-[14 px] text-primary-blue ${textStyles} `}
  >
    {Array.from(title).map((letter, index) => (
      <motion.span variants={textVariant2} key={index}>
        {letter === " " ? "\u00A0" : letter}
      </motion.span>
    ))}
  </motion.p>
);

export const TitleText = ({ title, textStyles }) => (
  <motion.h2
    variants={textVariant2}
    initial="hidden"
    whileInView="show"
    className={`mt-[8px] font-bold md:text-[64px] text-[40px] text-primary-blue ${textStyles}`}
  >
    {title}
  </motion.h2>
);
