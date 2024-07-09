import SliderComponent from "./SliderExpertMentor";

const ExpertMentorSection = () => {
  return (
    <section className=" my-24 bg-custom-blue-bg">
      <div className="container mx-auto py-7">
        <div className="grid justify-center">
          <h1 className="text-3xl font-medium font-poppins mb-5 text-custom-color-font text-center">
            Expert Mentor
          </h1>
          <p className="text-center text-xl text-custom-color-font font-normal font-poppins">
            SMKDEV menghadirkan expert daeri industri untuk mendampingi proses
            belajarmu.
          </p>
        </div>
        <SliderComponent />
      </div>
    </section>
  );
};

export default ExpertMentorSection;
