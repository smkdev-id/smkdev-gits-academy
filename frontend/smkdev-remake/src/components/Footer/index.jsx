import ContactForm from "./ContactForm";
import {
  FaYoutube,
  FaLinkedinIn,
  FaInstagram,
  FaTiktok,
  FaTelegram,
  FaWhatsapp,
} from "react-icons/fa";

const Footer = () => {
  return (
    <>
      <footer className="bg-custom-blue-bg mt-24 py-24 px-130">
        <div className="flex  gap-3">
          <div className="w-full">
            <h1 className="font-poppins text-custom-blue font-medium text-3xl">
              Kontak SMKDEV
            </h1>
            <p className="font-poppins text-base text-custom-color-font py-4">
              Berapa investasi untuk menggunaan SMKDEV? Anda bisa mulai dengan{" "}
              <span className="font-poppins font-bold">GRATIS</span>
            </p>
            <p className="font-poppins text-custom-color-font">
              Silahkan isi form disebelah untuk informasi lebih lanjut.{" "}
            </p>
          </div>
          <ContactForm />
        </div>

        <div className="mt-20">
          <img
            src="../../../public/logo-smk-dev.jpg"
            alt=""
            width={128}
            height={32}
          />
          <p className="font-poppins text-gray-500 font-medium pt-3">
            Creating High-Caliber Digital Talent
          </p>
          <p className="font-poppins text-gray-500 font-medium pt-5">
            Summarecon Bandung, Jl. Magna Barat Blok MD No.02, Rancabolang, Kec.
            Gedebage, Kota Bandung, Jawa Barat 40295
          </p>
          <p className="font-poppins text-custom-color-font-blue font-semibold pt-5">
            Follow Us
          </p>
          <div className="flex gap-1 pt-2">
            <div className=" rounded p-2 bg-white">
              <FaYoutube className="text-custom-color-font-blue" />
            </div>
            <div className=" rounded p-2 bg-white">
              <FaLinkedinIn className="text-custom-color-font-blue" />
            </div>
            <div className=" rounded p-2 bg-white">
              <FaInstagram className="text-custom-color-font-blue" />
            </div>
            <div className=" rounded p-2 bg-white">
              <FaTiktok className="text-custom-color-font-blue" />
            </div>
            <div className=" rounded p-2 bg-white">
              <FaTelegram className="text-custom-color-font-blue" />
            </div>
            <div className=" rounded p-2 bg-white">
              <FaWhatsapp className="text-custom-color-font-blue" />
            </div>
          </div>
        </div>
      </footer>
      <p className="text-center font-poppins text-custom-blue py-3">
        © 2024 SMKDEV – PT Eureka Merdeka Indonesia. All Rights Reserved.
      </p>
    </>
  );
};

export default Footer;
