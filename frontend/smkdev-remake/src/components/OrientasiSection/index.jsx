import ScrollAnimation from "./ScrollAnimation";

const OrientasiSection = () => {
  return (
    <section>
      <div className="w-full py-24">
        <div className="flex justify-center pb-5">
          <h1 className="text-3xl font-medium text-custom-blue font-poppins">
            Orientasi Belajar SMKDEV
          </h1>
        </div>
        <div className="flex justify-center ">
          <p className="text-lg text-center font-poppins ">
            Dapatkan pengalaman belajar berorientasi pengalaman kerja yang dapat
            mengantarkan Anda <br /> menjadi talenta yang dibutuhkan oleh
            industri digital terkini.
          </p>
        </div>
      </div>
      <div className="pb-20">
        <ScrollAnimation />
      </div>
    </section>
  );
};

export default OrientasiSection;
