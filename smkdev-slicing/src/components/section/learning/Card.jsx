const Card = ({ title, description, Icon }) => {
  return (
    <div className="flex flex-col w-full mt-6 text-gray-700 bg-white border shadow lg:w-1/3 md:w-1/3 bg-clip-border rounded-xl">
      <div className="p-6">
        <Icon className="w-10 h-10 mb-5 text-gray-800" />
        <h5 className="block mb-2 font-sans text-xl antialiased font-semibold leading-snug tracking-normal text-blue-gray-900">
          {title}
        </h5>
        <p className="block font-sans text-sm antialiased font-light leading-relaxed md:text-base lg:text-base min-h-20 text-inherit">
          {description}
        </p>
      </div>
      <div className="p-6 pt-0">
        <a href="#" className="inline-block">
          <button
            className="flex items-center gap-2 py-2 font-sans text-xs font-bold text-center text-gray-900 uppercase align-middle transition-all rounded-lg select-none hover:underline-offset-1 hover:underline"
            type="button"
          >
            Learn More
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth="2"
              stroke="currentColor"
              className="w-4 h-4"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M17.25 8.25L21 12m0 0l-3.75 3.75M21 12H3"
              ></path>
            </svg>
          </button>
        </a>
      </div>
    </div>
  );
};

export default Card;