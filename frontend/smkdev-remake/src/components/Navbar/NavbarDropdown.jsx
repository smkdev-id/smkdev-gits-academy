import { useState } from "react";
import PropTypes from "prop-types";

/**
 * Komponen NavbarDropdown digunakan untuk menampilkan dropdown di navbar.
 * Dropdown akan muncul saat mouse masuk ke area judul (title).
 *
 * @component
 *
 * @param {Object} props Properti yang diterima oleh komponen.
 * @param {string} props.title Judul untuk dropdown.
 * @param {Array<Object>} props.items Daftar item dropdown.
 * @param {string} props.items[].label Label dari setiap item dropdown.
 * @param {string} props.items[].href Link atau href dari setiap item dropdown.
 */
const NavbarDropdown = ({ title, items }) => {
  /**
   * State yang digunakan untuk mengontrol apakah dropdown dalam navbar terbuka atau tidak.
   *
   * @type {boolean}
   * @default false
   */
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);

  /**
   * Menangani peristiwa saat mouse masuk ke area dropdown.
   * Mengatur state isDropdownOpen menjadi true.
   *
   * @function
   * @name handleMouseEnter
   */
  const handleMouseEnter = () => {
    setIsDropdownOpen(true);
  };

  /**
   * Menangani peristiwa saat mouse keluar dari area dropdown.
   * Mengatur state isDropdownOpen menjadi false.
   *
   * @function
   * @name handleMouseLeave
   */
  const handleMouseLeave = () => {
    setIsDropdownOpen(false);
  };
  return (
    <div className="relative group" onMouseEnter={handleMouseEnter}>
      <span
        className={`cursor-pointer text-custom-blue font-medium font-poppins ${
          isDropdownOpen
            ? "underline underline-offset-4 decoration-2 decoration-custom-blue text-custom-blue font-medium"
            : "hover:underline hover:underline-offset-4 hover:decoration-2 hover:decoration-custom-blue-bg"
        }`}
      >
        {title}
      </span>
      {isDropdownOpen && (
        <div
          className="absolute mt-2 w-48 bg-white border border-gray-200 rounded-md shadow-lg"
          onMouseLeave={handleMouseLeave}
        >
          {items.map((item, index) => (
            <a
              key={index}
              href={item.href}
              className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 font-poppins"
            >
              {item.label}
            </a>
          ))}
        </div>
      )}
    </div>
  );
};

/**
 * PropTypes untuk komponen NavbarDropdown, yang mendefinisikan tipe data dari properti yang diterima.
 *
 * @typedef {Object} PropTypes
 * @property {string} title - Judul untuk dropdown. Harus berupa string dan wajib ada.
 * @property {Array<Object>} items - Daftar item dropdown. Setiap item harus memiliki properti label (string) dan href (string).
 */
NavbarDropdown.propTypes = {
  /**
   * Judul untuk dropdown.
   */
  title: PropTypes.string.isRequired,

  /**
   * Daftar item dropdown.
   * Setiap item harus memiliki properti label (string) dan href (string).
   */
  items: PropTypes.arrayOf(
    PropTypes.shape({
      label: PropTypes.string.isRequired,
      href: PropTypes.string.isRequired,
    })
  ).isRequired,
};

export default NavbarDropdown;
