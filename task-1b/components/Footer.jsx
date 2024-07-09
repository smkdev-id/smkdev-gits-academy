"use client";
import { motion } from "framer-motion";
import { socials } from "../constants";

import styles from "../styles";
import { footerVariants } from "../utils/motion";

const Footer = () => (
  <motion.footer
    variants={footerVariants}
    initial="hidden"
    whileInView="show"
    className={`${styles.paddings} py-8 relative`}
  >
    <div className={`${styles.innerWidth} mx-auto flex flex-col gap-8`}>
      <div className="flex items-center justify-between flex-warp gap-5 "></div>
      <div className="flex flex-col ">
        <div className="mb-[50px] h-[20px] bg-white-opacity-10">
          <div className="flex items-center justify-between flex-warp gap-4">
            <img
              src="/Logo.png"
              alt="logo"
              className="w-[250px] h-[24px] object-contain"
            />{" "}
            Creating High-Caliber Digital Talent
            <p className="text-primary-blue text-12 font-light mt-2">
              Summarecon Bandung, Jl. Magna Barat Blok MD No.02, Rancabolang,
              Kec. Gedebage, Kota Bandung, Jawa Barat 40295
              <br />
              <br /> © 2024 SMKDEV – PT Eureka Merdeka Indonesia. All Rights
              Reserved.
            </p>
            <div className="flex gap-4">
              {socials.map((social) => (
                <img
                  key={social.name}
                  src={social.url}
                  className="w-[50px] h-[50px] object-contain cursor-pointer"
                />
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  </motion.footer>
);

export default Footer;
