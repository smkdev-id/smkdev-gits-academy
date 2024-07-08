import Card from "./Card";

const Curriculum = () => {
  return (
    <div className="flex flex-wrap gap-8 md:flex-nowrap lg:flex-nowrap">
      <div className="w-full md:w-1/2 lg:w-1/2">
        <Card />
      </div>
      <div className="flex flex-col items-center justify-center w-full gap-5 md:w-1/2 lg:w-1/2">
        <h1 className="text-2xl font-bold lg:text-3xl md:text-3xl">MENGAPA HARUS MEMILIH BELAJAR DI <span className="text-primary">SMKDEV</span></h1>
        <p className="text-base md:text-lg lg:text-lg text-pretty">
          Pengalaman belajarmu jauh lebih menyenangkan dengan perpaduan belajar teori dan capstone project yang dapat menjadi portofolio pribadimu. Tidak lupa, Kamu juga dapat belajar langsung dari expert mentor dari dunia industri.
        </p>
      </div>
    </div>
  );
};

export default Curriculum;