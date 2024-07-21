import { useSprings } from "@react-spring/web";

/**
 * Hook untuk mengatur animasi spring pada kartu-kartu berdasarkan apakah kartu tersebut terlihat atau tidak.
 *
 * @param {Object[]} cards Array objek yang merepresentasikan setiap kartu.
 * @param {number[]} visibleCards Array indeks dari kartu yang terlihat.
 * @returns {Array} Array properti animasi untuk setiap kartu.
 */
const useCardSprings = (cards, visibleCards) => {
  return useSprings(
    cards.length,
    cards.map((_, index) => ({
      /**
       * Properti animasi untuk opacity.
       * @type {number}
       */
      opacity: visibleCards.includes(index) ? 1 : 0,

      /**
       * Properti animasi untuk transform.
       * @type {string}
       */
      transform: visibleCards.includes(index)
        ? "translateY(0px)"
        : "translateY(100px)",
        
      /**
       * Konfigurasi animasi spring.
       * @type {Object}
       * @property {number} tension Ketegangan animasi spring.
       * @property {number} friction Gesekan animasi spring.
       */
      config: { tension: 200, friction: 75 },
    }))
  );
};

export default useCardSprings;
