import { FC } from "react";
import MediaQuery from "./MediaQuery";
import Image from "next/image";

interface Props {}

const Footer: FC<Props> = ({}) => {
  return (
    <>
      <div className="bg-bg2 flex w-full justify-center py-10">
        <MediaQuery className="flex flex-col gap-y-5">
          <div className="flex w-full flex-col gap-y-4">
            <div className="">
              <div className="mb-3 w-[130px]">
                <Image
                  src="/assets/Logo-Smkdev.png"
                  alt="logo"
                  width="0"
                  height="0"
                  sizes="100vw"
                  className="h-auto w-full"
                />
              </div>
              <h1 className="text-base font-bold capitalize">
                Creating High-Caliber Digital Talent
              </h1>
            </div>
            <p className="text-base">
              Summarecon Bandung, Jl. Magna Barat Blok MD No.02, Rancabolang,
              Kec. Gedebage, Kota Bandung, Jawa Barat 40295
            </p>
          </div>

          <div className="flex flex-col gap-y-3">
            <h1 className="text-base font-bold capitalize text-primary">
              Follow Us
            </h1>
            <div></div>
          </div>
        </MediaQuery>
      </div>
      <div className="flex w-full justify-center py-5">
        <p className="font-medium text-primary">
          © 2024 SMKDEV – PT Eureka Merdeka Indonesia. All Rights Reserved.
        </p>
      </div>
    </>
  );
};

export default Footer;
