"use client";
import { Swiper, SwiperSlide } from "swiper/react";

import Image from "next/image";

export default function SwiperCompany() {
  const dataLogo = [
    { image: "/logo_gitsindo.png", name: "GITS Indonesia" },
    { image: "/logo_digits.png", name: "Digits.id" },
    { image: "/logo_arkana.png", name: "Arkana" },
    { image: "/logo_mantabone.png", name: "Mantab One" },
    { image: "/logo-hitopia.png", name: "Hitopia" },
    { image: "/logo-eudeka.png", name: "Eudeka" },
    { image: "/logo_gitsindo.png", name: "GITS Indonesia" },
    { image: "/logo_digits.png", name: "Digits.id" },
    { image: "/logo_arkana.png", name: "Arkana" },
    { image: "/logo_mantabone.png", name: "Mantab One" },
    { image: "/logo-hitopia.png", name: "Hitopia" },
    { image: "/logo-eudeka.png", name: "Eudeka" },
  ];
  return (
    <Swiper
      loop={true}
      spaceBetween={50}
      slidesPerView={7}
      onSlideChange={() => console.log("slide change")}
      onSwiper={(swiper) => console.log(swiper)}
    >
      {dataLogo.map((item) => (
        <SwiperSlide key={item.name}>
          <Image alt={item.name} src={item.image} width={150} height={150} />
        </SwiperSlide>
      ))}
    </Swiper>
  );
}
