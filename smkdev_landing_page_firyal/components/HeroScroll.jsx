"use client";
import React from "react";
import Image from "next/image";
import { ContainerScroll } from "./ui/ContainerScrollAnimation";
import { TypewriterEffect } from "./ui/TypewritterEffect";

export function HeroScroll() {
  const words = [
    {
      text: "Jadilah",
    },
    {
      text: "Talenta",
    },
    {
      text: "Digital",
    },
    {
      text: "Masa",
    },
    {
      text: "Depan",
    },
    {
      text: "Indonesia.",
      className: "text-[#004FC5]",
    },
  ];
  return (
    <div className="flex flex-col overflow-hidden">
      <ContainerScroll
        titleComponent={
          <>
            <div className="flex flex-col items-center justify-center h-[20rem]  ">
              <TypewriterEffect words={words} />
              <p className="text-black text-md mt-2">
                Kurikulum komprehensif dengan project-based learning dari para
                ahli industri. Siap bersaing di dunia digital!
              </p>
            </div>
          </>
        }
      >
        <Image
          src="/zoom.jpg"
          alt="hero"
          height={720}
          width={1400}
          className="mx-auto rounded-2xl object-cover h-full object-left-top"
          draggable={false}
        />
      </ContainerScroll>
    </div>
  );
}
