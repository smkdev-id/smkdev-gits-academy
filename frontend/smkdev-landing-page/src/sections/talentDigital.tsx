import { MediaQuery } from "@/components";
import { NextPage } from "next";
import Image from "next/image";

interface Props {}

const TalentDigital: NextPage<Props> = ({}) => {
  return (
    <div className="flex w-full justify-center py-20">
      <MediaQuery className="flex flex-col items-center justify-center gap-y-5">
        <h1 className="text-3xl font-semibold text-primary">
          Talenta Digital SMKDEV
        </h1>

        <div className="w-[75%]">
          <Image
            src="/assets/talenta.png"
            alt="Talenta Digital"
            width="0"
            height="0"
            sizes="100vw"
            className="h-auto w-full"
          />
        </div>
      </MediaQuery>
    </div>
  );
};

export default TalentDigital;
