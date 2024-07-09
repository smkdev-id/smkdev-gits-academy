import { IoLogoInstagram, IoLogoLinkedin } from "react-icons/io";
import Header from "../ui/Header";

const Mentor = () => {
  // example datas
  const mentors = [
    {
      name: "Helmi Adi Prasetyo",
      profile:
        "https://smkdev.storage.googleapis.com/wp/image-4ecbfe-4eee03-300x300.png",
      description: "Back End Developer",
      at: "SMKDEV",
    },
    {
      name: "Samuel Pandohan Terampil Gultom",
      profile:
        "https://smkdev.storage.googleapis.com/wp/Samuel-Pandohan-Terampil-Gultom.jpg",
      description: "Curriculum Developer",
      at: "SMKDEV",
    },
    {
      name: "Rachmawati Ari Taurisia",
      profile:
        "https://smkdev.storage.googleapis.com/wp/Rachmawati-Ari-Taurisia.jpg",
      description: "Curriculum Developer",
      at: "SMKDEV",
    },
    {
      name: "Sudaryatno",
      profile:
        "https://smkdev.storage.googleapis.com/wp/sudaryatno-fe5603-fe6f61-300x300.jpeg",
      description: "CTO",
      at: "GITS Indonesia",
    },
    {
      name: "Ibnu Sina Wardy",
      profile:
        "https://smkdev.storage.googleapis.com/wp/ibnu-sina-wardy-86f446-870bf6-a31bbc-a34c39-300x300.jpg",
      description: "CEO",
      at: "GITS Indonesia",
    },
  ];

  return (
    <div className="pb-10">
      <Header
        title="Expert Mentor"
        description="SMKDEV menghadirkan expert dari industri untuk mendampingi proses belajarmu"
      />

      <div className="grid gap-8 pt-5 lg:gap-16 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-5">
        {mentors?.map((mentor, i) => (
          <div key={i} className="text-center text-gray-500 dark:text-gray-400">
            <img
              className="mx-auto mb-4 rounded-full w-36 h-36"
              src={mentor.profile}
              alt={`${mentor.name}'s profile`}
            />
            <div className="flex flex-col">
              <p className="flex items-center justify-center h-16 mb-1 text-lg font-semibold text-gray-700">
                {mentor.name}
              </p>
              <p className="text-base">{mentor.description}</p>
              <p className="text-base">{mentor.at}</p>
            </div>
            <ul className="flex justify-center mt-4 space-x-4">
              <li>
                <IoLogoInstagram className="text-gray-900 w-7 h-7 hover:scale-105" />
              </li>
              <li>
                <IoLogoLinkedin className="text-gray-900 w-7 h-7 hover:scale-105" />
              </li>
            </ul>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Mentor;