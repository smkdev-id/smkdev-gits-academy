const CompanyList = () => {
  const images = {
    imageLink: [
      "https://smkdev.storage.googleapis.com/wp/logo-hitopia-4aaf6b-4acd41.png",
      "https://smkdev.storage.googleapis.com/wp/Logo-Mantab-One.png",
      "https://smkdev.storage.googleapis.com/wp/Logo-Eudeka.png",
      "https://smkdev.storage.googleapis.com/wp/Logo-Gits-Indonesia.png",
      "https://smkdev.storage.googleapis.com/wp/Logo-DIGITS.png",
      "https://smkdev.storage.googleapis.com/wp/Logo-Arkana.png",
    ],
  };

  return (
    <div className="px-4 py-10" id="company">
      <h1 className="text-xl font-bold text-center text-gray-700">
        Dipercaya Oleh Mitra Industri
      </h1>
      <div className="flex flex-wrap items-center justify-center mt-5 gap-x-5 gap-y-6">
        {images?.imageLink?.map((link, i) => {
          return (
            <div key={i} className="w-24 h-24 sm:w-32 sm:h-32 md:w-40 md:h-40">
              <img
                src={link}
                alt={link}
                className="object-contain w-full h-full transition-all duration-150 ease-in cursor-pointer hover:scale-110"
              />
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default CompanyList;
