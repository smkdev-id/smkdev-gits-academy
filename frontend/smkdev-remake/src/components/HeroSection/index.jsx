import TextTyping from "./TextTyping";

const HeroSection = () => {
  return (
    <section className="flex justify-center mt-16 py-16 ">
      <div className="flex justify-between gap-64 items-center">
        <div>
          <TextTyping />
          <div>
            <p className="text-lg  text-custom-color-font font- pt-10 font-poppins">
              Belajar langsung dengan expert dari industri dengan <br />{" "}
              kurikulum komperhensif berbasis project-based learning
            </p>
          </div>
        </div>
        <div>
          <img
            src="../../../public/hero.jpg"
            alt="hero-section"
            width={389}
            height={389}
            className="rounded-3xl"
          />
        </div>
      </div>
    </section>
  );
};

export default HeroSection;
