import { FeaturedCard } from "@/components";
import { NextPage } from "next";

const Featured: NextPage = () => {
  const cards = [
    {
      title: "Learning Path Industri",
      desc: "Learning path disusun berdasarkan kebutuhan dunia industri.",
      svg: (
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width={44}
          height={50}
          viewBox="0 0 44 50"
          fill="none"
        >
          <g clipPath="url(#clip0)">
            <path
              d="M9.375 0C4.19922 0 0 4.19922 0 9.375V40.625C0 45.8008 4.19922 50 9.375 50H37.5H40.625C42.3535 50 43.75 48.6035 43.75 46.875C43.75 45.1465 42.3535 43.75 40.625 43.75V37.5C42.3535 37.5 43.75 36.1035 43.75 34.375V3.125C43.75 1.39648 42.3535 0 40.625 0H37.5H9.375ZM9.375 37.5H34.375V43.75H9.375C7.64648 43.75 6.25 42.3535 6.25 40.625C6.25 38.8965 7.64648 37.5 9.375 37.5ZM12.5 14.0625C12.5 13.2031 13.2031 12.5 14.0625 12.5H32.8125C33.6719 12.5 34.375 13.2031 34.375 14.0625C34.375 14.9219 33.6719 15.625 32.8125 15.625H14.0625C13.2031 15.625 12.5 14.9219 12.5 14.0625ZM14.0625 18.75H32.8125C33.6719 18.75 34.375 19.4531 34.375 20.3125C34.375 21.1719 33.6719 21.875 32.8125 21.875H14.0625C13.2031 21.875 12.5 21.1719 12.5 20.3125C12.5 19.4531 13.2031 18.75 14.0625 18.75Z"
              fill="url(#paint0_linear)"
            />
          </g>
          <defs>
            <linearGradient
              id="paint0_linear"
              x1={3}
              y1="0"
              x2={55}
              y2={62}
              gradientUnits="userSpaceOnUse"
            >
              <stop stopColor="#005417" />
              <stop offset="0.5" stopColor="#00A92F" />
              <stop offset={1} stopColor="#AAE2BA" />
            </linearGradient>
            <clipPath id="clip0">
              <rect width="43.75" height={50} fill="white" />
            </clipPath>
          </defs>
        </svg>
      ),
    },
    {
      title: "Kurikulum Komprehensif",
      desc: "Kurikulum komprehensif dan senantiasa di perbarui berdasarkan pengalaman di dunia industri",
      svg: (
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width={57}
          height={50}
          viewBox="0 0 57 50"
          fill="none"
        >
          <g clipPath="url(#clip0)">
            <path
              d="M24.375 46.0449C25.4297 46.416 26.5625 45.6445 26.5625 44.5312V7.67578C26.5625 7.26562 26.4062 6.85547 26.0742 6.60156C24.1602 5.07812 19.7656 3.125 14.0625 3.125C9.13086 3.125 4.52148 4.42383 1.76758 5.47852C0.664062 5.9082 0 7.00195 0 8.18359V44.3457C0 45.5078 1.25 46.3184 2.35352 45.957C5.42969 44.9316 10.3027 43.75 14.0625 43.75C17.373 43.75 21.7773 45.1172 24.375 46.0449ZM31.875 46.0449C34.4727 45.1172 38.877 43.75 42.1875 43.75C45.9473 43.75 50.8203 44.9316 53.8965 45.957C55 46.3281 56.25 45.5078 56.25 44.3457V8.18359C56.25 7.00195 55.5859 5.9082 54.4824 5.48828C51.7285 4.42383 47.1191 3.125 42.1875 3.125C36.4844 3.125 32.0898 5.07812 30.1758 6.60156C29.8535 6.85547 29.6875 7.26562 29.6875 7.67578V44.5312C29.6875 45.6445 30.8301 46.416 31.875 46.0449Z"
              fill="url(#paint0_linear)"
            />
          </g>
          <defs>
            <linearGradient
              id="paint0_linear"
              x1={10}
              y1={3}
              x2="52.5"
              y2={48}
              gradientUnits="userSpaceOnUse"
            >
              <stop stopColor="#005417" />
              <stop offset="0.526042" stopColor="#00A92F" />
              <stop offset={1} stopColor="#AAE2BA" />
            </linearGradient>
            <clipPath id="clip0">
              <rect width="56.25" height={50} fill="white" />
            </clipPath>
          </defs>
        </svg>
      ),
    },
    {
      title: "Project-Based Learning",
      desc: "Materi pembelajaran disusun dengan pendekatan - pendekatan project based learning",
      svg: (
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width={63}
          height={50}
          viewBox="0 0 63 50"
          fill="none"
        >
          <g clipPath="url(#clip0)">
            <path
              d="M12.5 0C9.05273 0 6.25 2.80273 6.25 6.25V28.125H1.875C0.839844 28.125 0 28.9648 0 30C0 34.1406 3.35938 37.5 7.5 37.5H31.25V28.125H12.5V6.25H43.75V9.375H50V6.25C50 2.80273 47.1973 0 43.75 0H12.5ZM50 12.5H39.0625C36.4746 12.5 34.375 14.5996 34.375 17.1875V45.3125C34.375 47.9004 36.4746 50 39.0625 50H57.8125C60.4004 50 62.5 47.9004 62.5 45.3125V25H53.125C51.3965 25 50 23.6035 50 21.875V12.5ZM53.125 12.5V21.875H62.5L53.125 12.5Z"
              fill="url(#paint0_linear)"
            />
          </g>
          <defs>
            <linearGradient
              id="paint0_linear"
              x1={0}
              y1={0}
              x2={69}
              y2={57}
              gradientUnits="userSpaceOnUse"
            >
              <stop stopColor="#005417" />
              <stop offset="0.494792" stopColor="#00A92F" />
              <stop offset={1} stopColor="#AAE2BA" />
            </linearGradient>
            <clipPath id="clip0">
              <rect width="62.5" height={50} fill="white" />
            </clipPath>
          </defs>
        </svg>
      ),
    },
  ];

  return (
    <div className="flex w-full justify-center py-20">
      <div className="flex w-[60%] flex-col items-center justify-center">
        <div className="mb-10 flex w-full flex-col items-center gap-y-3">
          <h1 className="text-center text-3xl font-semibold text-primary">
            Orientasi Belajar SMKDEV
          </h1>
          <p className="w-[80%] text-center text-xl font-medium">
            Dapatkan pengalaman belajar berorientasi pengalaman kerja yang dapat
            mengantarkan Anda menjadi talenta yang dibutuhkan oleh industri
            digital terkini.
          </p>
        </div>
        <div className="flex w-full justify-between gap-x-6">
          {cards.map((item, index) => (
            <FeaturedCard
              key={index}
              title={item.title}
              desc={item.desc}
              svg={item.svg}
            />
          ))}
        </div>
      </div>
    </div>
  );
};

export default Featured;
