const Card = ({ title, description, Icon }) => {
  return (
    <div className="flex flex-col w-full text-gray-700 bg-white border shadow-md bg-clip-border rounded-xl group hover:shadow-lg">
      <div className="p-6">
        <div className="flex w-full gap-5">
          <div className="flex items-center justify-center w-1/5 md:items-end lg:items-end">
            <Icon className="w-12 h-auto mb-5 text-gray-800 transition duration-150 ease-in-out group-hover:scale-110" />
          </div>
          <div className="flex flex-col w-4/5">
            <h5 className="block mb-2 font-sans text-xl antialiased font-bold leading-snug tracking-normal text-blue-gray-900">
              {title}
            </h5>
            <p className="block font-sans text-sm antialiased leading-relaxed md:text-base lg:text-base text-inherit">
              {description}
            </p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Card;
