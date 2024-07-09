import Blog from "./blog";
import Comments from "./comments";
import CompanyList from "./company";
import Contact from "./contact";
import Curriculum from "./curriculum";
import Hero from "./hero";
import Learning from "./learning";
import Mentor from "./mentor";

import { IoLogoWhatsapp } from "react-icons/io";

const Section = () => {
  return (
    <div className="w-full min-h-screen px-5 lg:px-20 md:px-20">
      <Hero />
      <CompanyList />
      <Learning />
      <Curriculum />
      <Comments />
      <Mentor />
      <Blog />
      <Contact />
      {/* <IoLogoWhatsapp className="cursor-pointer text-white justify-center items-center fixed z-[999] bottom-4 right-4 p-4 block h-14 w-14 btn btn-primary btn-circle" /> */}
    </div>
  );
};

export default Section;
