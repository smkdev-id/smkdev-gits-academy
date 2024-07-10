"use client";
import { MediaQuery, WhyUsCard } from "@/components";
import { NextPage } from "next";
import { Swiper, SwiperSlide } from "swiper/react";
import { Autoplay, EffectCards } from "swiper/modules";
import "swiper/css/effect-cards";
import "swiper/css";

interface Props {}

const WhyUs: NextPage<Props> = ({}) => {
  const slides = [
    {
      title: "Kurikulum",
      desc: "Kurikulum berstandar industri global, senantiasa diperbaharui",
      color: "#004fc5",
      svg: (
        <svg
          width={35}
          height={35}
          aria-hidden="true"
          viewBox="0 0 512 512"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            className="fill-white"
            d="M352 96c0-53.02-42.98-96-96-96s-96 42.98-96 96 42.98 96 96 96 96-42.98 96-96zM233.59 241.1c-59.33-36.32-155.43-46.3-203.79-49.05C13.55 191.13 0 203.51 0 219.14v222.8c0 14.33 11.59 26.28 26.49 27.05 43.66 2.29 131.99 10.68 193.04 41.43 9.37 4.72 20.48-1.71 20.48-11.87V252.56c-.01-4.67-2.32-8.95-6.42-11.46zm248.61-49.05c-48.35 2.74-144.46 12.73-203.78 49.05-4.1 2.51-6.41 6.96-6.41 11.63v245.79c0 10.19 11.14 16.63 20.54 11.9 61.04-30.72 149.32-39.11 192.97-41.4 14.9-.78 26.49-12.73 26.49-27.06V219.14c-.01-15.63-13.56-28.01-29.81-27.09z"
          />
        </svg>
      ),
    },
    {
      title: "Metode",
      desc: "Metode belajar project-based learning, dapat menjadi portfolio siswa",
      color: "#ee951a",
      svg: (
        <svg
          aria-hidden="true"
          className=""
          viewBox="0 0 512 512"
          width={35}
          height={35}
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            className="fill-white"
            d="M497.9 142.1l-46.1 46.1c-4.7 4.7-12.3 4.7-17 0l-111-111c-4.7-4.7-4.7-12.3 0-17l46.1-46.1c18.7-18.7 49.1-18.7 67.9 0l60.1 60.1c18.8 18.7 18.8 49.1 0 67.9zM284.2 99.8L21.6 362.4.4 483.9c-2.9 16.4 11.4 30.6 27.8 27.8l121.5-21.3 262.6-262.6c4.7-4.7 4.7-12.3 0-17l-111-111c-4.8-4.7-12.4-4.7-17.1 0zM124.1 339.9c-5.5-5.5-5.5-14.3 0-19.8l154-154c5.5-5.5 14.3-5.5 19.8 0s5.5 14.3 0 19.8l-154 154c-5.5 5.5-14.3 5.5-19.8 0zM88 424h48v36.3l-64.5 11.3-31.1-31.1L51.7 376H88v48z"
          />
        </svg>
      ),
    },
    {
      title: "Mentor",
      desc: "Dibimbing langsung oleh mentor expert dari industri digital",
      color: "#00a92f",
      svg: (
        <svg
          width={35}
          height={35}
          aria-hidden="true"
          className=""
          viewBox="0 0 640 512"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            className="fill-white"
            d="M208 352c-2.39 0-4.78.35-7.06 1.09C187.98 357.3 174.35 360 160 360c-14.35 0-27.98-2.7-40.95-6.91-2.28-.74-4.66-1.09-7.05-1.09C49.94 352-.33 402.48 0 464.62.14 490.88 21.73 512 48 512h224c26.27 0 47.86-21.12 48-47.38.33-62.14-49.94-112.62-112-112.62zm-48-32c53.02 0 96-42.98 96-96s-42.98-96-96-96-96 42.98-96 96 42.98 96 96 96zM592 0H208c-26.47 0-48 22.25-48 49.59V96c23.42 0 45.1 6.78 64 17.8V64h352v288h-64v-64H384v64h-76.24c19.1 16.69 33.12 38.73 39.69 64H592c26.47 0 48-22.25 48-49.59V49.59C640 22.25 618.47 0 592 0z"
          />
        </svg>
      ),
    },
  ];
  return (
    <div className="bg-bg2 flex w-full justify-center overflow-visible">
      <MediaQuery className="flex items-center justify-between gap-x-14 overflow-visible py-20">
        <div className="flow-x-visible relative h-[230px] w-full overflow-visible">
          <Swiper
            className="mt-6 h-full w-full"
            spaceBetween={0}
            slidesPerView={1}
            direction="vertical"
            loop
            modules={[Autoplay, EffectCards]}
            autoplay={{
              pauseOnMouseEnter: false,
              delay: 2000,
              disableOnInteraction: false,
            }}
            effect="cards"
            cardsEffect={{
              perSlideOffset: 10,
              slideShadows: false,
              perSlideRotate: 3,
            }}
          >
            {slides.map((i, index) => (
              <SwiperSlide key={index} className="">
                <WhyUsCard
                  color={i.color}
                  desc={i.desc}
                  svg={i.svg}
                  title={i.title}
                />
              </SwiperSlide>
            ))}
          </Swiper>
        </div>

        <div className="flex w-full flex-col gap-y-3">
          <h1 className="text-3xl font-semibold uppercase text-primary">
            MENGAPA HARUS MEMILIH BELAJAR DI{" "}
            <span className="text-[#004fc5]">SMKDEV</span>
          </h1>
          <p className="text-base">
            Pengalaman belajarmu jauh lebih menyenangkan dengan perpaduan
            belajar teori dan capstone project yang dapat menjadi portofolio
            pribadimu. Tidak lupa, Kamu juga dapat belajar langsung dari expert
            mentor dari dunia industri.
          </p>
        </div>
      </MediaQuery>
    </div>
  );
};

export default WhyUs;
