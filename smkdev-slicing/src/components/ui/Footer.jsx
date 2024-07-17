import img from "../../assets/image";

import {
  FaYoutube,
  FaLinkedin,
  FaInstagram,
  FaTiktok,
  FaTelegram,
  FaWhatsapp,
} from "react-icons/fa";

const Footer = () => {
  const media = [
    {
      icon: FaYoutube,
      link: "https://www.youtube.com/@smkdev",
    },
    {
      icon: FaLinkedin,
      link: "https://www.linkedin.com/company/smkdev",
    },
    {
      icon: FaInstagram,
      link: "https://www.instagram.com/smkdev.official/",
    },
    {
      icon: FaTiktok,
      link: "https://www.tiktok.com/@smkdev.official",
    },
    {
      icon: FaTelegram,
      link: "https://t.me/+wPXVCS29UKgwNmJl",
    },
    {
      icon: FaWhatsapp,
      link: "https://chat.whatsapp.com/D5gVS4FbxSc2ij1zbZVxAV",
    },
  ];
  return (
    <footer className="mt-10 bg-gray-900 lg:px-10 md:px-10">
      <div className="w-full max-w-screen-xl p-4 py-6 mx-auto lg:py-8">
        <div className="md:flex md:justify-between">
          <div className="mb-6 md:mb-0">
            <a href="https://www.smk.dev/" className="flex items-center">
              <img
                src={img.logo}
                className="h-8 pb-2 me-3"
                alt="smkdev"
                width={120}
                height={100}
              />
            </a>
            <span className="self-center text-base font-semibold text-slate-200 whitespace-nowrap">
              Creating High-Caliber Digital Talent
            </span>
          </div>
          <div className="grid grid-cols-1 gap-8 sm:gap-6 sm:grid-cols-2">
            <div>
              <h2 className="mb-6 text-sm font-semibold text-white uppercase">
                Resources
              </h2>
              <ul className="font-medium text-gray-400">
                <li className="mb-4">
                  <a href="https://www.smk.dev/" className="hover:underline">
                    SMKDEV
                  </a>
                </li>
                <li>
                  <a
                    href="https://tailwindcss.com/"
                    className="hover:underline"
                  >
                    Tailwind CSS
                  </a>
                </li>
              </ul>
            </div>
            <div>
              <h2 className="mb-6 text-sm font-semibold text-white uppercase">
                Legal
              </h2>
              <ul className="font-medium text-gray-400">
                <li className="mb-4">
                  <a href="#" className="hover:underline">
                    Privacy Policy
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:underline">
                    Terms &amp; Conditions
                  </a>
                </li>
              </ul>
            </div>
          </div>
        </div>
        <hr className="my-6 border-gray-700 sm:mx-auto lg:my-8" />
        <div className="sm:flex sm:items-center sm:justify-between">
          <span className="text-sm text-gray-400 sm:text-center">
            © 2004{" "}
            <a href="https://www.smk.dev/" className="hover:underline">
              SMKDEV™
            </a>
            . All Rights Reserved.
          </span>
          <div className="flex mt-4 sm:justify-center sm:mt-0">
            {media.map((provider, i) => {
              return (
                <a
                  href={provider.link}
                  target="blank"
                  className="text-gray-500 hover:text-white ms-5"
                  key={i}
                >
                  <provider.icon className="w-4 h-4" />
                </a>
              );
            })}
          </div>
        </div>
      </div>
    </footer>
  );
};

export default Footer;