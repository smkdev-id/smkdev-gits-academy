import { useEffect, useState, createRef } from "react";
import { animated } from "@react-spring/web";
import CardKurikulumSection from "./CardKurikulumSection";
import { useInView } from "react-intersection-observer";
import useParallax from "./useParallax";
import useCardSprings from "./lib/cardSpringsUtils.js";
import getSectionSpring from "./lib/sectionSpringUtils.js";
import { cards } from "./cardsConstant.js";

/**
 * Komponen untuk bagian kurikulum dengan efek paralaks dan animasi kartu.
 *
 * Komponen ini menggunakan React hooks dan animasi dari `@react-spring/web` untuk efek paralaks dan animasi kartu.
 *
 * @component
 * @returns {JSX.Element} Komponen React untuk bagian kurikulum
 */
const KurikulumSection = () => {
  /**
   * Hook untuk menghitung offset parallax.
   * @type {import('@react-spring/web').SpringValue<number>}
   */
  const parallaxOffset = useParallax();

  /**
   * Hook untuk mendeteksi apakah bagian dalam viewport.
   * @type {Object}
   * @property {React.MutableRefObject<Element>} ref - Ref untuk bagian yang diamati.
   * @property {boolean} inView - Status apakah bagian dalam viewport.
   */
  const { ref: sectionRef, inView: sectionInView } = useInView({
    triggerOnce: true,
    threshold: 0.5,
  });

  /**
   * State untuk menyimpan referensi terhadap setiap kartu yang dirender.
   * @type {React.MutableRefObject<HTMLDivElement>[]}
   */
  const [cardRefs, setCardRefs] = useState([]);

  /**
   * State untuk melacak kartu mana yang terlihat dalam viewport.
   * @type {number[]}
   */
  const [visibleCards, setVisibleCards] = useState([]);

  /**
   * Efek untuk membuat referensi kartu menggunakan createRef() saat komponen pertama kali dirender.
   */
  useEffect(() => {
    setCardRefs(cards.map(() => createRef()));
  }, []);

  /**
   * Efek untuk mengamati setiap kartu yang terlihat menggunakan IntersectionObserver.
   */
  useEffect(() => {
    if (cardRefs.length > 0) {
      cardRefs.forEach((ref, index) => {
        const observer = new IntersectionObserver(
          ([entry]) => {
            if (entry.isIntersecting) {
              setVisibleCards((prev) => [...new Set([...prev, index])]);
            } else {
              setVisibleCards((prev) => prev.filter((i) => i !== index));
            }
          },
          { threshold: 0.5 }
        );
        if (ref.current) {
          observer.observe(ref.current);
        }
      });
    }
  }, [cardRefs]);

  /**
   * Mendapatkan properti animasi untuk setiap kartu yang terlihat.
   * @type {import('@react-spring/web').SpringValue<import('react').CSSProperties>[]}
   */
  const cardSprings = useCardSprings(cards, visibleCards);

  /**
   * Mendapatkan properti animasi untuk bagian (section) berdasarkan apakah dalam viewport.
   * @type {import('@react-spring/web').SpringValue<import('react').CSSProperties>}
   */
  const sectionSpring = getSectionSpring(sectionInView);

  return (
    <section className="relative">
      <animated.div
        style={{
          transform: parallaxOffset.to(
            (offset) => `translateY(${offset - 240}px)`
          ),
        }}
        className="absolute inset-0 bg-custom-blue-bg bg-opacity-80 bg-smkdev-dashed"
      />
      <div className="relative grid grid-cols-2 px-130 py-16 gap-x-20 items-center z-10">
        <div className="grid gap-5">
          {cardSprings.map((props, index) => (
            <animated.div key={index} style={props} ref={cardRefs[index]}>
              <CardKurikulumSection
                title={cards[index].title}
                icon={cards[index].icon}
                description={cards[index].description}
                bgColor={cards[index].bgColor}
                color={cards[index].color}
                isFirst={index === 0}
              />
            </animated.div>
          ))}
        </div>
        <animated.div
          ref={sectionRef}
          style={sectionSpring}
          className="mb-6 z-10"
        >
          <h2 className="text-3xl font-poppins font-medium text-custom-blue">
            MENGAPA HARUS MEMILIH BELAJAR DI{" "}
            <span className="text-custom-blue-secondary">SMKDEV</span>
          </h2>
          <p className="text-gray-600 mt-4 font-poppins">
            Pengalaman belajarmu jauh lebih menyenangkan dengan perpaduan
            belajar teori dan capstone project yang dapat menjadi portofolio
            pribadimu. Tidak lupa, Kamu juga dapat belajar langsung dari expert
            mentor dari dunia industri.
          </p>
        </animated.div>
      </div>
    </section>
  );
};

export default KurikulumSection;
