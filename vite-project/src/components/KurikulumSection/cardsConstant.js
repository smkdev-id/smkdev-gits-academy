
import { FaBook, FaPencilAlt, FaChalkboardTeacher } from "react-icons/fa";


/**
 * Data Card yang digunakan dalam komponen KurikulumSection.
 * Setiap objek merepresentasikan satu Card dengan judul, ikon, deskripsi, warna latar belakang, dan warna teks.
 * 
 * @type {Object[]}
 * @property {string} title Judul dari Card.
 * @property {import('react').ComponentType} icon Ikon yang digunakan untuk Card.
 * @property {string} description Deskripsi singkat tentang topik Card.
 * @property {string} bgColor Warna latar belakang Card dalam format CSS.
 * @property {string} color Warna teks Card dalam format CSS.
 */
export const cards = [
  {
    title: "Kurikulum",
    icon: FaBook,
    description:
      "Kurikulum berstandar industri global, senantiasa diperbaharui",
    bgColor: "bg-[#f3f8fb]",
    color: "text-[#1c3965]",
  },
  {
    title: "Metode",
    icon: FaPencilAlt,
    description:
      "Metode belajar project-based learning, dapat menjadi portofolio siswa",
    bgColor: "bg-[#fef7ef]",
    color: "text-[#ee951a]",
  },
  {
    title: "Mentor",
    icon: FaChalkboardTeacher,
    description:
      "Dibimbing langsung oleh mentor expert dari industri digital",
    bgColor: "bg-[#f0f9f3]",
    color: "text-[#00a92f]",
  },
];
