import { useSpring } from "@react-spring/web";

/**
 * Hook untuk mengatur animasi spring pada sebuah section berdasarkan apakah section tersebut terlihat atau tidak.
 *
 * @param {boolean} sectionInView Menunjukkan apakah section sedang terlihat di viewport.
 * @returns {Object} Properti animasi untuk section.
 * @returns {import('@react-spring/web').SpringValue<number>} returns.opacity Nilai animasi untuk opacity.
 * @returns {import('@react-spring/web').SpringValue<string>} returns.transform Nilai animasi untuk transform.
 */
const useSectionSpring = (sectionInView) => {
  return useSpring({
    /**
     * Properti animasi untuk opacity.
     * @type {number}
     */
    opacity: sectionInView ? 1 : 0,

    /**
     * Properti animasi untuk transform.
     * @type {string}
     */

    transform: sectionInView ? "translateY(0px)" : "translateY(-100px)",

    /**
     * Konfigurasi animasi spring.
     * @type {Object}
     * @property {number} tension Ketegangan animasi spring.
     * @property {number} friction Gesekan animasi spring.
     */
    config: { tension: 200, friction: 75 },
  });
};

export default useSectionSpring;
