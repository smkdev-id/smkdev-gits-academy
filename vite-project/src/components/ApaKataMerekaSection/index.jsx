// ApaKataMerekaSection.js
import { useState, useRef, useEffect } from "react";
import { testimonials } from "./testimonialConstant";
import Slider from "react-slick";
import Accordion from "./Accordion"; // Impor Accordion dari file terpisah
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";

/**
 * Komponen untuk bagian "Apa Kata Mereka tentang SMKDEV".
 * Menampilkan testimonial dari pengguna SMKDEV menggunakan slider.
 * @component
 */
const ApaKataMerekaSection = () => {
  const sliderRef = useRef(null);
  const [openAccordion, setOpenAccordion] = useState(null);
  const [, setSliderHeight] = useState("auto");

  /**
   * Konfigurasi pengaturan slider menggunakan react-slick.
   * @type {Object}
   */
  const settings = {
    dots: true,
    infinite: true,
    speed: 500,
    slidesToShow: 3,
    slidesToScroll: 1,
    autoplay: true,
    autoplaySpeed: 5000,
    pauseOnHover: true,
    adaptiveHeight: true,
    responsive: [
      {
        breakpoint: 1024,
        settings: {
          slidesToShow: 2,
        },
      },
      {
        breakpoint: 600,
        settings: {
          slidesToShow: 1,
        },
      },
    ],
    /**
     * Callback yang dipanggil sebelum perubahan slide.
     * Mengatur ketinggian slider kembali ke nilai default sebelum perubahan.
     */
    beforeChange: () => {
      setSliderHeight("auto");
      setOpenAccordion(null);
    },
    /**
     * Callback yang dipanggil setelah perubahan slide.
     * Memperbarui ketinggian slider setelah perubahan selesai.
     */
    afterChange: () => {
      setTimeout(() => {
        updateSliderHeight();
      }, 300);
    },
  };

  /**
   * Efek samping untuk memperbarui ketinggian slider saat komponen dimuat dan saat jendela diubah ukurannya.
   */
  useEffect(() => {
    updateSliderHeight();
    window.addEventListener("resize", updateSliderHeight);

    return () => {
      window.removeEventListener("resize", updateSliderHeight);
    };
  }, []);

  /**
   * Fungsi untuk menghitung dan mengatur ketinggian slider berdasarkan slide yang aktif.
   */
  const updateSliderHeight = () => {
    if (sliderRef.current && sliderRef.current.innerSlider) {
      const slides =
        sliderRef.current.innerSlider.list.querySelectorAll(".slick-slide");
      const activeSlides = Array.from(slides).filter((slide) =>
        slide.classList.contains("slick-active")
      );
      const maxHeight = Math.max(
        ...activeSlides.map((slide) => slide.offsetHeight)
      );
      setSliderHeight(`${maxHeight}px`);
    }
  };

  /**
   * Mengatur status accordion terbuka atau tertutup.
   * @param {number} id - ID testimonial yang akan diubah status accordion-nya.
   */
  const handleAccordionToggle = (id) => {
    setOpenAccordion(openAccordion === id ? null : id);
    setTimeout(() => {
      updateSliderHeight();
    }, 0);
  };

  return (
    <section className="kata-mereka-section my-24 ">
      <div className="container mx-auto ">
        <h2 className="text-3xl font-medium font-poppins text-center mb-5 text-custom-color-font">
          Kata Mereka tentang SMKDEV
        </h2>
        <p className="text-center text-xl text-custom-color-font font-normal font-poppins">
          Mereka telah merasakan serunya belajar skill digital di SMKDEV. Apakah
          kamu ingin seperti mereka?
        </p>
        <div className="slider-container transition-all duration-300 pt-10">
          <Slider ref={sliderRef} {...settings}>
            {testimonials.map((testimonial) => (
              <div key={testimonial.id} className="px-4 h-full">
                <div className="bg-white rounded-lg shadow-md p-6 mb-6 h-full  py-4">
                  <div>
                    <p className="text-lg font-semibold font-poppins mb-2 text-custom-color-font italic">
                      {testimonial.name}
                    </p>
                    <p className="text-cus mb-2 font-poppins italic">
                      {testimonial.job}
                    </p>
                    <Accordion
                      onToggle={() => handleAccordionToggle(testimonial.id)}
                      isOpen={openAccordion === testimonial.id}
                      isThreeSlides={settings.slidesToShow === 3}
                    >
                      {testimonial.quote}
                    </Accordion>
                  </div>
                </div>
              </div>
            ))}
          </Slider>
        </div>
      </div>
    </section>
  );
};

export default ApaKataMerekaSection;
