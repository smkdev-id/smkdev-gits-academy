import gsap from "gsap";
import { ScrollTrigger } from "gsap/ScrollTrigger";

/**
 * Mendaftarkan plugin ScrollTrigger ke GSAP.
 * Ini memungkinkan  untuk menggunakan fitur ScrollTrigger dalam animasi GSAP.
 */
gsap.registerPlugin(ScrollTrigger);

/**
 * Mengatur animasi pada elemen-elemen yang diberikan dan juga pada elemen section yang diberikan.
 *
 * @param {HTMLElement[]} boxes - Array elemen-elemen yang akan dianimasikan.
 * @param {React.RefObject} sectionRef - Referensi ke elemen section yang akan memiliki animasi background.
 */
export const setupAnimations = (boxes, sectionRef) => {
  boxes.forEach((box, index) => {
    if (index === 0) {
      /**
       * Animasi untuk elemen pertama dalam array boxes.
       * Elemen ini akan muncul dengan opacity dari 0 ke 1 dan bergerak dari kiri ke posisi semula.
       */
      gsap.fromTo(
        box,
        { opacity: 0, x: -500 },
        {
          opacity: 1,
          x: 0,
          duration: 0.5,
          scrollTrigger: {
            trigger: box,
            start: "center 90%",
            toggleActions: "play none none reset",
          },
        }
      );
    } else if (index === 2) {
      /**
       * Animasi untuk elemen ketiga dalam array boxes.
       * Elemen ini akan muncul dengan opacity dari 0 ke 1 dan bergerak dari kanan ke posisi semula.
       */
      gsap.fromTo(
        box,
        { opacity: 0, x: 500 },
        {
          opacity: 1,
          x: 0,
          duration: 0.5,
          scrollTrigger: {
            trigger: box,
            start: "center 90%",
            toggleActions: "play none none reset",
          },
        }
      );
    } else {
      /**
       * Animasi untuk elemen-elemen lainnya dalam array boxes.
       * Elemen-elemen ini akan muncul dengan opacity dari 0 ke 1.
       */
      gsap.fromTo(
        box,
        { opacity: 0 },
        {
          opacity: 1,
          duration: 1,
          scrollTrigger: {
            trigger: box,
            start: "center 90%",
            toggleActions: "play none none reset",
          },
        }
      );
    }
  });

  if (sectionRef.current) {
    /**
     * Animasi untuk elemen section yang diberikan.
     * Background position dari elemen ini akan berubah seiring dengan scroll.
     */
    gsap.to(sectionRef.current, {
      backgroundPositionY: "50%",
      ease: "none",
      scrollTrigger: {
        trigger: sectionRef.current,
        scrub: true,
        start: "top bottom",
        end: "bottom top",
      },
    });
  }
};
