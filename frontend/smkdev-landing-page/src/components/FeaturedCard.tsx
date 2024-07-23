import { FC } from "react";

interface Props {
  title: string;
  desc: string;
  svg: React.ReactNode;
}

const FeaturedCard: FC<Props> = ({ title, desc, svg }) => {
  return (
    <div className="flex w-full flex-col justify-between overflow-hidden rounded-2xl shadow-md">
      <div className="w-fit rounded-br-[40px] bg-secondary bg-opacity-5 p-8">
        {svg}
      </div>
      <div className="w-full px-8 py-8">
        <h1 className="mb-4 text-xl font-semibold capitalize text-primary">
          {title}
        </h1>
        <p className="text-base text-[#828282]">{desc}</p>
      </div>
    </div>
  );
};

export default FeaturedCard;
