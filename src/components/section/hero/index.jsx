import { GoArrowRight } from "react-icons/go";
import img from "../../../assets/image";

const Hero = () => {
  return (
    <div className="relative pt-20 md:pt-0 lg:pt-0">
      <section className="py-2 lg:py-20 md:py-20">
        <div className="items-center justify-between max-w-screen-xl mx-auto overflow-hidden text-gray-600 gap-x-12 md:flex md:px-8">
          <div className="flex-none w-full px-4 space-y-5 md:w-1/2 lg:w-1/2">
            <h2 className="text-4xl font-extrabold text-gray-800 md:text-5xl">
              Jadilah Talenta Digital Masa Depan Indonesia
            </h2>
            <p>
              Belajar langsung dengan expert dari industri dengan kurikulum
              komperhensif berbasis project-based learning
            </p>
            <div className="items-center space-y-3 gap-x-3 sm:flex sm:space-y-0">
              <a
                href="#company"
                className="font-medium text-center text-white duration-150 btn bg-violet-800/95 hover:bg-violet-800"
              >
                Let's get started 🚀
              </a>
              <a className="flex items-center justify-center text-white btn bg-dark hover:bg-black/80 gap-x-2 md:inline-flex">
                Get access
                <GoArrowRight />
              </a>
            </div>
          </div>
          <div className="flex-none mt-14 md:w-1/2 lg:w-1/2">
            <img
              src={img?.skills}
              className="w-full h-auto md:rounded-tl-[108px] lg:rounded-tl-[108px] rounded-2xl"
              alt="img skill"
              width={500}
              height={500}
            />
          </div>
        </div>
      </section>
    </div>
  );
};

export default Hero;