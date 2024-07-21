import Slider from "react-slick";
import { slidesExpertMentor } from "./SlidesExpertConstant.js";
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";
import { useRef } from "react";

/**
 * Komponen untuk menampilkan slider dengan data mentor ahli.
 * Menggunakan React Slick untuk tampilan slider responsif.
 * @component
 */
const SliderComponent = () => {
  /**
   * Referensi untuk elemen slider menggunakan useRef.
   * Digunakan untuk mengakses instance Slider dari React Slick.
   * @type {React.MutableRefObject<Slider | null>}
   */
  const sliderRef = useRef(null);

  /**
   * Pengaturan konfigurasi slider menggunakan React Slick.
   * @type {Object}
   */
  const settings = {
    dots: false,
    infinite: true,
    speed: 500,
    slidesToShow: 5,
    autoplay: true,
    autoplaySpeed: 5000,
    pauseOnHover: true,
    adaptiveHeight: false,
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
  };

  return (
    <div className="">
      <Slider ref={sliderRef} {...settings}>
        {slidesExpertMentor.map((expertMentor, index) => (
          <div key={index} className="px-4 h-full">
            <div className="  p-6 mb-6 h-full  py-4">
              <div className="grid justify-center">
                <img
                  src={expertMentor.image}
                  alt={expertMentor.name}
                  className="rounded-full"
                  width={200}
                />
                <div className="mt-6 text-center">
                  <p className="text-lg font-semibold font-poppins mb-2 text-custom-color-font">
                    {expertMentor.name}
                  </p>
                  <div className="border-t-2 border-custom-color-font mx-16"></div>
                  <p className="text-cus mb-1 font-poppins pt-3">
                    {expertMentor.job}
                  </p>
                  <p className="text-cus mb-2 font-poppins">
                    {expertMentor.company}
                  </p>
                </div>
              </div>
            </div>
          </div>
        ))}
      </Slider>
    </div>
  );
};

export default SliderComponent;
