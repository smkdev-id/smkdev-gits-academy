"use client";
import React from "react";
import Image from "next/image";
import { motion } from "framer-motion";

export function MentorCard() {
  const dataMentor = [
    {
      image: "/mentor1.png",
      name: "Helmi Adi Prasetyo",
      role: "Backend Developer ",
      company: "SMKDEV",
    },
    {
      image: "/mentor2.jpg",
      name: "Samuel Pandoham ",
      role: "Curriculum Developer",
      company: "SMKDEV",
    },
    {
      image: "/mentor3.jpg",
      name: "Rachmawati Ari Taurisia",
      role: "Curriculum Developer",
      company: "SMKDEV",
    },
    {
      image: "/mentor4.jpeg",
      name: "Sudaryarno",
      role: "CTO",
      company: "GITS Indonesia",
    },
    {
      image: "/mentor5.jpg",
      name: "Ibnu Sina Wardy",
      role: "CEO ",
      company: "GITS Indonesia",
    },
  ];

  return (
    <div className="relative mt-20 py-20 max-w-6xl flex flex-col justify-center">
      <div className="flex flex-col justify-center text-center">
        <h1 className="text-3xl font-bold">
          Expert <span className="text-[#004FC5]">Mentor</span>
        </h1>
        <h1 className="text-xl flex pt-4 text-center justify-center">
          SMKDEV menghadirkan expert dari industri untuk mendampingi proses
          belajarmu.
        </h1>
      </div>

      <div className="flex flex-row justify-center items-center pt-6">
        {dataMentor.map((item, idx) => {
          const words = item.name.split(" ");
          const firstWord = words.shift();
          const restOfName = words.join(" ");

          return (
            <motion.div
              key={"item" + idx}
              style={{
                rotate: 0,
              }}
              whileHover={{
                scale: 1.1,
                rotate: Math.random() * 10 - 5,
                zIndex: 9,
              }}
              whileTap={{
                scale: 1.1,
                rotate: 0,
                zIndex: 9,
              }}
              className="w-1/5 h-[260px] mt-4 rounded-xl mr-6 p-4 bg-white border-neutral-200 border flex-shrink-0 overflow-hidden"
            >
              <div className="flex justify-center">
                <Image
                  className="rounded-full"
                  src={item.image}
                  alt={item.name}
                  width={100}
                  height={100}
                />
              </div>
              <h1 className="font-semibold text-center text-md pt-4 ">
                <span className="text-[#00A92F]">{firstWord}</span>{" "}
                <span className="text-black">{restOfName}</span>
              </h1>
              <p className="line-clamp-2 text-center text-sm mt-2">
                {item.role}
              </p>
              <p className="line-clamp-2 text-center text-sm">{item.company}</p>
            </motion.div>
          );
        })}
      </div>
    </div>
  );
}
