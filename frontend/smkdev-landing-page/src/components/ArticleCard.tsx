import Image from "next/image";
import Link from "next/link";
import { FC } from "react";

interface Props {
  title: string;
  desc: string;
  tag: string;
  img: string;
  url: string;
}

const ArticleCard: FC<Props> = ({ desc, img, tag, title, url }) => {
  return (
    <div className="relative flex h-[300px] w-full cursor-pointer items-start rounded-lg border-gray-300 shadow-lg">
      <div className="relative h-full w-1/2 overflow-hidden">
        <Image
          alt={title}
          src={img}
          fill
          style={{ objectFit: "cover" }}
          className="rounded-lg rounded-br-none rounded-tr-none"
        />
      </div>
      <div className="flex h-full w-full flex-col justify-between p-7">
        <div>
          <h1 className="mb-3 text-xl font-semibold capitalize text-primary">
            {title}
          </h1>
          <p className="line-clamp-4">{desc}</p>
        </div>
        <Link
          href={url}
          className="cursor-pointer text-sm font-medium uppercase hover:underline"
        >
          baca lebih lanjut &gt;&gt;
        </Link>
      </div>

      <div className="absolute left-3 top-3 rounded-full bg-primary bg-opacity-70 px-3 py-1 text-white">
        <span className="font-mono text-sm uppercase text-white">
          {tag === "" ? "general" : tag}
        </span>
      </div>
    </div>
  );
};

export default ArticleCard;
