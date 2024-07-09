import { FaBookOpen, FaBook } from "react-icons/fa";
import { FaLaptopFile } from "react-icons/fa6";

/**
 * Objek yang berisi referensi ke komponen ikon dari react-icons/fa dan react-icons/fa6.
 * Digunakan untuk memetakan nama ikon ke komponen ikon yang sesuai.
 *
 * @type {Object}
 * @property {React.ComponentType} bookOpen - Komponen ikon FaBookOpen dari react-icons/fa.
 * @property {React.ComponentType} book - Komponen ikon FaBook dari react-icons/fa.
 * @property {React.ComponentType} laptopFile - Komponen ikon FaLaptopFile dari react-icons/fa6.
 */
export const icons = {
  bookOpen: FaBookOpen,
  book: FaBook,
  laptopFile: FaLaptopFile,
};

/**
 * Array yang berisi data kartu-kartu dengan ikon, judul, dan konten.
 *
 * @type {Object[]}
 * @property {string} icon - Nama ikon yang sesuai dengan properti pada objek icons.
 * @property {string} title - Judul kartu.
 * @property {string} content - Konten atau deskripsi kartu.
 */
export const cards = [
  {
    icon: "bookOpen",
    title: "Learning Path Industri",
    content: "Learning path disusun berdasarkan kebutuhan dunia industri.",
  },
  {
    icon: "book",
    title: "Kurikulum Komprehensif",
    content:
      "Kurikulum komprehensif dan senantiasa di perbarui berdasarkan pengalaman di dunia industri",
  },
  {
    icon: "laptopFile",
    title: "Project-Based Learning",
    content:
      "Materi pembelajaran disusun dengan pendekatan - pendekatan project based learning",
  },
];
