import { MediaQuery } from "@/components";
import { NextPage } from "next";

interface Props {}

const Mentor: NextPage<Props> = ({}) => {
  return (
    <div className="flex w-full justify-center py-20">
      <MediaQuery className="flex flex-col items-center justify-center gap-y-5">
        <div className="mb-10 flex w-full flex-col items-center gap-y-3">
          <h1 className="text-center text-3xl font-semibold capitalize text-primary">
            Expert Mentor
          </h1>
          <p className="w-[80%] text-center text-xl font-medium">
            SMKDEV menghadirkan expert dari industri untuk mendampingi proses
            belajarmu.
          </p>
        </div>

        <div className=""></div>
      </MediaQuery>
    </div>
  );
};

export default Mentor;
