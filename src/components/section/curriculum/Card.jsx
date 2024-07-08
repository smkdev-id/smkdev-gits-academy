import { FaUserGraduate, FaWandMagicSparkles, FaComputer } from "react-icons/fa6";

const Card = () => {
  return (
    <div className="flex flex-wrap w-full gap-9">
      <div className="flex flex-col w-full text-gray-700 bg-white border shadow-md bg-clip-border rounded-xl group hover:shadow-lg">
        <div className="p-6">
          <div className="flex w-full gap-5">
            <div className="flex items-center justify-center w-1/5 md:items-end lg:items-end">
              <FaUserGraduate className="w-12 h-auto mb-5 text-gray-800 transition duration-150 ease-in-out group-hover:scale-110" />
            </div>
            <div className="flex flex-col w-4/5">
              <h5 className="block mb-2 font-sans text-xl antialiased font-bold leading-snug tracking-normal text-blue-gray-900">
                Kurikulum
              </h5>
              <p className="block font-sans text-sm antialiased leading-relaxed md:text-base lg:text-base text-inherit">
                Kurikulum berstandar industri global, senantiasa diperbaharui
              </p>
            </div>
          </div>
        </div>
      </div>
      <div className="flex flex-col w-full text-gray-700 bg-white border shadow-md bg-clip-border rounded-xl group hover:shadow-lg">
        <div className="p-6">
          <div className="flex w-full gap-5">
            <div className="flex items-center justify-center w-1/5 transition duration-150 ease-in-out md:items-end lg:items-end group-hover:scale-110">
              <FaWandMagicSparkles className="w-12 h-auto mb-5 text-gray-800" />
            </div>
            <div className="flex flex-col w-4/5">
              <h5 className="block mb-2 font-sans text-xl antialiased font-bold leading-snug tracking-normal text-blue-gray-900">
                Metode
              </h5>
              <p className="block font-sans text-sm antialiased leading-relaxed md:text-base lg:text-base text-inherit">
                Metode belajar project-based learning, dapat menjadi portfolio siswa
              </p>
            </div>
          </div>
        </div>
      </div>
      <div className="flex flex-col w-full text-gray-700 bg-white border shadow-md bg-clip-border rounded-xl group hover:shadow-lg">
        <div className="p-6">
          <div className="flex w-full gap-5">
            <div className="flex items-center justify-center w-1/5 md:items-end lg:items-end">
              <FaComputer className="w-12 h-auto mb-5 text-gray-800 transition duration-150 ease-in-out group-hover:scale-110" />
            </div>
            <div className="flex flex-col w-4/5">
              <h5 className="block mb-2 font-sans text-xl antialiased font-bold leading-snug tracking-normal text-blue-gray-900">
                Mentor
              </h5>
              <p className="block font-sans text-sm antialiased leading-relaxed md:text-base lg:text-base text-inherit">
                Dibimbing langsung oleh mentor expert dari industri digital
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Card;