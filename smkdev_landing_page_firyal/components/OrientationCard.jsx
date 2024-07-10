"use client";
import React from "react";
import Image from "next/image";
import { motion } from "framer-motion";

export function OrientationCard() {
  const dataOrientation = [
    {
      desc: "Learning path disusun menyesuaikan dengan perkembangan dan kebutuhan dunia industri.",
      name: "Learning Path Industri",
      image: "/paper.png",
    },
    {
      desc: "Kurikulum komprehensif dan senantiasa di perbarui berdasarkan pengalaman di dunia industri",
      name: "Kurikulum Komprehensif",
      image: "/book.png",
    },
    {
      desc: "Materi pembelajaran disusun dengan pendekatan - pendekatan project based learning",
      name: "Project-Based Learning",
      image: "/laptop.png",
    },
  ];

  return (
    <div className="py-40 px-20 flex flex-col items-center justify-center">
      <h1 className="text-3xl px-20 font-bold">Orientasi Belajar SMKDEV</h1>
      <h1 className="text-xl px-20 flex justify-center text-center pb-20 pt-4">
        Dapatkan pengalaman belajar berorientasi pengalaman kerja yang dapat
        mengantarkan Anda menjadi talenta yang dibutuhkan oleh industri digital
        terkini.
      </h1>
      <div className="flex justify-center items-center">
        <div className="flex flex-wrap">
          {dataOrientation.map((item, idx) => {
            const words = item.name.split(" ");
            const lastWord = words.pop();
            const restOfName = words.join(" ");

            return (
              <motion.div
                key={"item" + idx}
                style={{
                  rotate: 0,
                }}
                whileHover={{
                  scale: 1.1,
                  rotate: Math.random() * 15 - 10,
                  zIndex: 9,
                }}
                whileTap={{
                  scale: 1.1,
                  rotate: 0,
                  zIndex: 9,
                }}
                className="w-60 rounded-xl mr-12 p-4  bg-white border-neutral-200 border  flex-shrink-0 overflow-hidden"
              >
                <div className="w-1/5 flex justify-center rounded-br-lg">
                  <Image
                    src={item.image}
                    alt={item.name}
                    width={30}
                    height={30}
                  />
                </div>
                <h1 className="font-bold text-lg pt-4">
                  <span className="text-black">{restOfName}</span>{" "}
                  <span className="text-[#00A92F]">{lastWord}</span>
                </h1>
                <p className=" line-clamp-4 ">{item.desc}</p>
              </motion.div>
            );
          })}
        </div>
      </div>
    </div>
  );
}
