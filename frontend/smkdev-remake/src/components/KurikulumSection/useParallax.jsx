import { useEffect } from "react";
import { useSpring } from "react-spring";

/**
 * Hook untuk efek paralaks berdasarkan pergerakan scroll.
 * Menggunakan `react-spring` untuk mengatur animasi efek paralaks.
 *
 * @returns {number} Nilai animasi untuk efek paralaks
 */
const useParallax = () => {
  /**
   * State untuk mengontrol animasi efek paralaks.
   * @type {import('react-spring').SpringValue<{ y: number }>}
   */
  const [{ y }, set] = useSpring(() => ({ y: 0 }));

  useEffect(() => {
    /**
     * Event listener untuk mengatur efek paralaks berdasarkan pergerakan scroll.
     */
    const handleScroll = () => {
      const scrollTop = window.scrollY;
      set({ y: scrollTop / 5 });
    };

    // Melekatkan event listener pada event scroll
    window.addEventListener("scroll", handleScroll);

    // Membersihkan event listener saat komponen dilepas
    return () => window.removeEventListener("scroll", handleScroll);
  }, [set]);

  // Mengembalikan nilai animasi untuk efek paralaks
  return y;
};

export default useParallax;
