import { IoIosBookmarks, IoIosBook, IoIosDesktop } from "react-icons/io";
import Card from "./Card";

const Learning = () => {
  const learning = [
    {
      title: "Learning Path Industri",
      description:
        "Learning path disusun berdasarkan kebutuhan dunia industri.",
      icon: IoIosBookmarks,
    },
    {
      title: "Kurikulum Komprehensif",
      description:
        "Kurikulum komprehensif dan senantiasa di perbarui berdasarkan pengalaman di dunia industri.",
      icon: IoIosBook,
    },
    {
      title: "Project-Based Learning",
      description:
        "Materi pembelajaran disusun dengan pendekatan - pendekatan project based learning.",
      icon: IoIosDesktop,
    },
  ];

  return (
    <div className="py-10">
      <div className="flex flex-col gap-4">
        <h1 className="text-3xl font-bold text-center text-gray-700">
          Orientasi Belajar SMKDEV
        </h1>
        <p className="text-center text-gray-500 lg:text-xl md:text-xl text-wrap">
          Dapatkan pengalaman belajar berorientasi pengalaman kerja yang dapat
          mengantarkan Anda menjadi talenta yang dibutuhkan oleh industri
          digital terkini.
        </p>
      </div>
      <div className="flex flex-wrap w-full gap-6 py-10 md:flex-nowrap lg:flex-nowrap">
        {learning.map((data, i) => {
          return (
            <Card
              key={i}
              title={data.title}
              description={data.description}
              Icon={data.icon}
            />
          );
        })}
      </div>
    </div>
  );
};

export default Learning;