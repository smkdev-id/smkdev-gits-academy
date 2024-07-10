import { MediaQuery, MentorList } from "@/components";
import { NextPage } from "next";

interface Props {}

const Mentor: NextPage<Props> = ({}) => {
  const MENTOR_LIST = [
    {
      name: "Sudaryatno",
      role: "CTO",
      at: "GITS Indonesia",
      img: "https://smkdev.storage.googleapis.com/wp/sudaryatno-fe5603-fe6f61-300x300.jpeg",
    },
    {
      name: "Ibnu Sina Wardy",
      role: "CEO",
      at: "GITS Indonesia",
      img: "https://smkdev.storage.googleapis.com/wp/ibnu-sina-wardy-86f446-870bf6-a31bbc-a34c39-300x300.jpg",
    },
    {
      name: "Helmi Adi Prasetyo",
      role: "Backend Developer",
      at: "SMKDEV",
      img: "https://smkdev.storage.googleapis.com/wp/image-4ecbfe-4eee03-300x300.png",
    },
    {
      name: "Samuel Pandohan Terampil Gultom",
      role: "Curriculum Developer",
      at: "SMKDEV",
      img: "https://smkdev.storage.googleapis.com/wp/Samuel-Pandohan-Terampil-Gultom.jpg",
    },
    {
      name: "Rachmawati Ari Taurisia",
      role: "Curriculum Developer",
      at: "SMKDEV",
      img: "https://smkdev.storage.googleapis.com/wp/Rachmawati-Ari-Taurisia.jpg",
    },
  ];

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

        <div className="w-full">
          <MentorList list={MENTOR_LIST} />
        </div>
      </MediaQuery>
    </div>
  );
};

export default Mentor;
