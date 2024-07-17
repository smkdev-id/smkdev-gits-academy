const Card = ({ title, value, Icon }) => {
  return (
    <div className="flex flex-col items-center justify-center w-full p-6 mt-6 text-gray-700 lg:w-1/3 md:w-1/3 bg-clip-border rounded-xl">
      <div className="flex items-center justify-center px-4 py-2 my-1 rounded bg-dark">
        <Icon className="w-10 h-10 text-white" />
      </div>
      <h5 className="block mb-2 font-sans text-2xl antialiased font-bold leading-snug tracking-normal text-blue-gray-900">
        {title} :
      </h5>
      <p className="block font-sans text-sm antialiased leading-relaxed text-center md:text-base lg:text-base min-h-20 text-inherit text-slate-600">
        {value}
      </p>
    </div>
  );
};

export default Card;
