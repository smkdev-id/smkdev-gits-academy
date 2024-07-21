import PropTypes from "prop-types";

/**
 * Komponen untuk menampilkan Card kurikulum dengan judul, ikon, deskripsi, warna latar belakang, dan warna teks.
 *
 * @component
 * @param {Object} props Properti yang diterima oleh komponen.
 * @param {string} props.title Judul dari Card.
 * @param {React.ElementType} props.icon Ikon yang digunakan untuk Card.
 * @param {string} props.description Deskripsi singkat tentang topik Card.
 * @param {string} props.bgColor Warna latar belakang Card dalam format CSS.
 * @param {string} props.color Warna teks Card dalam format CSS.
 * @param {boolean} props.isFirst Menunjukkan apakah Card ini adalah Card pertama.
 * @returns {JSX.Element} Komponen CardKurikulumSection
 */
const CardKurikulumSection = ({
  title,
  icon: Icon,
  description,
  bgColor,
  color,
  isFirst,
}) => {
  return (
    <div
      className={`flex items-center w-full p-4 ${
        isFirst ? "mt-20" : ""
      } bg-white rounded-2xl shadow-md`}
    >
      <div className={`text-4xl mr-4 ${color} p-5 rounded-2xl ${bgColor}`}>
        <Icon className="icon" />
      </div>
      <div className="h-full">
        <h3 className="text-2xl font-poppins font-medium text-custom-blue pb-2">
          {title}
        </h3>
        <p className="text-gray-600 font-poppins">{description}</p>
      </div>
    </div>
  );
};

CardKurikulumSection.propTypes = {
  /**
   * Judul dari kartu.
   */
  title: PropTypes.string.isRequired,

  /**
   * Ikon yang digunakan untuk kartu.
   */
  icon: PropTypes.elementType.isRequired,

  /**
   * Deskripsi singkat tentang topik kartu.
   */
  description: PropTypes.string.isRequired,

  /**
   * Warna latar belakang kartu dalam format CSS.
   */
  bgColor: PropTypes.string.isRequired,

  /**
   * Warna teks kartu dalam format CSS.
   */
  color: PropTypes.string.isRequired,

  /**
   * Menunjukkan apakah kartu ini adalah kartu pertama.
   */
  isFirst: PropTypes.bool.isRequired,
};

export default CardKurikulumSection;
