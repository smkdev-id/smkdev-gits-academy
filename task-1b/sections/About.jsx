"use client";
import { motion } from "framer-motion";

import styles from "../styles";
import { TypingText } from "../components";
import { fadeIn, staggerContainer, textVariant } from "../utils/motion";

const About = () => (
  <section className={`${styles.paddings} relative z-10`}>
    <div className="gradient-02 z-0" />
    <motion.div
      variants={staggerContainer}
      initial="hidden"
      whileInView="show"
      viewport={{ once: false, amount: 0.25 }}
      className={`${styles.innerWidth} mx-auto ${styles.flexCenter} flex-col`}
    >
      <TypingText title="| About SMKDEV" textStyles="text-center" />
      <motion.p
        variants={fadeIn("up", "tween", 0.2, 1)}
        className="mt-[8px] font-normal sm:text-[32px] text-[20px] text-center text-secondary-white"
      >
        <span className="font-extrabold">SMKDEV</span>
        adalah komunitas yang berkomitmen untuk membina talenta SMK melalui{" "}
        <span className="font-extrabold">
          pendekatan project-based learning
        </span>
        , di mana para siswa diberi kesempatan untuk terlibat langsung dalam
        berbagai studi kasus project IT yang berasal dari dunia industri. Dengan
        metode ini, siswa tidak hanya mendapatkan pengetahuan teoritis, tetapi
        juga pengalaman praktis yang
        <span className="font-extrabold">
          relevan dengan kebutuhan pasar kerja saat ini
        </span>
        . Melalui bimbingan dan dukungan dari para profesional, SMKDEV bertujuan
        untuk{" "}
        <span className="font-extrabold">
          mempersiapkan generasi muda agar siap menghadapi tantangan di industri
          teknologi informasi, sekaligus mendorong pengembangan keterampilan
          teknis dan soft skills yang dibutuhkan dalam dunia kerja
        </span>
        . Komunitas ini menjadi jembatan yang menghubungkan pendidikan dengan
        industri, memastikan bahwa talenta muda memiliki kompetensi yang tepat
        untuk sukses di masa depan.
      </motion.p>
      <motion.img
        variants={fadeIn("up", "tween", 0.3, 1)}
        src="/arrow-down.svg"
        alt="arrow down"
        className="w-[18px] h-[28px] object-contain mt-[28px]"
      />
    </motion.div>
  </section>
);

export default About;
