import { NextPage } from "next";
import Image from "next/image";
import Marquee from "react-fast-marquee";

const Mitra: NextPage = () => {
  const slides = [
    { src: "/assets/Logo-Digits.png", alt: "Logo Digits" },
    { src: "/assets/Logo-Gitsid.png", alt: "Logo Gitsid" },
    { src: "/assets/Logo-Eudeka.png", alt: "Logo Eudeka" },
    { src: "/assets/Logo-Arkana.png", alt: "Logo Arkana" },
    { src: "/assets/Logo-Hitopia.png", alt: "Logo Hitopia" },
    { src: "/assets/Logo-Mantabone.png", alt: "Logo Mantabone" },
  ];

  return (
    <div className="bg-bg2 flex w-full justify-center py-16">
      <div className="flex w-[60%] flex-col items-center justify-center gap-y-3">
        <h1 className="text-center text-xl font-semibold capitalize text-primary">
          Dipercaya Oleh Mitra Industri
        </h1>
        <Marquee className="w-full">
          {slides.concat(slides).map((slide, index) => (
            <div key={index} className="flex w-[200px] justify-center">
              <Image
                src={slide.src}
                alt={slide.alt}
                width={130}
                height={130}
                sizes="100vw"
                className=""
              />
            </div>
          ))}
        </Marquee>
      </div>
    </div>
  );
};

export default Mitra;
