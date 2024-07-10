"use client";
import React from "react";
import Image from "next/image";
import { motion } from "framer-motion";

export function Why() {
  const dataOrientation = [
    {
      desc: "Kurikulum berstandar industri global, senantiasa diperbaharui",
      name: "Kurikulum",
      image: "/book.png",
    },
    {
      desc: "Metode belajar project-based learning, dapat menjadi portfolio siswa",
      name: "Metode",
      image: "/paper.png",
    },
    {
      desc: "Dibimbing langsung oleh mentor expert dari industri digital",
      name: "Mentor",
      image: "/teachings.png",
    },
  ];

  return (
    <div className="relative py-20 px-20 flex flex-row justify-center">
      <div className="absolute inset-0 z-0 flex justify-center items-center w-full h-auto">
        <Image
          src="/decoration.png"
          width="1500"
          height="100"
          alt="decoration"
          className="opacity-30"
        />
      </div>
      <div className="flex flex-col w-1/2 justify-center px-20 z-10">
        <h1 className="text-3xl font-bold">
          Mengapa Harus Memilih Belajar di{" "}
          <span className="text-[#004FC5]">SMKDEV?</span>
        </h1>
        <h1 className="text-xl flex pt-4">
          Pengalaman belajarmu jauh lebih menyenangkan dengan perpaduan belajar
          teori dan capstone project yang dapat menjadi portofolio pribadimu.
          Tidak lupa, Kamu juga dapat belajar langsung dari expert mentor dari
          dunia industri.
        </h1>
      </div>

      <div className="flex flex-col justify-center items-center w-1/2 z-10">
        <div className="flex flex-col">
          {dataOrientation.map((item, idx) => {
            return (
              <motion.div
                key={"item" + idx}
                style={{
                  rotate: 0,
                }}
                whileHover={{
                  scale: 1.1,
                  rotate: Math.random() * 10 - 5,
                  zIndex: 100,
                }}
                whileTap={{
                  scale: 1.1,
                  rotate: 0,
                  zIndex: 100,
                }}
                className="w-210 mt-4 rounded-xl mr-12 p-4 bg-white border-neutral-200 border flex-shrink-0 overflow-hidden"
              >
                <div className="flex justify-start rounded-br-lg">
                  <Image
                    src={item.image}
                    alt={item.name}
                    width={30}
                    height={30}
                  />
                </div>
                <h1 className="font-bold text-lg pt-4 text-[#EEA642]">
                  {item.name}
                </h1>
                <p className="line-clamp-2">{item.desc}</p>
              </motion.div>
            );
          })}
        </div>
      </div>
    </div>
  );
}
