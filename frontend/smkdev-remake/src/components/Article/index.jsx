import Slider from "react-slick";
import ArticleCard from "./ArticleCard";
import { useRef } from "react";
import { articles } from "./articlesConstant";

/**
 * Komponen untuk menampilkan daftar artikel menggunakan slider.
 * @component
 */
const Article = () => {
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
    dots: true,
    infinite: true,
    speed: 500,
    slidesToShow: 4,
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
  };

  return (
    <div>
      <h1 className="text-3xl font-medium font-poppins mb-5 text-custom-color-font text-center">
        Artikel Pilihan SMKDEV
      </h1>
      <div className="px-130">
        <Slider ref={sliderRef} {...settings}>
          {articles.map((article, index) => (
            <ArticleCard
              key={index}
              image={article.image}
              title={article.title}
              description={article.description}
              link={article.link}
            />
          ))}
        </Slider>
        <div className="flex justify-center mt-10 ">
          <button className="font-poppins bg-custom-blue py-2 px-6 rounded-md text-white font-medium transform hover:shadow-lg transition-transform duration-300">
            Lihat Semua Artikel
          </button>
        </div>
      </div>
    </div>
  );
};

export default Article;
