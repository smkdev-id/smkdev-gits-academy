import { useState, useEffect } from "react";
import NavbarDropdown from "./NavbarDropdown";
import { dropDownItems } from "./dropDownItemsConstant.js";
import { useSpring, animated } from "react-spring";

/**
 * Komponen Navbar yang menampilkan navigasi dengan animasi muncul dari atas
 * menggunakan react-spring.
 * 
 * @component
 */
const Navbar = () => {
    /**
   * State yang digunakan untuk melacak apakah komponen telah dimuat atau belum.
   * 
   * @type {boolean}
   * @default false
   */
  const [isLoaded, setIsLoaded] = useState(false);

  /**
   * Konfigurasi animasi untuk navbar menggunakan useSpring dari react-spring.
   * 
   * - Animasi akan mengatur posisi top dari "-100px" ke "0px" dan
   *   opacity dari 0 ke 1 ketika komponen telah dimuat (isLoaded = true).
   * - Animasi dimulai dari posisi { top: "-100px", opacity: 5 }.
   * - Durasi animasi adalah 600ms.
   * 
   * @type {Object}
   */
  const navbarAnimation = useSpring({
    top: isLoaded ? "0px" : "-100px",
    opacity: isLoaded ? 1 : 0,
    from: { top: "-100px", opacity: 5 },
    config: { duration: 600 },
  });

   /**
   * Mengatur isLoaded menjadi true ketika komponen telah dimuat.
   * 
   * @function
   * @name useEffect
   */
  useEffect(() => {
    setIsLoaded(true);
  }, []);
  return (
    <animated.div
      className="fixed w-full z-50 bg-white shadow-sm"
      style={navbarAnimation}
    >
      <nav className="flex justify-between items-center px-130  py-4 border-b-2 bg-white ">
        <img
          src="../../public/logo-smk-dev.jpg"
          alt=""
          width={128}
          height={32}
        />
        <div className="flex items-center gap-5 ">
          <NavbarDropdown title="Learn" items={dropDownItems} />
          <span className="font-poppins text-custom-blue font-medium cursor-pointer hover:underline hover:underline-offset-4 hover:decoration-2 hover:decoration-custom-blue">
            Community
          </span>
          <span className="font-poppins text-custom-blue font-medium cursor-pointer hover:underline hover:underline-offset-4 hover:decoration-2 hover:decoration-custom-blue">
            Blog
          </span>
          <button className="font-poppins bg-custom-blue py-2 px-6 rounded-md text-white font-medium">
            Dashboard
          </button>
        </div>
      </nav>
    </animated.div>
  );
};

export default Navbar;
