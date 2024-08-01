"use client";
import { motion } from "framer-motion";

import styles from "../styles";
import { navVariants } from "../utils/motion";

const Navbar = () => (
  <motion.nav
    variants={navVariants}
    initial="hidden"
    whileInView="show"
    className={`${styles.xPaddings} py-8 relative`}
  >
    <div className="absolute w-[50%] inset-0" />
    <div className={`${styles.innerWidth} mx-auto flex justify-between gap-8`}>
      <img
        src="/Search.png"
        alt="search"
        className="w-[50px] h-[50px] object-contain"
      />
      <img
        src="/Logo.png"
        alt="logo"
        className="w-[250px] h-[24px] object-contain"
      />
      <img
        src="/Menu.png"
        alt="menu"
        className="w-[24px] h-[24px] object-contain"
      />
    </div>
  </motion.nav>
);

export default Navbar;
