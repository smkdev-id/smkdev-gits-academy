import Slider from "react-slick";
import { slides } from "./slidesConstant.js";
import { settingsTop } from "./lib/sliderSettingsTop.js";
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";

/**
 * Komponen untuk menampilkan slider menggunakan react-slick.
 *
 * @component
 */
const SliderComponent = () => {
  // Referensi untuk slider react-slick
  let sliderTop;

  /**
   * Menghentikan otomatisasi slider ketika mouse masuk.
   *
   * @param {Slider} slider - Instance slider react-slick.
   */
  const handleMouseEnter = (slider) => {
    slider.slickPause();
  };

  /**
   * Memulai otomatisasi slider ketika mouse keluar.
   *
   * @param {Slider} slider - Instance slider react-slick.
   */
  const handleMouseLeave = (slider) => {
    slider.slickPlay();
  };

  return (
    <div className="px-130">
      <Slider ref={(slider) => (sliderTop = slider)} {...settingsTop}>
        {slides.map((slide, index) => (
          <div
            key={index}
            className="slide p-4"
            onMouseEnter={() => handleMouseEnter(sliderTop)}
            onMouseLeave={() => handleMouseLeave(sliderTop)}
          >
            <img
              src={slide.logo}
              alt={`${slide.name} logo`}
              width={150}
              className="object-contain"
            />
          </div>
        ))}
      </Slider>
    </div>
  );
};

export default SliderComponent;
