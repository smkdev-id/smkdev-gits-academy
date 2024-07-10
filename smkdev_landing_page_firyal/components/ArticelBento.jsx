import { cn } from "@/utils/cn";
import React from "react";
import {
  IconArrowWaveRightUp,
  IconBoxAlignRightFilled,
  IconBoxAlignTopLeft,
  IconClipboardCopy,
  IconFileBroken,
  IconSignature,
  IconTableColumn,
} from "@tabler/icons-react";
import { BentoGrid, BentoGridItem } from "./ui/Bento";
import Image from "next/image";

export function ArticelBento() {
  return (
    <div>
      <div className="flex flex-col justify-center text-center mt-32">
        <h1 className="text-3xl font-bold">
          <span className="text-[#004FC5] mr-2">Artikel</span>
          Pilihan SMKDEV
        </h1>
        <h1 className="text-xl flex pt-2 text-center justify-center">
          Temukan Informasi Terkini dan Terpercaya untuk Menjadi Ahli di Dunia
          Digital.
        </h1>
      </div>
      <BentoGrid className="max-w-6xl pt-4">
        {items.map((item, i) => (
          <BentoGridItem
            key={i}
            title={item.title}
            description={item.description}
            header={item.header}
            icon={item.icon}
          />
        ))}
      </BentoGrid>
    </div>
  );
}
const Skeleton = ({ img }) => (
  <div className="flex flex-1 w-full h-full min-h-[6rem] rounded-xl overflow-hidden relative">
    <Image
      alt="article1"
      src={img}
      layout="fill"
      objectFit="cover"
      className="rounded-xl"
    />
  </div>
);
const items = [
  {
    title: "Prompt Engineering: Keterampilan, Prospek Karir Dan Kursus AI",
    description:
      "Kenapa Prompt Engineering penting untuk AI generatif seperti ChatGPT dan Google Gemini?",
    header: <Skeleton img={"/article1.jpg"} />,
    icon: <IconClipboardCopy className="h-4 w-4 " />,
  },
  {
    title: "Panduan Pemula Menguasai AI Tools (Alat AI) ",
    description:
      "Temukan langkah-langkah, alat populer, dan tips untuk sukses dalam memanfaatkan teknologi.",
    header: <Skeleton img={"/article2.jpg"} />,
    icon: <IconFileBroken className="h-4 w-4 text-neutral-500" />,
  },
  {
    title: "Prompt Engineering: Keterampilan, Prospek Karir Dan Kursus AI",
    description:
      "Kenapa Prompt Engineering penting untuk AI generatif seperti ChatGPT dan Google Gemini?",
    header: <Skeleton img={"/article3.jpg"} />,
    icon: <IconSignature className="h-4 w-4 text-neutral-500" />,
  },

  {
    title: "Panduan Pemula Menguasai AI Tools (Alat AI) ",
    description:
      "Temukan langkah-langkah, alat populer, dan tips untuk sukses dalam memanfaatkan teknologi.",
    header: <Skeleton img={"/article4.jpg"} />,
    icon: <IconFileBroken className="h-4 w-4 text-neutral-500" />,
  },
  {
    title: "Prompt Engineering: Keterampilan, Prospek Karir Dan Kursus AI",
    description:
      "Kenapa Prompt Engineering penting untuk AI generatif seperti ChatGPT dan Google Gemini?",
    header: <Skeleton img={"/article1.jpg"} />,
    icon: <IconBoxAlignRightFilled className="h-4 w-4 text-neutral-500" />,
  },
  {
    title: "Panduan Pemula Menguasai AI Tools (Alat AI) ",
    description:
      "Temukan langkah-langkah, alat populer, dan tips untuk sukses dalam memanfaatkan teknologi.",
    header: <Skeleton img={"/article4.jpg"} />,
    icon: <IconTableColumn className="h-4 w-4 text-neutral-500" />,
  },
];
