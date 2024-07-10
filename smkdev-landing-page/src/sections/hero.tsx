"use client";
import { MediaQuery } from "@/components";
import { NextPage } from "next";
import Image from "next/image";
import Typewriter from "typewriter-effect";

interface Props {}

const Hero: NextPage<Props> = ({}) => {
  return (
    <div className="flex w-full justify-center py-40">
      <MediaQuery className="flex items-center justify-center gap-x-8">
        <div className="flex flex-col gap-y-9">
          <h1 className="text-5xl font-semibold capitalize">
            jadilah talenta digital{" "}
            <Typewriter
              options={{
                strings: ["Global", "Masa Depan Indonesia"],
                autoStart: true,
                loop: true,
                wrapperClassName: "text-[#004fc5]",
                cursorClassName: "text-[#004fc5]",
              }}
            />
          </h1>
          <p className="text-xl font-medium">
            Belajar langsung dengan expert dari industri dengan kurikulum
            komperhensif berbasis project-based learning
          </p>
        </div>
        <div className="relative">
          <div className="z-20 w-[415px]">
            <Image
              src="/assets/person1.jpg"
              alt="Person"
              width="0"
              height="0"
              sizes="100vw"
              className="h-auto w-full rounded-[40px]"
            />
          </div>
          <div className="absolute -left-12 -top-12 -z-10">
            <Image
              src="/assets/graphics-2.svg"
              alt="graphics-2"
              width="0"
              height="0"
              sizes="100vw"
              className="h-auto w-full"
            />
          </div>
        </div>
      </MediaQuery>
    </div>
  );
};

export default Hero;
