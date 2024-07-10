"use client";
import { FC } from "react";
import { Swiper, SwiperSlide } from "swiper/react";
import { Autoplay } from "swiper/modules";
import "swiper/css";
import Image from "next/image";

interface Props {
  list: {
    img: string;
    name: string;
    role: string;
    at: string;
  }[];
}

const MentorList: FC<Props> = ({ list }) => {
  return (
    <>
      <Swiper
        className="w-full cursor-pointer"
        spaceBetween={10}
        slidesPerView={5}
        loop
        modules={[Autoplay]}
        autoplay={{
          pauseOnMouseEnter: false,
          delay: 1500,
          disableOnInteraction: false,
        }}
      >
        {list.map((i, index) => (
          <SwiperSlide key={index} className="w-full">
            <div className="flex w-full flex-col items-center gap-y-4">
              <div className="w-[170px]">
                <Image
                  src={i.img}
                  alt={i.name}
                  width="0"
                  height="0"
                  sizes="100vw"
                  className="h-auto w-full rounded-xl"
                />
              </div>

              <div className="flex flex-col items-center gap-y-3">
                <h1 className="text-center text-xl font-semibold capitalize text-primary">
                  {i.name}
                </h1>
                <div className="h-[2px] w-24 bg-primary" />
                <div className="flex flex-col items-center">
                  <p className="text-center text-base">{i.role}</p>
                  <p className="text-center text-base">At</p>
                  <p className="text-center text-base">{i.at}</p>
                </div>
              </div>
            </div>
          </SwiperSlide>
        ))}
      </Swiper>
    </>
  );
};

export default MentorList;
