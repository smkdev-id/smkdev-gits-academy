import { useEffect, useRef } from "react";
import gsap from "gsap";
import { icons, cards } from "./cardsConstant";
import "tailwindcss/tailwind.css";
import { setupAnimations } from "./lib/gsapUtils.js";
import { addToRefs } from "./lib/refUtils.js";

/**
 * Komponen ScrollAnimation menampilkan card dengan animasi
 * menggunakan GSAP untuk efek scroll.
 *
 * @component
 */
const ScrollAnimation = () => {
  /**
   * Referensi untuk elemen-elemen card yang digunakan dalam animasi GSAP.
   *
   * @type {React.MutableRefObject<Array<Element>>}
   */
  const boxRefs = useRef([]);

  /**
   * Referensi untuk elemen section utama yang digunakan dalam animasi GSAP.
   *
   * @type {React.MutableRefObject<Element | null>}
   */
  const sectionRef = useRef(null);

  /**
   * Efek samping yang dijalankan sekali saat komponen `ScrollAnimation` dimuat.
   * Menggunakan GSAP untuk mengambil elemen dengan kelas 'box' dan
   * menyiapkan animasi dengan setupAnimations pada sectionRef.
   *
   * @effect
   * @name useEffect
   * @function
   */
  useEffect(() => {
    // Mengambil semua elemen dengan kelas 'box' menggunakan GSAP
    const boxes = gsap.utils.toArray(".box");

    // Setup animasi menggunakan utilitas GSAP pada sectionRef
    setupAnimations(boxes, sectionRef);
  }, []);

  return (
    <div
      ref={sectionRef}
      className="flex items-center justify-center font-sans mb-10 bg-fixed bg-cover bg-center overflow-hidden"
    >
      <div className="flex justify-center space-x-8">
        {cards.map((card, index) => {
          const IconComponent = icons[card.icon];

          if (!IconComponent) {
            console.error(`Icon component "${card.icon}" not found.`);
            return null;
          }

          return (
            <div
              key={index}
              ref={(el) => addToRefs(el, boxRefs)}
              className="shadow-md border box rounded-2xl transform hover:shadow-lg transition-transform duration-300 w-96 h-80 my-4 flex-shrink-0 relative px-8 bg-white items-center"
            >
              <div className="bg-[#e5f7ed] w-28 h-28 rounded-tl-2xl rounded-br-lg flex items-center justify-center mb-4 absolute left-0 top-0">
                <IconComponent className="text-green-600 w-11 h-11" />
              </div>
              <div className="mt-32">
                <h2 className="text-xl text-custom-color-font font-medium mb-2 font-poppins">
                  {card.title}
                </h2>
                <p className="font-poppins my-4 text-gray-500 text-md">
                  {card.content}
                </p>
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default ScrollAnimation;
