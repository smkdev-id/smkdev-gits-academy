import PropTypes from "prop-types";

/**
 * Komponen Accordion digunakan untuk menampilkan konten dengan opsi untuk dibuka atau ditutup.
 *
 * @component
 *
 * @param {Object} props Properti yang diterima oleh komponen.
 * @param {React.ReactNode} props.children Konten yang akan ditampilkan di dalam Accordion.
 * @param {Function} props.onToggle Fungsi callback untuk mengubah status accordion (terbuka/tutup).
 * @param {boolean} props.isOpen Status saat ini dari accordion (terbuka/tutup).
 * @param {boolean} props.isThreeSlides Menunjukkan apakah komponen ini digunakan dalam slider dengan 3 slides atau tidak.
 */
const Accordion = ({ children, onToggle, isOpen, isThreeSlides }) => {
  /**
   * Panjang minimum konten untuk menampilkan tombol "Baca Selengkapnya".
   *
   * @type {number}
   * @constant
   */
  const minLengthToShowButton = 150;

  /**
   * Callback untuk mengaktifkan fungsi onToggle saat tombol "Baca Selengkapnya" diklik.
   *
   * @function
   * @name toggleAccordion
   */
  const toggleAccordion = () => {
    onToggle();
  };

  return (
    <div className="accordion">
      {children.length > minLengthToShowButton ? (
        <>
          <p
            className={`font-poppins text-gray-500 italic ${
              isOpen ? "" : "line-clamp-3"
            }`}
          >
            {children}
          </p>
          {isThreeSlides && (
            <button
              className="font-poppins text-custom-blue italic hover:text-custom-blue-secondary focus:outline-none transition-all duration-300 mt-2"
              onClick={toggleAccordion}
            >
              {isOpen ? "Sembunyikan" : "Baca Selengkapnya"}
            </button>
          )}
        </>
      ) : (
        <p className=" font-poppins">{children}</p>
      )}
    </div>
  );
};

/**
 * PropTypes untuk komponen Accordion, yang mendefinisikan tipe data dari properti yang diterima.
 *
 * @typedef {Object} PropTypes
 * @property {React.ReactNode} children Konten yang akan ditampilkan di dalam Accordion.
 * @property {Function} onToggle Fungsi callback untuk mengubah status accordion (terbuka/tutup).
 * @property {boolean} isOpen Status saat ini dari accordion (terbuka/tutup).
 * @property {boolean} isThreeSlides Menunjukkan apakah komponen ini digunakan dalam slider dengan 3 slides atau tidak.
 */
Accordion.propTypes = {
  /**
   * Konten yang akan ditampilkan di dalam Accordion.
   */
  children: PropTypes.node.isRequired,

  /**
   * Fungsi callback untuk mengubah status accordion (terbuka/tutup).
   */
  onToggle: PropTypes.func.isRequired,

  /**
   * Status saat ini dari accordion (terbuka/tutup).
   */
  isOpen: PropTypes.bool.isRequired,

  /**
   * Menunjukkan apakah komponen ini digunakan dalam slider dengan 3 slides atau tidak.
   */
  isThreeSlides: PropTypes.bool.isRequired,
};

export default Accordion;
